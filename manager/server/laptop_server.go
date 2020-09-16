package server

import (
	"context"
	"fmt"

	laptop "manager/pb"
	"manager/sample"
	laptopservice "manager/service/laptop"
	storageservice "manager/service/storage"
)

const maxImageSize = 1 << 20

// LaptopServer struct.
type LaptopServer struct {
	LaptopService laptopservice.LaptopService
	ImageService  storageservice.Store
}

// NewLaptopServer create new laptop store.
func NewLaptopServer(laptopStore laptopservice.LaptopService, imageStore storageservice.Store) *LaptopServer {
	return &LaptopServer{LaptopService: laptopStore, ImageService: imageStore}
}

// CreateLaptop creates new laptop.
func (service *LaptopServer) CreateLaptop(ctx context.Context, req *laptop.CreateLaptopRequest) (*laptop.CreateLaptopResponse, error) {
	sample := sample.NewLaptop()

	result, err := service.LaptopService.SaveLaptop(context.Background(), sample)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sample.Id = result

	return &laptop.CreateLaptopResponse{Laptop: sample}, nil
}

// SayHello hello func.
func (service *LaptopServer) SayHello(ctx context.Context, req *laptop.Hello) (*laptop.CreateLaptopResponse, error) {
	sample := sample.NewLaptop()

	result, err := service.LaptopService.SaveLaptop(context.Background(), sample)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sample.Id = result

	return &laptop.CreateLaptopResponse{Laptop: sample}, nil
}

// FetchAllLaptops fetch all laptops form the data base.
func (service *LaptopServer) FetchAllLaptops(void *laptop.Void, stream laptop.LaptopService_FetchAllLaptopsServer) error {
	fmt.Println("hello i am called.")

	result, err := service.LaptopService.FetchAll(stream.Context())
	if err != nil {
		return err
	}
	
	if err := stream.Send(result); err != nil {
		return err
	}

	fmt.Printf("Successfully send laptop with id: %s", result.Laptop.GetId())

	return nil

}

//
// UploadLaptopImage upload image to server.
func (service *LaptopServer) UploadLaptopImage(stream laptop.LaptopService_UploadLaptopImageServer) error {
	// 	fmt.Println("hello i am called here")
	// //
	// 	req, err := stream.Recv()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return err
	// 	}
	// //
	// 	laptopID := req.GetInfo().GetLaptopId()
	// 	imageType := req.GetInfo().GetImageType()
	// 	log.Printf("Request received with image type : %s and laptop id :%s", imageType, laptopID)
	// //
	// 	result, err := service.LaptopStore.FindLaptop(context.Background(), laptopID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if result == "" {
	// 		return fmt.Errorf("No data found with id %s", laptopID)
	// 	}
	// //
	// 	imageData := bytes.Buffer{}
	// 	imageSize := 0
	// //
	// 	for {
	// 		fmt.Println("Starting to receive chunks of image data")
	// //
	// 		req, err := stream.Recv()
	// 		if err == io.EOF {
	// 			log.Println("No more image data")
	// 			break
	// 		}
	// 		if err != nil {
	// 			return err
	// 		}
	// //
	// 		chunk := req.GetChunkData()
	// 		imageSize += len(chunk)
	// 		if imageSize > maxImageSize {
	// 			return fmt.Errorf("Image size exceeds. %d", maxImageSize)
	// 		}
	// //
	// 		_, err = imageData.Write(chunk)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// //
	// 	imageID, err := service.ImageStore.Save(laptopID, imageType, imageData)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if imageID == "" {
	// 		return fmt.Errorf("No image is saved %s", "err")
	// 	}
	// //
	// 	err = service.LaptopStore.UpdateLaptop(context.Background(), laptopID, imageID)
	// 	if err != nil {
	// 		return err
	// 	}
	// //
	// 	res := &laptop.UploadImageResponse{
	// 		ImageId:   imageID,
	// 		LaptopId:  laptopID,
	// 		ImageSize: uint32(imageSize),
	// 	}
	// //
	// 	err = stream.SendAndClose(res)
	// 	if err != nil {
	// 		return err
	// 	}
	// //
	return nil
}

//
