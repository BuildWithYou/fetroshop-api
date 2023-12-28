package users

type UserRepository interface {
	Create(user *User) error
}
