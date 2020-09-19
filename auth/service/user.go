package service

import (
	"auth/model"
	auth "auth/pbauth"
	"auth/store"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

// UserService interface to implement.
type UserService interface {
	SaveUser(ctx context.Context, req *auth.User) (*model.User, error)
}

// CreateUserService initializes user service.
type CreateUserService struct {
	UserStore store.UserStore
}

// NewUserService creates new user service.
func NewUserService(store store.UserStore) UserService {
	return &CreateUserService{
		UserStore: store,
	}
}

// SaveUser create and save new user.
func (service *CreateUserService) SaveUser(ctx context.Context, req *auth.User) (*model.User, error) {
	hashedPasswored, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("i am called2")
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Unable to bcrypt password %v", err)
	}

	user := model.NewUser(req.GetName(), string(hashedPasswored), req.GetRole())
	fmt.Println(user)

	res, err := service.UserStore.Save(user)
	if err != nil {
		fmt.Println("i am called3")
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
