// Package internalauth manages authorizing gRPC calls.
package internalauth

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/jrockway/monorepo/jsso2/pkg/sessions"
	"github.com/jrockway/monorepo/jsso2/pkg/store"
	"github.com/jrockway/monorepo/jsso2/pkg/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Config struct {
	RootPassword string `long:"root_password" env:"ROOT_PASSWORD" description:"If set, allow a requestor full privileges if they include this password in their requests.  Should only be used to bootstrap a normal administrative user."`
}

// RPCConfig configures permissions for an RPC.
type RPCConfig struct {
	// An RPC must tolerate all session taints in order to be executed.
	Tolerations []string
}

// Permissions manages all authorization in JSSO.
type Permissions struct {
	// If set, a password that can be provided to bypass all access controls.
	RootPassword string
	RPCConfig    map[string]*RPCConfig
	Store        *store.Connection
	Cookies      *sessions.CookieConfig
}

// NewFromConfig builds a Permissions object from configuration.
func NewFromConfig(c *Config, s *store.Connection) *Permissions {
	return &Permissions{
		Store:        s,
		RootPassword: c.RootPassword,
		RPCConfig: map[string]*RPCConfig{
			"/grpc.health.v1.Health/Check": {
				Tolerations: []string{sessions.TaintAnonymous},
			},
			"/grpc.health.v1.Health/Watch": {
				Tolerations: []string{sessions.TaintAnonymous},
			},
			"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo": {
				Tolerations: []string{sessions.TaintAnonymous},
			},
			"/jsso.User/WhoAmI": {
				Tolerations: []string{sessions.TaintAnonymous},
			},
			"/jsso.Session/AuthorizeHTTP": {
				Tolerations: []string{sessions.TaintAnonymous},
			},
			"/jsso.Enrollment/Start": {
				Tolerations: []string{sessions.TaintEnrollment},
			},
			"/jsso.Enrollment/Finish": {
				Tolerations: []string{sessions.TaintEnrollment},
			},
			"/jsso.Login/Start": {
				Tolerations: []string{sessions.TaintAnonymous},
			},
			"/jsso.Login/Finish": {
				Tolerations: []string{sessions.TaintStartLogin},
			},
		},
	}
}

// AuthorizeRPC returns whether the credentials provided allow the RPC to be called.
func (p *Permissions) AuthorizeRPC(ctx context.Context, session *types.Session, fullMethod string) error {
	haveTaints := make(map[string]struct{})
	for _, t := range session.GetTaints() {
		haveTaints[t] = struct{}{}
	}
	var tolerations []string
	if cfg, ok := p.RPCConfig[fullMethod]; ok {
		tolerations = cfg.Tolerations
	}
	for _, t := range tolerations {
		delete(haveTaints, t)
	}
	var remainingTaints []string
	for k := range haveTaints {
		remainingTaints = append(remainingTaints, k)
	}
	if len(remainingTaints) == 0 {
		return nil
	}
	sort.Strings(remainingTaints)
	return status.Error(codes.PermissionDenied, fmt.Sprintf("rpc does not tolerate session taints %v", remainingTaints))
}

func (p *Permissions) isRoot(md metadata.MD) bool {
	want := fmt.Sprintf("root %s", p.RootPassword)
	for _, auth := range md.Get("Authorization") {
		if auth == want {
			return true
		}
	}
	return false
}

func (p *Permissions) getSession(ctx context.Context) (*types.Session, error) {
	l := ctxzap.Extract(ctx).Named("internalauth")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata in incoming context")
	}

	if p.isRoot(md) {
		return sessions.Root(), nil
	}

	// Extract sessions from Authorization / Cookie headers.
	ss, unusedHeaders, unusedCookies := p.Cookies.SessionsFromMetadata(md)
	session, errs := p.Store.AuthenticateUser(ctx, l, ss, unusedHeaders, unusedCookies)
	if len(errs) > 0 && (len(unusedHeaders) > 0 || len(unusedCookies) > 0) {
		l.Debug("errors may prevent the use of a provided credential", zap.Errors("error", errs))
	}
	if session != nil {
		return session, nil
	}
	return sessions.Anonymous(), nil
}

