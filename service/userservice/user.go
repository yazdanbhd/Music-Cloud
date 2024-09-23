package userservice

import (
	"github.com/golang-jwt/jwt"
	"github.com/yazdanbhd/Music-Cloud/delivery/authjwt"
	"github.com/yazdanbhd/Music-Cloud/entity"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	IsAuthenticated(userName, password string) (bool, error)
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type Service struct {
	repo Repository
}

func New(r Repository) Service {
	return Service{repo: r}
}

func (s *Service) UserRegister(req RegisterRequest) (RegisterResponse, error) {
	// Store the user data to the database
	user := entity.User{
		ID:          0,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}
	u, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, err
	}
	return RegisterResponse{UserID: u.ID, Name: u.Name}, nil
}

func (s *Service) UserLogin(loginReq LoginRequest) (LoginResponse, error) {
	isAuth, err := s.repo.IsAuthenticated(loginReq.Name, loginReq.Password)

	if err != nil || isAuth == false {
		return LoginResponse{}, err
	}

	token := authjwt.New([]byte(`secret-key`), jwt.SigningMethodHS256)

	tokenString, err := token.CreateToken(loginReq.Name)

	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{AccessToken: tokenString}, nil
}
