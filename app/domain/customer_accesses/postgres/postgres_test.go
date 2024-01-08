package postgres

import (
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/stretchr/testify/assert"
)

var conn = connection.OpenDBConnection(connection.DB_TEST, confighelper.GetConfig())

func TestCustomerAccessRepoProvider(t *testing.T) {
	type args struct {
		db *connection.Connection
	}
	tests := []struct {
		name string
		args args
		want customer_accesses.CustomerAccessRepo
	}{
		{
			name: "CustomerAccessRepoProvider",
			args: args{
				db: conn,
			},
			want: &PostgreSQL{DB: conn.DB},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomerAccessRepoProvider(tt.args.db)
			assert.NotNil(t, got)
		})
	}
}
