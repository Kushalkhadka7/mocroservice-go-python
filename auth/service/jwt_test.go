package service_test

import (
	"auth/model"
	"auth/service"
	"fmt"
	"testing"
	"time"
)

var mockJwtManger = service.NewJWTMnanager("kushal", 5*time.Second)
var user = model.NewUser("kushal", "khadka", "admin")

func TestJwt(t *testing.T) {
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

			res, err := mockJwtManger.GenerateToken(user)
			if err != nil {
				panic(err)
			}

			fmt.Println(res)
		})

	}
}
