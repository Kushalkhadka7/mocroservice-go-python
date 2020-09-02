package laptopservice

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		error string
	}{
		{
			name:  "Kushal",
			error: "hello wold",
		},
		{
			name:  "Nischal",
			error: "NOt hello worl",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var db LaptopInterface

			laptopstore := NewLaptopStore(db)

			err := laptopstore.Save()

			fmt.Println(err)

		})
	}

}
