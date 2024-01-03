package postgres

import "github.com/BuildWithYou/fetroshop-api/app/domain/users"

type PostgreSQL struct{}

func NewUserRepository() users.UserRepository {
	return &PostgreSQL{}
}
