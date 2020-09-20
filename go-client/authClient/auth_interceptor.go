package authclient

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type AuthInterceptor struct {
	authClient  *AuthClient
	accessToken string
}

func NewAuthInterceptor(authClient *AuthClient, refreshDuration time.Duration) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		authClient: authClient,
	}

	err := interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {

	return func(
		ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("------->unary interceptor: %s", method)

		return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
	}

}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(duration time.Duration) error {

	err := interceptor.refreshToken()
	if err != nil {
		return err
	}

	go func() {
		wait := duration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()
			if err != nil {
				wait = time.Second
			} else {
				wait = duration
			}
		}

	}()

	return nil

}

func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}

	interceptor.accessToken = accessToken
	log.Printf("Token refreshed: %v", accessToken)

	return nil
}
