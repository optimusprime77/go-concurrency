package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"
)

type httpbinResponse struct {
	Args map[string]string `json:"args"`
	Data string            `json:"data"`
	Form map[string]string `json:"form"`
}

func TestRequestBytes(t *testing.T) {
	timeout := 2 * time.Second
	c := NewHTTPClient(Parameters{Timeout: &timeout})

	getPayload := &url.Values{}
	getPayload.Add("test", "ok")

	tcs := []struct {
		n           string
		d           HTTPRequestData
		expResponse httpbinResponse
		expError    error
	}{
		{
			"GET without GetPayload should be ok",
			HTTPRequestData{
				Method: http.MethodGet,
				URL:    "https://httpbin.org/get",
			},
			httpbinResponse{},
			nil,
		},
		{
			"GET with GetPayload should be ok",
			HTTPRequestData{
				Method:     http.MethodGet,
				URL:        "https://httpbin.org/get",
				GetPayload: getPayload,
			},
			httpbinResponse{
				Args: map[string]string{"test": "ok"},
			},
			nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.n, func(t *testing.T) {
			ctx := context.Background()
			res, err := c.RequestBytes(ctx, tc.d)

			if err != nil {
				if err.Error() != tc.expError.Error() {
					t.Errorf("unexpected response code: got %d, exp %d", err, tc.expError)
				}
			} else {
				if err != tc.expError {
					t.Errorf("unexpected response code: got %d, exp %d", err, tc.expError)
				}
			}

			if err == nil {
				var jsonResponse httpbinResponse

				json.Unmarshal(res, &jsonResponse)

				for k, v := range jsonResponse.Args {
					if v != tc.expResponse.Args[k] {
						t.Errorf("unexpected args returned: got %s, exp %s", v, tc.expResponse.Args[k])
					}
				}

				if jsonResponse.Data != tc.expResponse.Data {
					t.Errorf("unexpected data returned: got %s, exp %s", jsonResponse.Data, tc.expResponse.Data)
				}

				for k, v := range jsonResponse.Form {
					if v != tc.expResponse.Form[k] {
						t.Errorf("unexpected args returned: got %s, exp %s", v, tc.expResponse.Form[k])
					}
				}
			}
		})
	}
}
