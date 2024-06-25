package user

type (
	UserRepository interface {
		FindByPhone(phone string) (User, error)
		Save(user User) error
		Delete(user User) error
		List() ([]User, error)
	}
	User struct {
		phone    string
		password string
	}
)

func NewUser(phone string, password string) User {
	return User{phone: phone, password: password}
}

func (u User) Phone() string {
	return u.phone
}

func (u User) Password() string {
	return u.password
}
