package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jrockway/monorepo/jsso2/pkg/client"
	"github.com/jrockway/monorepo/jsso2/pkg/jssopb"
	"github.com/jrockway/monorepo/jsso2/pkg/jtesting"
	"github.com/jrockway/monorepo/jsso2/pkg/sessions"
	"github.com/jrockway/monorepo/jsso2/pkg/testserver"
	"github.com/jrockway/monorepo/jsso2/pkg/types"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestUsers(t *testing.T) {
	jsonOutput = true
	s := testserver.New()
	r := &jtesting.R{Logger: true, Database: true}
	defaultCreds := &client.Credentials{Root: "root"}
	s.ToR(r)
	s.Credentials = &client.Credentials{}
	jtesting.Run(t, "users", *r, func(t *testing.T, e *jtesting.E) {
		testData := []struct {
			name         string
			args         []string
			wantFail     bool
			wantOutProto proto.Message
			cmpopts      []cmp.Option
			wantErr      string
			creds        *client.Credentials
		}{
			{
				name: "add",
				args: []string{"users", "add", "test"},
				wantOutProto: &jssopb.EditUserReply{
					User: &types.User{
						Id:       1,
						Username: "test",
					},
				},
				wantErr: "OK\n",
			},
			{
				name:     "add again",
				args:     []string{"users", "add", "test"},
				wantFail: true,
			},
			{
				name: "enroll",
				args: []string{"users", "enroll", "--username=test"},
				wantOutProto: &jssopb.GenerateEnrollmentLinkReply{
					Token: sessions.ToBase64(&types.Session{Id: make([]byte, 64)}),
					Url:   "http://jsso.example.com/#/enroll/",
				},
				cmpopts: []cmp.Option{
					protocmp.FilterField(
						&jssopb.GenerateEnrollmentLinkReply{},
						"token",
						cmp.Comparer(func(x, y string) bool {
							return len(x) == len(y)
						}),
					),
					protocmp.FilterField(
						&jssopb.GenerateEnrollmentLinkReply{},
						"url",
						cmp.Comparer(func(x, y string) bool {
							if len(x) > len(y) {
								return strings.HasPrefix(x, y)
							}
							return strings.HasPrefix(y, x)
						}),
					),
				},
				wantErr: "OK\n",
			},
			{
				name:     "enroll with invalid user",
				args:     []string{"users", "enroll", "--username=invalid"},
				wantFail: true,
			},
			{
				name: "whoami",
				args: []string{"users", "whoami"},
				wantOutProto: &jssopb.WhoAmIReply{
					User: sessions.Root().GetUser(),
				},
				wantErr: "OK\n",
			},
			{
				name:         "whoami as anonymous user",
				creds:        &client.Credentials{},
				args:         []string{"users", "whoami"},
				wantOutProto: &jssopb.WhoAmIReply{},
				wantErr:      "OK\n",
			},
			{
				name: "whoami with invalid session",
				creds: &client.Credentials{
					Token: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAa",
				},
				args:         []string{"users", "whoami"},
				wantOutProto: &jssopb.WhoAmIReply{},
				wantErr:      "OK\n",
			},
		}

		clientset = client.FromCC(e.ClientConn)
		noClose = true

		for _, test := range testData {
			t.Run(test.name, func(t *testing.T) {
				rootCmd.SetArgs(test.args)
				out := new(bytes.Buffer)
				err := new(bytes.Buffer)
				rootCmd.SetOut(out)
				rootCmd.SetErr(err)
				defer rootCmd.SetOut(os.Stderr)
				defer rootCmd.SetErr(os.Stderr)

				creds := test.creds
				if creds == nil {
					creds = defaultCreds
				}
				s.Credentials.Root = creds.Root
				s.Credentials.Token = creds.Token
				s.Credentials.Bearer = creds.Bearer

				if err := rootCmd.ExecuteContext(e.Context); !test.wantFail && err != nil {
					t.Fatalf("execute: %v", err)
				} else if test.wantFail && err == nil {
					t.Error("execute: expected error")
				}

				if want := test.wantOutProto; want != nil {
					got := proto.Clone(want)
					proto.Reset(got)
					if err := protojson.Unmarshal(out.Bytes(), got); err != nil {
						t.Fatalf("parse result: %v", err)
					}

					opts := []cmp.Option{protocmp.Transform()}
					opts = append(opts, test.cmpopts...)
					if diff := cmp.Diff(got, want, opts...); diff != "" {
						t.Errorf("output: %s", diff)
					}
				}

				if want := test.wantErr; want != "" {
					got := err.String()
					if got != want {
						t.Errorf("status output:\n  got: %v\n want: %v", got, want)
					}
				}
			})
		}
	})
}
