package imageservice

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	laptop "manager/pb"
	laptopservice "manager/service/laptop"
	laptopstore "manager/store"
	"manager/util"

	"google.golang.org/grpc/codes"
)

const maxImageSize = 1 << 20

// CreateImageStorage image storage service interface.
type CreateImageStorage interface {
	Save(stream laptop.LaptopService_UploadLaptopImageServer) (*laptop.UploadImageResponse, error)
}

// ImageService initializes new image service.
type ImageService struct {
	LaptopService laptopservice.LaptopService
	ImageStore    laptopstore.ImageStore
}

// NewImageService creates new image service.
func NewImageService(store laptopstore.ImageStore, laptopService laptopservice.LaptopService) CreateImageStorage {
	return &ImageService{ImageStore: store, LaptopService: laptopService}
}

// Save image to imaage store.
func (service *ImageService) Save(stream laptop.LaptopService_UploadLaptopImageServer) (*laptop.UploadImageResponse, error) {

	req, err := stream.Recv()
	if err != nil {
		return nil, util.Error(codes.Internal, "Unable to receive data", err)
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("Request received with image type : %s and laptop id :%s", imageType, laptopID)

	result, err := service.LaptopService.FindLaptop(context.Background(), laptopID)
	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		fmt.Printf("Starting to receive chunks of image data: %v \n", imageSize)

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("No more image data")
			break
		}
		if err != nil {
			return nil, util.Error(codes.Internal, "Unable to receive data", err)
		}

		chunk := req.GetChunkData()
		imageSize += len(chunk)
		if imageSize > maxImageSize {
			return nil, fmt.Errorf("Image size exceeddatas. %d", maxImageSize)
		}

		_, err = imageData.Write(chunk)
		if err != nil {
			return nil, util.Error(codes.Internal, "Unable write data to file", err)
		}
	}

	imageInfo, err := service.ImageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		return nil, err
	}

	err = service.LaptopService.UpdateLaptop(context.Background(), laptopID, imageInfo.ImageID)
	if err != nil {
		return nil, err
	}

	res := &laptop.UploadImageResponse{
		ImageId:   imageInfo.ImageID,
		LaptopId:  imageInfo.LaptopID,
		ImageSize: uint32(imageSize),
	}

	return res, nil

}
