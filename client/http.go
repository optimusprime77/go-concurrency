package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Parameters provides the parameters used when creating a new http client
type Parameters struct {
	Timeout *time.Duration
}

// NewHTTPClient instantiates a new HTTPClient
func NewHTTPClient(parameters Parameters) HTTPClient {
	if parameters.Timeout == nil {
		timeout := 1 * time.Second
		parameters.Timeout = &timeout
	}

	client := &http.Client{
		Timeout: *parameters.Timeout,
	}

	return HTTPClient{client}
}

// HTTPRequestData contains the request data
type HTTPRequestData struct {
	Method      string
	URL         string
	Headers     map[string]string
	PostPayload []byte
	GetPayload  *url.Values
}

// HTTPClient contains the http client
type HTTPClient struct {
	*http.Client
}

// RequestBytes returns whatever the request returns in a slice of byte.
func (client *HTTPClient) RequestBytes(ctx context.Context, reqData HTTPRequestData) ([]byte, error) {
	r, err := client.request(ctx, reqData)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return ioutil.ReadAll(r.Body)
}

func (client *HTTPClient) request(ctx context.Context, reqData HTTPRequestData) (*http.Response, error) {
	var req *http.Request
	var err error

	if reqData.Method == http.MethodPost {
		req, err = http.NewRequest(reqData.Method, reqData.URL, bytes.NewBuffer(reqData.PostPayload))
	} else {
		req, err = http.NewRequest(reqData.Method, reqData.URL, nil)
	}

	if err != nil {
		return nil, err
	}

	if reqData.GetPayload != nil {
		req.URL.RawQuery = reqData.GetPayload.Encode()
	}

	for k, v := range reqData.Headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("User-Agent", "go-concurrency")

	resp, err := client.Do(req)

	if err != nil {
		if reqData.Method == http.MethodPost {
			return resp, fmt.Errorf("Error making request: %v. Body: %s", err, reqData.PostPayload)
		}

		return resp, fmt.Errorf("Error making request: %v. Query: %v", err, req.URL.RawQuery)
	}

	if resp.StatusCode >= 400 {
		return resp, fmt.Errorf("Error response from %s, got status: %d", reqData.URL, resp.StatusCode)
	}

	return resp, nil
}
