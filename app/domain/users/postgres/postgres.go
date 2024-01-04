package postgres

import "github.com/BuildWithYou/fetroshop-api/app/domain/users"

type PostgreSQL struct{}

func UserRepositoryProvider() users.UserRepository {
	return &PostgreSQL{}
}
