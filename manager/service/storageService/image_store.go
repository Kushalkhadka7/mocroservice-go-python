package storageservice

import (
	"bytes"
	"fmt"
	"os"

	guuid "github.com/google/uuid"
)

// ImageStore implements storage server.
type ImageStore struct {
	imageFolder string
	image       map[string]*ImageInfo
}

// ImageInfo image metadata.
type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
}

//Store interface.
type Store interface {
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error)
}

// NewStorageServer new server.
func NewStorageServer(imageFolder string) *ImageStore {
	return &ImageStore{
		imageFolder: imageFolder,
		image:       make(map[string]*ImageInfo),
	}
}

// Save image file.
func (store *ImageStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error) {
	imageID := guuid.New()

	imagePath := fmt.Sprintf("%s/%s.%s", store.imageFolder, imageID, imageType)
	imageFile, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}

	_, err = imageData.WriteTo(imageFile)
	if err != nil {
		return "", err
	}

	return imageID.String(), nil
}
