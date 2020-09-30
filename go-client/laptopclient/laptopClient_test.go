package laptopclient

import (
	pb "client/pb"
	"client/sample"
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
)

type laptopServiceClient struct{}

func (c *laptopServiceClient) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	return nil
}
func (c *laptopServiceClient) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, nil
}

var service = pb.NewLaptopServiceClient(&laptopServiceClient{})

type mockLaptopClient struct{}

var mockClinet = NewLapotpClient(&mockLaptopClient{})

func (c *mockLaptopClient) CreateLaptop(ctx context.Context, in *pb.CreateLaptopRequest, opts ...grpc.CallOption) (*pb.CreateLaptopResponse, error) {
	res := &pb.CreateLaptopResponse{
		Laptop: sample.NewLaptop(),
	}
	return res, nil
}
func (c *mockLaptopClient) FetchAllLaptops(ctx context.Context, in *pb.Void, opts ...grpc.CallOption) (pb.LaptopService_FetchAllLaptopsClient, error) {
	return nil, nil
}
func (c *mockLaptopClient) UploadLaptopImage(ctx context.Context, opts ...grpc.CallOption) (pb.LaptopService_UploadLaptopImageClient, error) {
	return nil, nil
}
func (c *mockLaptopClient) SayHello(ctx context.Context, in *pb.Hello, opts ...grpc.CallOption) (*pb.CreateLaptopResponse, error) {
	return nil, nil
}

func TestCreateLaptop(t *testing.T) {
	mockClinet.CreateNewLaptop()

}

func BenchmarkCreateLaptop(b *testing.B) {

	for n := 1; n <= 10; n *= 2 {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mockClinet.CreateNewLaptop()
			}
		})
	}
}
