package pigeon

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/optimusprime77/go-concurrency/client"
	log "github.com/sirupsen/logrus"
)

// FetchData returns response from endpoint specified.
func (c *Client) FetchData(ctx context.Context, url string) ([]byte, error) {

	// Implement semaphores to ensure maximum concurrency threshold.
	c.semaphore <- struct{}{}
	defer func() { <-c.semaphore }()

	// If there is an in-flight request for a unique URL, send response
	// from the in-flight request. Else, create the in-flight request.
	responseRaw, err, shared := c.RequestGroup.Do(url, func() (interface{}, error) {
		return c.fetchResponse(ctx)
	})

	if err != nil {
		return []byte{}, err
	}

	log.Infof("in-flight status : %t", shared)

	//time.Sleep(time.Second * 4)

	response := responseRaw.([]byte)

	return response, err
}

func (c *Client) fetchResponse(ctx context.Context) ([]byte, error) {
	payload := struct {
		Number int `json:"number"`
	}{
		Number: 4,
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return []byte{}, err
	}

	reqData := client.HTTPRequestData{
		Method:      http.MethodPost,
		URL:         c.Endpoint + "/post",
		Headers:     map[string]string{"Content-Type": "application/json"},
		PostPayload: jsonPayload,
	}

	return c.HTTPClient.RequestBytes(ctx, reqData)
}
