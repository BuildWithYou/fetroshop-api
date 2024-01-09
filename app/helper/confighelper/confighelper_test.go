package confighelper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	t.Run("GetConfig", func(t *testing.T) {
		config := GetConfig()
		assert.NotNil(t, config)
	})
}

func TestPrintConfig(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		typeData string
	}{
		{
			name:     "Print environment",
			key:      "environment",
			typeData: "string",
		},
		{
			name:     "Print database.logLevel",
			key:      "database.logLevel",
			typeData: "string",
		},
		{
			name:     "Print security.jwt.tokenKey",
			key:      "security.jwt.tokenKey",
			typeData: "string",
		},
	}

	config := GetConfig()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.typeData {
			case "string":
				fmt.Println(config.GetString(tt.key))
			}
		})
	}
}
