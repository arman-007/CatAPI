package utils

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
)

type APIResponse struct {
	Key   string
	Data  interface{}
	Error error
}

func FetchData(url, key string, ch chan<- APIResponse, queryParams map[string]string) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", "live_GQGS0iyuOQPXMeMpC7aTQle8rd1Go6WB3rmtDNBNxSg3xeK1INujU9tRhtZdH8v3")

	// Add query parameters
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- APIResponse{Key: key, Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- APIResponse{Key: key, Error: err}
		return
	}

	// Log the raw response body for debugging
	// fmt.Printf("Response from %s: %s\n", url, string(body))

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- APIResponse{Key: key, Error: err}
		return
	}

	ch <- APIResponse{Key: key, Data: data}
}