package authclient

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// AuthInterceptor initializes new auth interceptor .
type AuthInterceptor struct {
	authClient  *AuthClient
	accessToken string
}

// NewAuthInterceptor creates new interceprtor.
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

// Unary creates new unary interceptor.
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
		log.Printf("-------unary interceptor: %s -------", method)

		return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
	}

}

// attachToken to the context of the request going to the server.
func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}

// scheduleRefreshToken refresh the token in given time interval.
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

// refreshToken refreshes the token.b
func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}

	interceptor.accessToken = accessToken
	log.Printf("Token refreshed: %v", accessToken)

	return nil
}
