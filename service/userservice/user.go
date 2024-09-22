package userservice

import "github.com/yazdanbhd/Music-Cloud/entity"

type Repository interface {
	Register(u entity.User) (entity.User, error)
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
