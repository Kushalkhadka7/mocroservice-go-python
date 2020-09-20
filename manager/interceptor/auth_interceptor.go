package interceptor

import (
	"context"
	"fmt"
	pb "manager/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type AuthInterceptor struct {
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, err
		}

		values := md["authorization"]
		if len(values) == 0 {
			return nil, status.Error(codes.PermissionDenied, "Unauthorized")
		}

		accessToken := values[0]

		authReq := &pb.VerifyUserTokenRequest{
			AccessToken: accessToken,
		}

		authCleint, err := createAuthClient()
		if err != nil {
			return nil, err
		}

		res, err := authCleint.VerifyUser(ctx, authReq)
		if err != nil {
			return nil, err

		}

		fmt.Println("from interceprot")
		fmt.Println(res)

		var result *pb.VerifyUserTokenResponse

		context := context.WithValue(ctx, result, res)

		return handler(context, req)
	}

}

func createAuthClient() (pb.AuthServiceClient, error) {
	serverConn, err := grpc.Dial("localhost:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	authCleint := pb.NewAuthServiceClient(serverConn)

	return authCleint, nil
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		serv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		fmt.Println("hello world i am called again.")

		return handler(serv, stream)
	}

}
