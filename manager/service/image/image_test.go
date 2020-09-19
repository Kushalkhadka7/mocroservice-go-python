package imageservice

import (
	"bytes"
	imagestore "manager/store"
	"testing"
)

type MockImageStore struct{}

func (service *MockImageStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (*imagestore.ImageInfo, error) {
	return nil, nil
}

var imageService = NewImageService(&MockImageStore{}, nil)

func TestSave(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
	}{
		{
			name: "kushal",
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Cannot mock stream.
		})
	}
}
