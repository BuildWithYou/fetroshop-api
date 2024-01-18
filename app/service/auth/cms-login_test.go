package auth_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmsServiceLogin(t *testing.T) {
	type args struct {
		username  string
		password  string
		platform  string
		userAgent string
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
				username:  "testercmsapi",
				password:  "testertester",
				platform:  "tester-platform",
				userAgent: "tester-userAgent",
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
			request := httptest.NewRequest(http.MethodPost, "/api/auth/login", body)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Sec-Ch-Ua-Platform", "sec-ch-ua-platform-test")
			request.Header.Set("User-Agent", "user-agent-test")
			response, err := cmsServer.FiberApp.Test(request)
			assert.Nil(t, err)
			assert.Equal(t, tt.wantResponseCode, response.StatusCode)

			bytes, err := io.ReadAll(response.Body)
			assert.Nil(t, err)
			assert.NotNil(t, bytes)

			if response.StatusCode != tt.wantResponseCode {
				fmt.Println("cmsLogger : ", cmsLogger)
				cmsLogger.LogConsole.Error(fmt.Sprintln("Response : ", string(bytes)))
			}

		})
	}
}
