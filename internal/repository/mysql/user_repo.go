package mysql

import (
	"github.com/dedegunawan/golang-clean-architecture/internal/domain/user"
	g "gorm.io/gorm"
)

type userRepository struct{ db *g.DB }

func NewUserRepository(db *g.DB) user.Repository {
	_ = db.AutoMigrate(&user.User{})
	return &userRepository{db}
}

func (r *userRepository) Create(u *user.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) FindByID(id uint64) (*user.User, error) {
	var u user.User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	var u user.User
	if err := r.db.Where("Email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) List(offset, limit int) ([]user.User, int64, error) {
	var items []user.User
	var total int64
	if err := r.db.Model(&user.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error
	return items, total, err
}

func (r *userRepository) UpdateAvatar(id uint64, url string) error {
	return r.db.Model(&user.User{}).Where("id = ?", id).Update("avatar_url", url).Error
}

func (r *userRepository) SetActive(id uint64, active bool) error {
	return r.db.Model(&user.User{}).Where("id = ?", id).Update("is_active", active).Error
}
