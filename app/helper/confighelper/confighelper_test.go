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
