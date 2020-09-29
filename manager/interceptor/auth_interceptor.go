package interceptor

import (
	"context"
	pb "manager/pb"
	"manager/util"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// AuthInterceptor initializes new auth interceptor.
type AuthInterceptor struct{}

// NewAuthInterceptor creates new auth interceptor.
func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

// Unary server interceptor.
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, util.Error(codes.Aborted, "No incoming context.", err)
		}

		values := md["authorization"]
		if len(values) == 0 {
			return nil, util.Error(codes.Internal, "Unauthorized, No authorization mechanism provided.", err)
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

		var result *pb.VerifyUserTokenResponse

		context := context.WithValue(ctx, result, res)

		return handler(context, req)
	}

}

// Stream creates new stream interceptor.
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		serv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {

		return handler(serv, stream)
	}

}

// createAuthClient creates new auth client.
func createAuthClient() (pb.AuthServiceClient, error) {
	serverConn, err := grpc.Dial("auth:6001",
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, util.Error(codes.Unauthenticated, "Unable to create auth client.", err)
	}

	authCleint := pb.NewAuthServiceClient(serverConn)

	return authCleint, nil
}
