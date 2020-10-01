package server_test

import (
	"context"
	laptop "manager/pb"
	"manager/sample"
	"manager/server"
	"manager/util"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

type MockLaptopService struct{}

// Mock save laptop funciton.
func (service MockLaptopService) SaveLaptop(ctx context.Context, laptop *laptop.Laptop) (string, error) {
	oid, err := util.GenerateOID("new Laptop")
	if err != nil {
		return "", err
	}

	return oid, nil
}
func (service *MockLaptopService) FetchAllLaptops(ctx context.Context) ([]*laptop.Laptop, error) {
	return nil, nil
}
func (service *MockLaptopService) FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error) {
	return nil, nil
}
func (service *MockLaptopService) UpdateLaptop(ctx context.Context, laptopID string, imageID string) error {
	return nil
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
			require.NotEmpty(t, res.Laptop.Brand)
		})
	}

}

func TestCreateLapotp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		req  *laptop.CreateLaptopRequest
		code codes.Code
	}{
		{
			name: "Success",
			req:  &laptop.CreateLaptopRequest{Laptop: sample.NewLaptop()},
			code: codes.OK,
		},
		{
			name: "Failed",
			req:  &laptop.CreateLaptopRequest{Laptop: sample.NewLaptop()},
			code: codes.Internal,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := laptopServer.CreateLaptop(ctx, tc.req)
			require.NoError(t, err)
			require.NotNil(t, res)
			require.NotEmpty(t, res)
			require.NotEmpty(t, res.Laptop)
			require.NotEmpty(t, res.Laptop.Brand)
		})
	}
}
