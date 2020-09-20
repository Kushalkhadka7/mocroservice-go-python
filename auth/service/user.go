package service

import (
	"auth/model"
	pb "auth/pb"
	"auth/store"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

// UserService interface to implement.
type UserService interface {
	SaveUser(ctx context.Context, req *pb.User) (string, error)
	VerifyUser(ctx context.Context, req *pb.VerifyUserTokenRequest) (*UserInfo, error)
}

// CreateUserService initializes user service.
type CreateUserService struct {
	UserStore  store.UserStore
	JWTManager *JWTManager
}

// NewUserService creates new user service.
func NewUserService(store store.UserStore, jwtManger *JWTManager) UserService {
	return &CreateUserService{
		UserStore:  store,
		JWTManager: jwtManger,
	}
}

// SaveUser create and save new user.
func (service *CreateUserService) SaveUser(ctx context.Context, req *pb.User) (string, error) {
	hashedPasswored, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("i am called2")
		fmt.Println(err)
		return "", status.Errorf(codes.Internal, "Unable to bcrypt password %v", err)
	}

	user := model.NewUser(req.GetName(), string(hashedPasswored), req.GetRole())
	fmt.Println(user)

	token, err := service.JWTManager.GenerateToken(user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}

func (service *CreateUserService) VerifyUser(ctx context.Context, req *pb.VerifyUserTokenRequest) (*UserInfo, error) {
	accessToken := req.AccessToken

	userInfo, err := service.JWTManager.VerifyToken(accessToken)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
