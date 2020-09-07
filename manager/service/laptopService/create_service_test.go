package laptopservice

import (
	"context"
	"fmt"
	laptop "manager/pb"
	"manager/sample"
	"testing"
)

var store = NewLaptopStore(db)
var service = NewLaptopService(store)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		error string
		req   *laptop.CreateLaptopRequest
	}{
		{
			name:  "Test1",
			error: "Not found",
			req:   &laptop.CreateLaptopRequest{Laptop: sample.NewLaptop()},
		},
		{
			name:  "Test2",
			error: "Found",
			req:   &laptop.CreateLaptopRequest{Laptop: sample.NewLaptop()},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := service.CreateLaptop(context.Background(), tc.req)
			if err != nil {
				panic(err)
			}

			fmt.Println(result)

		})
	}
}
