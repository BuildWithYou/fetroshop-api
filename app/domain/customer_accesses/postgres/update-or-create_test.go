package postgres

import (
	"reflect"
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"gorm.io/gorm"
)

var dbConnection = connection.OpenDBConnection(confighelper.GetConfig())

func TestPostgreSQLUpdateOrCreate(t *testing.T) {
	type args struct {
		data      *customer_accesses.CustomerAccess
		condition *customer_accesses.CustomerAccess
	}
	tests := []struct {
		name string
		p    *PostgreSQL
		args args
		want *gorm.DB
	}{
		{
			name: "UpdateOrCreate",
			p:    &PostgreSQL{DB: dbConnection},
			args: args{
				data:      &customer_accesses.CustomerAccess{},
				condition: &customer_accesses.CustomerAccess{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.UpdateOrCreate(tt.args.data, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostgreSQL.UpdateOrCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
