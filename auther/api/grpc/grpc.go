package grpc

import (
	"auther/api"
	"auther/models"
	context "context"
	"errors"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type auther struct {
	l  models.Logger
	jw models.JwtSigner
}

var (
	ErrTimeout = errors.New("call has timed out")
)

func InitGRPCServer(socket string, bp api.Backpack) {
	ks := bp.Sigterm
	l := bp.Logger
	jw := bp.Signer

	defer func() { l.Info("GRPC Shutting down"); ks.WaitGroup.Done() }()

	lis, err := net.Listen("tcp", socket)
	if err != nil {
		l.Error("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	RegisterAutherServer(s, &auther{l, jw})
	reflection.Register(s)

	go func() {
		l.Info("GRPC server started - server listening at ", lis.Addr())
		if err := s.Serve(lis); err != nil {
			l.Error("failed to serve", err)
		}
	}()

	sig := <-ks.KillChannel
	ks.KillChannel <- sig
	s.GracefulStop()
}

func (s *auther) ValidateJWT(ctx context.Context, t *JWTToken) (*JWTClaims, error) {
	go func(done <-chan struct{}) {
		<-done
		s.l.Info("client connection has gone away, request got cancelled")
	}(ctx.Done())

	token := t.GetToken()
	s.l.Info("Got token to validate", "jtw", token)

	claims, err := s.jw.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf("error validating the token; %w", err)
	}
	s.l.Info("Got token was valid!", "user.mail", claims.Email)

	return &JWTClaims{Name: claims.Name, Email: claims.Email, Expiry: timestamppb.New(time.Unix(claims.ExpiresAt, 0))}, nil
}

func (s *auther) mustEmbedUnimplementedAutherServer() {}
