package laptopservice

import (
	"context"
	"fmt"
	laptop "manager/pb"
	"testing"
)

func TestSaveService(t *testing.T) {
	var db LaptopInterface
	LaptopStore := NewLaptopStore(db)
	laptopService := NewLaptopService(LaptopStore)

	res, err := laptopService.CreateLaptop(context.Background(), &laptop.CreateLaptopRequest{})

	if err != nil {
		fmt.Printf("error %s", err)
	}

	fmt.Printf("Success%s", res)
}
