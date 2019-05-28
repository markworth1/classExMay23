package service

import (
	"context"
	"encoding/json"

	"github.com/gSchool/classExMay23/users/pkg/repo"
)

// User struct
type User struct {
	UserEmail       string `json:"email"`
	UserName        string `json:"name"`
	UserPhonenumber string `json:"phone"`
}

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	Create(ctx context.Context, user User) error
	GetAll(ctx context.Context) ([]string, error)
}

type basicUsersService struct {
	repo.Repo
}

func (b *basicUsersService) Create(ctx context.Context, user User) (e0 error) {
	tmp, _ := json.Marshal(user)
	b.Repo.Insert(string(tmp))
	return e0
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	return &basicUsersService{}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUsersService) GetAll(ctx context.Context) (s0 []string, e1 error) {
	s0 = b.Repo.GetAllData()
	return s0, e1
}
