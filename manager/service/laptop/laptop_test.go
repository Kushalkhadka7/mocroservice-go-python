package laptopservice_test

import (
	"manager/service/laptop"
	"testing"
	"context"
	laptop "manager/pb"
	"manager/sample"
)

type MockLaptopStore struct{}


func (store *MockLaptopStore)Save(ctx context.Context, laptop *laptop.Laptop) (string, error){
	return "",nil
}
func (store *MockLaptopStore)FetchAll(ctx context.Context) (*laptop.FetchLaptopResposne, error){
	return nil,nil
}
func (store *MockLaptopStore)FindLaptop(ctx context.Context, laptopID string) (string, error){
	return "",nil
}
func (store *MockLaptopStore)UpdateLaptop(ctx context.Context, laptopID string, imageID string) error{
	return nil
}


var ctx = context.Background()

func TestSaveLaptop(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
	}{
		{
			name: "Kushal",
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			laptopService := laptopservice.New(MockLaptopStore{})
			res,err:=laptopService.SaveLaptop(ctx,sample.NewLaptop())
		})
	}
}
