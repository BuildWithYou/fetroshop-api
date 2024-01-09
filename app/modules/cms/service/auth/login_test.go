package auth_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/injector"
	"github.com/stretchr/testify/assert"
)

var cmsApp = injector.InitializeCmsServer()

func TestCmsServiceLogin(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name             string
		args             args
		wantResponseCode int
		wantErr          bool
	}{
		{
			name: "Login200",
			args: args{
				username: "string",
				password: "string",
			},
			wantResponseCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(fmt.Sprintf(`{
            "password": "%s",
            "username": "%s"
         }`, tt.args.password, tt.args.username))
			request := httptest.NewRequest("POST", "/api/auth/login", body)
			request.Header.Set("Content-Type", "application/json")
			response, err := cmsApp.FiberApp.Test(request)
			assert.Nil(t, err)
			assert.Equal(t, tt.wantResponseCode, response.StatusCode)

			bytes, err := io.ReadAll(response.Body)
			assert.Nil(t, err)
			assert.NotNil(t, bytes)

			if response.StatusCode != tt.wantResponseCode {
				fmt.Println("Response : ", string(bytes)) // #marked: logging
			}

		})
	}
}
