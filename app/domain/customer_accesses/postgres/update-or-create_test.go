package postgres

import (
	"reflect"
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
)

func TestCustomerAccessUpdateOrCreate(t *testing.T) {
	type args struct {
		data      *customer_accesses.CustomerAccess
		condition *customer_accesses.CustomerAccess
	}
	tests := []struct {
		name string
		p    *PostgreSQL
		args args
		want any
	}{
		{
			name: "UpdateOrCreate",
			p:    &PostgreSQL{DB: conn.DB},
			args: args{
				data: &customer_accesses.CustomerAccess{
					Token:      "test-token2",
					CustomerID: 1,
					Platform:   "test-platform",
					UserAgent:  "test-user-agent",
				},
				condition: &customer_accesses.CustomerAccess{
					CustomerID: 1,
					Platform:   "test-platform",
					UserAgent:  "test-user-agent"},
			},
			want: nil,
		},
	}

	if validatorhelper.IsNotNil(conn.Err) {
		t.Errorf("PostgreSQL.UpdateOrCreate() failed. error = %v", conn.Err.Error())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.UpdateOrCreate(tt.args.data, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostgreSQL.UpdateOrCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
