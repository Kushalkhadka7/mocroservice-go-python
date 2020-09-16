package server_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	laptop "manager/pb"
	"manager/server"
	"testing"
	"manager/sample"
)

type MockLaptopService struct{}

func (service MockLaptopService) SaveLaptop(ctx context.Context, laptop *laptop.Laptop) (string, error) {
	return "Hello", nil
}

var ctx = context.Background()
var laptopServer = server.NewLaptopServer(MockLaptopService{}, nil)


func TestLaptopServer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		req  *laptop.Hello
		code codes.Code
	}{
		{
			name: "Success",
			req:  &laptop.Hello{Hello: "Success"},
			code: codes.OK,
		},
		{
			name: "Error",
			req:  &laptop.Hello{Hello: "Error"},
			code: codes.Unknown,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := laptopServer.SayHello(ctx, tc.req)
			if err != nil {
				panic(err)
			}

			require.NoError(t, err)
			require.NotNil(t, res)
			require.NotEmpty(t, res.Laptop.Id)
		})
	}

}

func TestCreateLapotp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		req *laptop.CreateLaptopRequest
		code codes.Code
	}{
		{
			name: "Success with laptop id.",
			req: &laptop.CreateLaptopRequest{Laptop:sample.NewLaptop()},
			code: codes.OK,
		},
		{
			name: "Failed",
			req: &laptop.CreateLaptopRequest{Laptop:sample.NewLaptop()},
			code: codes.Internal,
		},
	}

	for i:= range testCases{
		tc := testCases[i]

		t.Run(tc.name,func(t *testing.T){
			t.Parallel()

			res,err:= laptopServer.CreateLaptop(ctx,tc.req)
			if err != nil {
				require.Error(t,err)
				t.Fatalf("Error, %v",err)
			}
			require.NoError(t,err)
			require.NotNil(t,res)
			require.NotEmpty(t,res)
			require.NotEmpty(t,res.Laptop)
			require.NotEmpty(t,res.Laptop.Id)
		})
	}
}
