package laptopservice

import (
	"context"
	"fmt"
	laptop "manager/pb"
	"manager/sample"
	"testing"
)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		error string
	}{
		{
			name:  "Test1",
			error: "Not found",
		},
		{
			name:  "Test2",
			error: "Found",
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := NewLaptopService(NewLaptopStore(db))
			req := &laptop.CreateLaptopRequest{Laptop: sample.NewLaptop()}

			result, err := service.CreateLaptop(context.Background(), req)
			if err != nil {
				panic(err)
			}

			fmt.Println(result)

		})
	}
}
