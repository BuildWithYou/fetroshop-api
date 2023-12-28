package postgres

type PostgreSQL struct{}

func New() *PostgreSQL {
	return &PostgreSQL{}
}