func (p *Permissions) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		rootCtx := ss.Context()
		session, err := p.getSession(rootCtx)
		if u := session.GetUser().GetUsername(); u != "" {
			ctxzap.AddFields(rootCtx, zap.String("session.user", u))
		}
		if t := session.GetTaints(); len(t) > 0 {
			ctxzap.AddFields(rootCtx, zap.Any("session.taints", t))
		}
		if err != nil {
			return status.Error(codes.Unauthenticated, fmt.Sprintf("get user from session: %v", err))
		}
		ctx := sessions.NewContext(rootCtx, session)
		if err := p.AuthorizeRPC(ctx, session, info.FullMethod); err != nil {
			l := ctxzap.Extract(ctx)
			l.Debug("user not authorized to perform RPC", zap.Error(err))
			return err
		}
		return handler(srv, ss)
	}
}

func (p *Permissions) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(rootCtx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		session, err := p.getSession(rootCtx)
		if u := session.GetUser().GetUsername(); u != "" {
			ctxzap.AddFields(rootCtx, zap.String("session.user", u))
		}
		if t := session.GetTaints(); len(t) > 0 {
			ctxzap.AddFields(rootCtx, zap.Any("session.taints", t))
		}
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("get user from session: %v", err))
		}
		ctx := sessions.NewContext(rootCtx, session)
		if err := p.AuthorizeRPC(ctx, session, info.FullMethod); err != nil {
			l := ctxzap.Extract(ctx)
			l.Debug("user not authorized to perform RPC", zap.Error(err))
			return nil, err
		}
		return handler(ctx, req)
	}
}

func sessionMetadataFromContext(ctx context.Context) *types.SessionMetadata {
	result := &types.SessionMetadata{}
	if ctx == nil {
		return result
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return result
	}
	userAgent := md.Get("user-agent")
	if len(userAgent) == 1 {
		result.UserAgent = userAgent[0]
	}
	// It's assumed that Envoy is always in front of requests and that it will always set this
	// to a legitimate value that the client can't tamper with.
	ip := md.Get("x-forwarded-for")
	if len(ip) == 1 {
		result.IpAddress = ip[0]
	}
	return result
}

// General policy decisions start here.

func (p *Permissions) EnrollmentSessionPrototype(ctx context.Context, target *types.User) (*types.Session, error) {
	id, err := sessions.GenerateID()
	if err != nil {
		return nil, fmt.Errorf("generate session id: %w", err)
	}
	now := time.Now()
	return &types.Session{
		Id:        id,
		User:      target,
		CreatedAt: timestamppb.New(now),
		ExpiresAt: timestamppb.New(now.Add(3 * 24 * time.Hour)),
		Taints:    []string{sessions.TaintEnrollment},
		Metadata:  sessionMetadataFromContext(ctx),
	}, nil
}

func (p *Permissions) LoginSessionPrototype(ctx context.Context, target *types.User) (*types.Session, error) {
	id, err := sessions.GenerateID()
	if err != nil {
		return nil, fmt.Errorf("generate session id: %w", err)
	}
	now := time.Now()
	return &types.Session{
		Id:        id,
		User:      target,
		CreatedAt: timestamppb.New(now),
		ExpiresAt: timestamppb.New(now.Add(18 * time.Hour)),
		Taints:    []string{sessions.TaintStartLogin},
		Metadata:  sessionMetadataFromContext(ctx),
	}, nil
}

func (p *Permissions) AllowRedirect(destination string) error {
	return nil
}

// The per-operation permissions start here.

func (p *Permissions) AllowUserEdit(ctx context.Context, target *types.User, actor *types.Session) error {
	return nil
}

func (p *Permissions) AllowGenerateEnrollmentLink(ctx context.Context, target *types.User, actor *types.Session) error {
	return nil
}

func (p *Permissions) AllowStartEnrollment(ctx context.Context, target *types.Session) error {
	return nil
}

func (p *Permissions) AllowFinishEnrollment(ctx context.Context, target *types.Session) error {
	return nil
}

func (p *Permissions) AllowStartLogin(ctx context.Context, target *types.User) error {
	return nil
}

func (p *Permissions) AllowAuthorizeHTTP(ctx context.Context, proxyUser *types.User) error {
	return nil
}

func (p *Permissions) AllowWebVisit(ctx context.Context, session *types.Session, requestURL *url.URL) error {
	if ts := session.GetTaints(); len(ts) > 0 {
		return fmt.Errorf("session is tainted: %v", ts)
	}
	if id := session.GetUser().GetId(); id < 1 {
		return errors.New("you must be logged in to visit this site")
	}
	return nil
}
