package connection

import (
	"reflect"
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestOpenDBConnection(t *testing.T) {
	type args struct {
		dbType DBType
		config *viper.Viper
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
		{
			name: "OpenMainDBConnection",
			args: args{
				dbType: DB_MAIN,
				config: confighelper.GetConfig(),
			},
			want: nil,
		},
		{
			name: "OpenTestDBConnection",
			args: args{
				dbType: DB_TEST,
				config: confighelper.GetConfig(),
			},
			want: nil,
		},
		{
			name: "OpenWrongDBConnection",
			args: args{
				dbType: "wrong",
				config: confighelper.GetConfig(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenDBConnection(tt.args.dbType, tt.args.config)
			switch tt.name {
			case "OpenMainDBConnection", "OpenTestDBConnection":
				{
					_, err := OpenDBConnection(tt.args.dbType, tt.args.config)
					if validatorhelper.IsNotNil(err) {
						t.Errorf("%s failed. Error: %v", tt.name, err.Error())
					}
					assert.Nil(t, err)
				}
			case "OpenWrongDBConnection":
				{
					_, err := OpenDBConnection(tt.args.dbType, tt.args.config)
					assert.NotNil(t, err)
				}
			default:
				{
					if err != nil {
						t.Errorf("%s failed. Error: %v", tt.name, err.Error())
					}
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
					}

				}
			}
		})
	}
}
