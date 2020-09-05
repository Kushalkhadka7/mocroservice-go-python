package laptopservice

import (
	"context"
	"fmt"
	"manager/sample"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func TestSave(t *testing.T) {
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

			store := NewLaptopStore(db)
			result, err := store.Save(context.Background(), sample.NewLaptop())
			if err != nil {
				panic(err)
			}

			if result == sample.NewLaptop() {
				fmt.Println("Laptop creation successful")
			}

		})
	}

}
