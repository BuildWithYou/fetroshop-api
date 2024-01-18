package brand_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/injector"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/stretchr/testify/assert"
)

var fetroshopApp = injector.InitializeWebServer()
var webLogger = logger.NewWebLogger(confighelper.GetConfig())

func TestCategoryServiceList(t *testing.T) {
	tests := []struct {
		name             string
		args             model.ListCategoriesRequest
		wantResponseCode int
		wantErr          bool
	}{
		{
			name: "ListEmptyParentCode200",
			args: model.ListCategoriesRequest{
				ParentCode:     "",
				Offset:         0,
				Limit:          10,
				OrderBy:        "display_order",
				OrderDirection: "ASC",
			},
			wantResponseCode: 200,
		},
		{
			name: "ListCorrectParentCode200",
			args: model.ListCategoriesRequest{
				ParentCode:     "kebutuhan-dapur",
				Offset:         0,
				Limit:          10,
				OrderBy:        "display_order",
				OrderDirection: "ASC",
			},
			wantResponseCode: 200,
		},
		{
			name: "ListWrongParentCode400",
			args: model.ListCategoriesRequest{
				ParentCode:     "wrong-parent-code",
				Offset:         0,
				Limit:          10,
				OrderBy:        "display_order",
				OrderDirection: "ASC",
			},
			wantResponseCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := url.Values{}
			query.Add("parentCode", tt.args.ParentCode)
			query.Add("offset", strconv.Itoa(int(tt.args.Offset)))
			query.Add("limit", strconv.Itoa(int(tt.args.Limit)))
			query.Add("orderBy", tt.args.OrderBy)
			query.Add("orderDirection", tt.args.OrderDirection)

			request := httptest.NewRequest(http.MethodGet, "/api/category/list?"+query.Encode(), nil)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			request.Header.Set("Sec-Ch-Ua-Platform", "sec-ch-ua-platform-test")
			request.Header.Set("User-Agent", "user-agent-test")

			response, err := fetroshopApp.FiberApp.Test(request)
			assert.Nil(t, err)
			assert.Equal(t, tt.wantResponseCode, response.StatusCode)

			bytes, err := io.ReadAll(response.Body)
			assert.Nil(t, err)
			assert.NotNil(t, bytes)

			if response.StatusCode != tt.wantResponseCode {
				webLogger.LogConsole.Error(fmt.Sprintln("Response : ", string(bytes)))
			}

		})
	}
}
