package postgres

import (
	"reflect"
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"gorm.io/gorm"
)

func TestCustomerAccessCreate(t *testing.T) {
	type args struct {
		data *customer_accesses.CustomerAccess
	}
	tests := []struct {
		name string
		p    *PostgreSQL
		args args
		want *gorm.DB
	}{
		{
			name: "Create",
			p:    &PostgreSQL{DB: conn.DB},
			args: args{
				data: &customer_accesses.CustomerAccess{
					Token:      "test-token",
					CustomerID: 1,
					Platform:   "test-platform",
					UserAgent:  "test-user-agent",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Create(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostgreSQL.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
