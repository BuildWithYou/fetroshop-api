package connection

import (
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestOpenDBConnection(t *testing.T) {
	type args struct {
		dbType DBType
		config *viper.Viper
	}
	config := confighelper.GetConfig()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "OpenMainDBConnection",
			args: args{
				dbType: DB_MAIN,
				config: config,
			},
		},
		{
			name: "OpenTestDBConnection",
			args: args{
				dbType: DB_TEST,
				config: config,
			},
		},
		{
			name: "OpenWrongDBConnection",
			args: args{
				dbType: "wrong",
				config: config,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := OpenDBConnection(tt.args.dbType, tt.args.config)
			got := conn.DB
			err := conn.Err
			switch tt.name {
			case "OpenMainDBConnection", "OpenTestDBConnection":
				{
					if validatorhelper.IsNotNil(err) {
						t.Errorf("%s failed. Error: %v", tt.name, err.Error())
					}
					assert.NotNil(t, got)
					assert.Nil(t, err)
				}
			case "OpenWrongDBConnection":
				{
					assert.Nil(t, got)
					assert.NotNil(t, err)
				}
			}
		})
	}
}
