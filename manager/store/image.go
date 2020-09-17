package laptopstore

import (
	"bytes"
	"fmt"
	"os"

	guuid "github.com/google/uuid"
)

//ImageStore interface.
type ImageStore interface {
	Save(laptopID string, imageType string, imageData bytes.Buffer) (*ImageInfo, error)
}

// DiskImageStore implements storage server.
type DiskImageStore struct {
	imageFolder string
	image       map[string]*ImageInfo
}

// ImageInfo image metadata.
type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
	ImageID string
}

// NewDiskImageStore new server.
func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imageFolder: imageFolder,
		image:       make(map[string]*ImageInfo),
	}
}

// Save image file.
func (store *DiskImageStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (*ImageInfo, error) {
	imageID := guuid.New()

	imagePath := fmt.Sprintf("%s/%s.%s", store.imageFolder, imageID, imageType)
	imageFile, err := os.Create(imagePath)
	if err != nil {
		return nil, err
	}

	_, err = imageData.WriteTo(imageFile)
	if err != nil {
		return nil, err
	}


	return &ImageInfo{
		LaptopID: laptopID,
		Type:     imageType,
		Path:     imagePath,
		ImageID: imageID.String(),
	}, nil
}
