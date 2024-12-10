// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package jssopb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// Edit adds a new user if the ID is 0, or updates an existing user.
	Edit(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserReply, error)
	// GenerateEnrollmentLink generates an enrollment token for the user.
	GenerateEnrollmentLink(ctx context.Context, in *GenerateEnrollmentLinkRequest, opts ...grpc.CallOption) (*GenerateEnrollmentLinkReply, error)
	// WhoAmI returns the user object associated with the current session.  When
	// called without a session, a null current user is returned rather than an
	// error.
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

var userEditStreamDesc = &grpc.StreamDesc{
	StreamName: "Edit",
}

func (c *userClient) Edit(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserReply, error) {
	out := new(EditUserReply)
	err := c.cc.Invoke(ctx, "/jsso.User/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var userGenerateEnrollmentLinkStreamDesc = &grpc.StreamDesc{
	StreamName: "GenerateEnrollmentLink",
}

func (c *userClient) GenerateEnrollmentLink(ctx context.Context, in *GenerateEnrollmentLinkRequest, opts ...grpc.CallOption) (*GenerateEnrollmentLinkReply, error) {
	out := new(GenerateEnrollmentLinkReply)
	err := c.cc.Invoke(ctx, "/jsso.User/GenerateEnrollmentLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var userWhoAmIStreamDesc = &grpc.StreamDesc{
	StreamName: "WhoAmI",
}

func (c *userClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIReply, error) {
	out := new(WhoAmIReply)
	err := c.cc.Invoke(ctx, "/jsso.User/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserService is the service API for User service.
// Fields should be assigned to their respective handler implementations only before
// RegisterUserService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type UserService struct {
	// Edit adds a new user if the ID is 0, or updates an existing user.
	Edit func(context.Context, *EditUserRequest) (*EditUserReply, error)
	// GenerateEnrollmentLink generates an enrollment token for the user.
	GenerateEnrollmentLink func(context.Context, *GenerateEnrollmentLinkRequest) (*GenerateEnrollmentLinkReply, error)
	// WhoAmI returns the user object associated with the current session.  When
	// called without a session, a null current user is returned rather than an
	// error.
	WhoAmI func(context.Context, *WhoAmIRequest) (*WhoAmIReply, error)
}

func (s *UserService) edit(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.User/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Edit(ctx, req.(*EditUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UserService) generateEnrollmentLink(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateEnrollmentLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GenerateEnrollmentLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.User/GenerateEnrollmentLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GenerateEnrollmentLink(ctx, req.(*GenerateEnrollmentLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UserService) whoAmI(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.User/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterUserService registers a service implementation with a gRPC server.
func RegisterUserService(s grpc.ServiceRegistrar, srv *UserService) {
	srvCopy := *srv
	if srvCopy.Edit == nil {
		srvCopy.Edit = func(context.Context, *EditUserRequest) (*EditUserReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
		}
	}
	if srvCopy.GenerateEnrollmentLink == nil {
		srvCopy.GenerateEnrollmentLink = func(context.Context, *GenerateEnrollmentLinkRequest) (*GenerateEnrollmentLinkReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GenerateEnrollmentLink not implemented")
		}
	}
	if srvCopy.WhoAmI == nil {
		srvCopy.WhoAmI = func(context.Context, *WhoAmIRequest) (*WhoAmIReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "jsso.User",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Edit",
				Handler:    srvCopy.edit,
			},
			{
				MethodName: "GenerateEnrollmentLink",
				Handler:    srvCopy.generateEnrollmentLink,
			},
			{
				MethodName: "WhoAmI",
				Handler:    srvCopy.whoAmI,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "jsso.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewUserService creates a new UserService containing the
// implemented methods of the User service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewUserService(s interface{}) *UserService {
	ns := &UserService{}
	if h, ok := s.(interface {
		Edit(context.Context, *EditUserRequest) (*EditUserReply, error)
	}); ok {
		ns.Edit = h.Edit
	}
	if h, ok := s.(interface {
		GenerateEnrollmentLink(context.Context, *GenerateEnrollmentLinkRequest) (*GenerateEnrollmentLinkReply, error)
	}); ok {
		ns.GenerateEnrollmentLink = h.GenerateEnrollmentLink
	}
	if h, ok := s.(interface {
		WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIReply, error)
	}); ok {
		ns.WhoAmI = h.WhoAmI
	}
	return ns
}

// UnstableUserService is the service API for User service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableUserService interface {
	// Edit adds a new user if the ID is 0, or updates an existing user.
	Edit(context.Context, *EditUserRequest) (*EditUserReply, error)
	// GenerateEnrollmentLink generates an enrollment token for the user.
	GenerateEnrollmentLink(context.Context, *GenerateEnrollmentLinkRequest) (*GenerateEnrollmentLinkReply, error)
	// WhoAmI returns the user object associated with the current session.  When
	// called without a session, a null current user is returned rather than an
	// error.
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIReply, error)
}

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionClient interface {
	// AuthorizeHTTP authorizes an incoming HTTP request, returning a
	// request-scoped bearer token to authenticate the request to upstream
	// systems.  A "deny" authorization decision will not be transmitted as a
	// gRPC error; a gRPC error like PermissionDenied means that the caller is
	// not authorized to call AuthorizeHTTP, not that the HTTP request is
	// unauthorized.  (Corollary: an OK response code does not mean the request
	// is authorized.)
	AuthorizeHTTP(ctx context.Context, in *AuthorizeHTTPRequest, opts ...grpc.CallOption) (*AuthorizeHTTPReply, error)
}

type sessionClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionClient(cc grpc.ClientConnInterface) SessionClient {
	return &sessionClient{cc}
}

var sessionAuthorizeHTTPStreamDesc = &grpc.StreamDesc{
	StreamName: "AuthorizeHTTP",
}

func (c *sessionClient) AuthorizeHTTP(ctx context.Context, in *AuthorizeHTTPRequest, opts ...grpc.CallOption) (*AuthorizeHTTPReply, error) {
	out := new(AuthorizeHTTPReply)
	err := c.cc.Invoke(ctx, "/jsso.Session/AuthorizeHTTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionService is the service API for Session service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSessionService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SessionService struct {
	// AuthorizeHTTP authorizes an incoming HTTP request, returning a
	// request-scoped bearer token to authenticate the request to upstream
	// systems.  A "deny" authorization decision will not be transmitted as a
	// gRPC error; a gRPC error like PermissionDenied means that the caller is
	// not authorized to call AuthorizeHTTP, not that the HTTP request is
	// unauthorized.  (Corollary: an OK response code does not mean the request
	// is authorized.)
	AuthorizeHTTP func(context.Context, *AuthorizeHTTPRequest) (*AuthorizeHTTPReply, error)
}

func (s *SessionService) authorizeHTTP(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeHTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.AuthorizeHTTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.Session/AuthorizeHTTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.AuthorizeHTTP(ctx, req.(*AuthorizeHTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSessionService registers a service implementation with a gRPC server.
func RegisterSessionService(s grpc.ServiceRegistrar, srv *SessionService) {
	srvCopy := *srv
	if srvCopy.AuthorizeHTTP == nil {
		srvCopy.AuthorizeHTTP = func(context.Context, *AuthorizeHTTPRequest) (*AuthorizeHTTPReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method AuthorizeHTTP not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "jsso.Session",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "AuthorizeHTTP",
				Handler:    srvCopy.authorizeHTTP,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "jsso.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSessionService creates a new SessionService containing the
// implemented methods of the Session service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSessionService(s interface{}) *SessionService {
	ns := &SessionService{}
	if h, ok := s.(interface {
		AuthorizeHTTP(context.Context, *AuthorizeHTTPRequest) (*AuthorizeHTTPReply, error)
	}); ok {
		ns.AuthorizeHTTP = h.AuthorizeHTTP
	}
	return ns
}

// UnstableSessionService is the service API for Session service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSessionService interface {
	// AuthorizeHTTP authorizes an incoming HTTP request, returning a
	// request-scoped bearer token to authenticate the request to upstream
	// systems.  A "deny" authorization decision will not be transmitted as a
	// gRPC error; a gRPC error like PermissionDenied means that the caller is
	// not authorized to call AuthorizeHTTP, not that the HTTP request is
	// unauthorized.  (Corollary: an OK response code does not mean the request
	// is authorized.)
	AuthorizeHTTP(context.Context, *AuthorizeHTTPRequest) (*AuthorizeHTTPReply, error)
}

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	Start(ctx context.Context, in *StartLoginRequest, opts ...grpc.CallOption) (*StartLoginReply, error)
	Finish(ctx context.Context, in *FinishLoginRequest, opts ...grpc.CallOption) (*FinishLoginReply, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

var loginStartStreamDesc = &grpc.StreamDesc{
	StreamName: "Start",
}

func (c *loginClient) Start(ctx context.Context, in *StartLoginRequest, opts ...grpc.CallOption) (*StartLoginReply, error) {
	out := new(StartLoginReply)
	err := c.cc.Invoke(ctx, "/jsso.Login/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var loginFinishStreamDesc = &grpc.StreamDesc{
	StreamName: "Finish",
}

func (c *loginClient) Finish(ctx context.Context, in *FinishLoginRequest, opts ...grpc.CallOption) (*FinishLoginReply, error) {
	out := new(FinishLoginReply)
	err := c.cc.Invoke(ctx, "/jsso.Login/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginService is the service API for Login service.
// Fields should be assigned to their respective handler implementations only before
// RegisterLoginService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type LoginService struct {
	Start  func(context.Context, *StartLoginRequest) (*StartLoginReply, error)
	Finish func(context.Context, *FinishLoginRequest) (*FinishLoginReply, error)
}

func (s *LoginService) start(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.Login/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Start(ctx, req.(*StartLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *LoginService) finish(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.Login/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Finish(ctx, req.(*FinishLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterLoginService registers a service implementation with a gRPC server.
func RegisterLoginService(s grpc.ServiceRegistrar, srv *LoginService) {
	srvCopy := *srv
	if srvCopy.Start == nil {
		srvCopy.Start = func(context.Context, *StartLoginRequest) (*StartLoginReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
		}
	}
	if srvCopy.Finish == nil {
		srvCopy.Finish = func(context.Context, *FinishLoginRequest) (*FinishLoginReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "jsso.Login",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Start",
				Handler:    srvCopy.start,
			},
			{
				MethodName: "Finish",
				Handler:    srvCopy.finish,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "jsso.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewLoginService creates a new LoginService containing the
// implemented methods of the Login service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewLoginService(s interface{}) *LoginService {
	ns := &LoginService{}
	if h, ok := s.(interface {
		Start(context.Context, *StartLoginRequest) (*StartLoginReply, error)
	}); ok {
		ns.Start = h.Start
	}
	if h, ok := s.(interface {
		Finish(context.Context, *FinishLoginRequest) (*FinishLoginReply, error)
	}); ok {
		ns.Finish = h.Finish
	}
	return ns
}

// UnstableLoginService is the service API for Login service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableLoginService interface {
	Start(context.Context, *StartLoginRequest) (*StartLoginReply, error)
	Finish(context.Context, *FinishLoginRequest) (*FinishLoginReply, error)
}

// EnrollmentClient is the client API for Enrollment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnrollmentClient interface {
	Start(ctx context.Context, in *StartEnrollmentRequest, opts ...grpc.CallOption) (*StartEnrollmentReply, error)
	Finish(ctx context.Context, in *FinishEnrollmentRequest, opts ...grpc.CallOption) (*FinishEnrollmentReply, error)
}

type enrollmentClient struct {
	cc grpc.ClientConnInterface
}

func NewEnrollmentClient(cc grpc.ClientConnInterface) EnrollmentClient {
	return &enrollmentClient{cc}
}

var enrollmentStartStreamDesc = &grpc.StreamDesc{
	StreamName: "Start",
}

func (c *enrollmentClient) Start(ctx context.Context, in *StartEnrollmentRequest, opts ...grpc.CallOption) (*StartEnrollmentReply, error) {
	out := new(StartEnrollmentReply)
	err := c.cc.Invoke(ctx, "/jsso.Enrollment/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var enrollmentFinishStreamDesc = &grpc.StreamDesc{
	StreamName: "Finish",
}

func (c *enrollmentClient) Finish(ctx context.Context, in *FinishEnrollmentRequest, opts ...grpc.CallOption) (*FinishEnrollmentReply, error) {
	out := new(FinishEnrollmentReply)
	err := c.cc.Invoke(ctx, "/jsso.Enrollment/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnrollmentService is the service API for Enrollment service.
// Fields should be assigned to their respective handler implementations only before
// RegisterEnrollmentService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type EnrollmentService struct {
	Start  func(context.Context, *StartEnrollmentRequest) (*StartEnrollmentReply, error)
	Finish func(context.Context, *FinishEnrollmentRequest) (*FinishEnrollmentReply, error)
}

func (s *EnrollmentService) start(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartEnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.Enrollment/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Start(ctx, req.(*StartEnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *EnrollmentService) finish(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishEnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/jsso.Enrollment/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Finish(ctx, req.(*FinishEnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterEnrollmentService registers a service implementation with a gRPC server.
func RegisterEnrollmentService(s grpc.ServiceRegistrar, srv *EnrollmentService) {
	srvCopy := *srv
	if srvCopy.Start == nil {
		srvCopy.Start = func(context.Context, *StartEnrollmentRequest) (*StartEnrollmentReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
		}
	}
	if srvCopy.Finish == nil {
		srvCopy.Finish = func(context.Context, *FinishEnrollmentRequest) (*FinishEnrollmentReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "jsso.Enrollment",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Start",
				Handler:    srvCopy.start,
			},
			{
				MethodName: "Finish",
				Handler:    srvCopy.finish,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "jsso.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewEnrollmentService creates a new EnrollmentService containing the
// implemented methods of the Enrollment service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewEnrollmentService(s interface{}) *EnrollmentService {
	ns := &EnrollmentService{}
	if h, ok := s.(interface {
		Start(context.Context, *StartEnrollmentRequest) (*StartEnrollmentReply, error)
	}); ok {
		ns.Start = h.Start
	}
	if h, ok := s.(interface {
		Finish(context.Context, *FinishEnrollmentRequest) (*FinishEnrollmentReply, error)
	}); ok {
		ns.Finish = h.Finish
	}
	return ns
}

// UnstableEnrollmentService is the service API for Enrollment service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableEnrollmentService interface {
	Start(context.Context, *StartEnrollmentRequest) (*StartEnrollmentReply, error)
	Finish(context.Context, *FinishEnrollmentRequest) (*FinishEnrollmentReply, error)
}
