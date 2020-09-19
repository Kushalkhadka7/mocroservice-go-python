package laptopstore

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	laptop "manager/pb"
	"os"
	"testing"
)

const maxImageSize = 1 << 20

type MockDiskImageStore struct{}

var mockDiskImageStore = NewDiskImageStore("../testimg")
var imagePath = "../tmp/laptop.jpg"

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

			imageFile, err := os.Open(imagePath)
			if err != nil {
				log.Fatal("cannot open image file :", err)
			}
			defer imageFile.Close()

			reader := bufio.NewReader(imageFile)
			buffer := make([]byte, 1024)
			imageSize := 0
			imageData := bytes.Buffer{}

			for {
				data, err := reader.Read(buffer)
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal("Unable to read image data.", err)
				}

				req := &laptop.UploadImageRequest{
					Data: &laptop.UploadImageRequest_ChunkData{
						ChunkData: buffer[:data],
					},
				}

				chunk := req.GetChunkData()
				imageSize += len(chunk)
				if imageSize > maxImageSize {
					fmt.Printf("Image size exceeddatas. %d", maxImageSize)
				}

				_, err = imageData.Write(chunk)
				if err != nil {
					fmt.Printf("Image size exceeddatas. %d", maxImageSize)
				}

			}

			res, err := mockDiskImageStore.Save("abc", "jpg", imageData)
			if err != nil {
				fmt.Printf("Unable to save image %d", err)
			}

			fmt.Println(res)
		})
	}
}
