package user

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/dedegunawan/golang-clean-architecture/pkg/logger"
)

type Service interface {
	Register(name, email, password string) (*User, error)
	Get(id uint64) (*User, error)
	List(page, size int) ([]User, int64, error)
	UpdateAvatar(id uint64, url string) error
	SetActive(id uint64, active bool) error
}

type service struct {
	repo Repository
	lg   *logger.Logger
}

func NewService(r Repository, lg *logger.Logger) Service {
	return &service{repo: r, lg: lg}
}

func (s *service) Register(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil { return nil, err }
	u := &User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
		IsActive:     true,
	}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Get(id uint64) (*User, error) { return s.repo.FindByID(id) }

func (s *service) List(page, size int) ([]User, int64, error) {
	if page < 1 { page = 1 }
	if size <= 0 || size > 100 { size = 10 }
	offset := (page - 1) * size
	return s.repo.List(offset, size)
}

func (s *service) UpdateAvatar(id uint64, url string) error { return s.repo.UpdateAvatar(id, url) }
func (s *service) SetActive(id uint64, active bool) error   { return s.repo.SetActive(id, active) }
