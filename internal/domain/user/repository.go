package user

type Repository interface {
	Create(u *User) error
	FindByID(id uint64) (*User, error)
	List(offset, limit int) ([]User, int64, error)

	UpdateAvatar(id uint64, url string) error
	SetActive(id uint64, active bool) error
}
