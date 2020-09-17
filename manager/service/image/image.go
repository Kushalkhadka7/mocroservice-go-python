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
)

const maxImageSize = 1 << 20

// CreateImageStorage image storage service interface.
type CreateImageStorage interface {
	Save(stream laptop.LaptopService_UploadLaptopImageServer) (*laptop.UploadImageResponse, error)
}

// ImageService initializes new image service.
type ImageService struct {
	laptopservice.LaptopService
	ImageStore laptopstore.ImageStore
}

// NewImageService creates new image service.
func NewImageService(store laptopstore.ImageStore) CreateImageStorage {
	return &ImageService{ImageStore: store}
}

// Save image to imaage store.
func (service *ImageService) Save(stream laptop.LaptopService_UploadLaptopImageServer) (*laptop.UploadImageResponse, error) {
	fmt.Println("hello i am called here")

	req, err := stream.Recv()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("Request received with image type : %s and laptop id :%s", imageType, laptopID)

	result, err := service.LaptopService.FindLaptop(context.Background(), "hello")
	fmt.Println(result)
	if err != nil {
		return nil, err
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		fmt.Println("Starting to receive chunks of image data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("No more image data")
			break
		}
		if err != nil {
			return nil, err
		}

		chunk := req.GetChunkData()
		imageSize += len(chunk)
		if imageSize > maxImageSize {
			return nil, fmt.Errorf("Image size exceeds. %d", maxImageSize)
		}

		_, err = imageData.Write(chunk)
		if err != nil {
			return nil, err
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
