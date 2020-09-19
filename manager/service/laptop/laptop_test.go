package laptopservice_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	laptop "manager/pb"
	"manager/sample"
	"manager/service/laptop"
	"testing"
)

var ctx = context.Background()

type MockLaptopStore struct{}

func (store *MockLaptopStore) FetchAll(ctx context.Context) ([]*laptop.Laptop, error) {
	return nil,nil
}

func (store *MockLaptopStore) Save(ctx context.Context, laptop *laptop.Laptop) (string, error) {
	return "", nil
}

func (store *MockLaptopStore) FindLaptop(ctx context.Context, laptopID string) (*laptop.Laptop, error) {
	return nil, nil
}

func (store *MockLaptopStore) UpdateLaptop(ctx context.Context, laptopID string, imageID string) (*mongo.UpdateResult, error) {
	return nil, nil
}

func TestSaveLaptop(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		req  *laptop.Laptop
		code codes.Code
	}{
		{
			name: "Success",
			req:  sample.NewLaptop(),
			code: codes.OK,
		},
		{
			name: "Failed",
			req:  sample.NewLaptop(),
			code: codes.Internal,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			laptopService := laptopservice.New(MockLaptopStore{})
			res, err := laptopService.SaveLaptop(ctx, tc.req)
			require.NoError(t, err)
			require.NotNil(t, res)
			require.NotEmpty(t, res)
		})
	}
}
