package user

import (
	"context"
	"strconv"
	"time"

	"github.com/TimurZheksimbaev/Golang-webchat/config"
	"github.com/TimurZheksimbaev/Golang-webchat/utils"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	repository Repository
	timeout time.Duration
	appConfig *config.AppConfig
}

func NewService(repository Repository, appConfig *config.AppConfig) Service {
	return &service{
		repository: repository,
		timeout: time.Duration(2) * time.Second,
		appConfig: appConfig,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, utils.ServiceError("Could not encrypt password", err)
	}

	u := &User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	r, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		return nil, utils.ServiceError("Could not create user", err)
	}
	response := &CreateUserResponse{
		ID: strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email: r.Email,
	}
	return response, nil 
}

type MyJWTClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserRequest) (*LoginUserResponse, error ) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserResponse{}, utils.ServiceError("Could not get user by email", err)
	}

	err = utils.ComparePasswords(u.Password, req.Password)
	if err != nil {
		return &LoginUserResponse{}, utils.ServiceError("Passwords do not match", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.appConfig.JWTExpiration)),
		},
	})

	ss, err := token.SignedString([]byte(s.appConfig.SecretKey))
	if err != nil {
		return &LoginUserResponse{}, utils.ServiceError("Could not sign token", err)
	}

	return &LoginUserResponse{
		accessToken: ss,
		Username: u.Username,
		ID: strconv.Itoa(int(u.ID)),
	}, nil



}