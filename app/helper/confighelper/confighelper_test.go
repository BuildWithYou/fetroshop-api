package confighelper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name string
		want *viper.Viper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintConfig(t *testing.T) {
	tests := []struct {
		name string
		key  string
	}{
		{
			name: "Print environment",
			key:  "environment",
		},
		{
			name: "Print database.logLevel",
			key:  "database.logLevel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := GetConfig()
			fmt.Println(config.GetString(tt.key))
		})
	}
}
