package category_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/BuildWithYou/fetroshop-api/app/injector"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/stretchr/testify/assert"
)

var fiberApp = injector.InitializeWebServer()

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
			name: "ListNotEmptyParentCode200",
			args: model.ListCategoriesRequest{
				ParentCode:     "kebutuhan-dapur",
				Offset:         0,
				Limit:          10,
				OrderBy:        "display_order",
				OrderDirection: "ASC",
			},
			wantResponseCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Define the form data
			/* formData := url.Values{
				"parentCode": {tt.args.ParentCode},
				"offset":     {strconv.Itoa(int(tt.args.Offset))},
				"limit":      {strconv.Itoa(int(tt.args.Limit))},
				"orderBy":    {tt.args.OrderBy},
			} */

			formData := url.Values{}
			formData.Set("parentCode", tt.args.ParentCode)
			formData.Set("offset", strconv.Itoa(int(tt.args.Offset)))
			formData.Set("limit", strconv.Itoa(int(tt.args.Limit)))
			formData.Set("orderBy", tt.args.OrderBy)
			formData.Set("orderDirection", tt.args.OrderDirection)

			fmt.Println("formData : ", formData)

			// Define the request body
			body := strings.NewReader(formData.Encode())

			request := httptest.NewRequest(http.MethodGet, "/api/category/list", body)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			request.Header.Set("Sec-Ch-Ua-Platform", "sec-ch-ua-platform-test")
			request.Header.Set("User-Agent", "user-agent-test")
			response, err := fiberApp.FiberApp.Test(request)
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
