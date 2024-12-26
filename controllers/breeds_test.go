package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// MockHTTPClient simulates an HTTP client for testing
// type MockHTTPClient struct {
// 	DoFunc func(req *http.Request) (*http.Response, error)
// }

// func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
// 	if m.DoFunc != nil {
// 		return m.DoFunc(req)
// 	}
// 	return nil, errors.New("DoFunc not implemented")
// }

func setupTestBreedsController(queryParams string) (*BreedsController, *httptest.ResponseRecorder) {
	r := httptest.NewRequest("GET", "/api/breeds/images?"+queryParams, nil)
	w := httptest.NewRecorder()

	// Create Beego context
	context := beecontext.NewContext()
	context.Reset(w, r)

	// Create and initialize controller
	c := &BreedsController{}
	c.Init(context, "", "", nil)

	return c, w
}

func TestGetBreedImages(t *testing.T) {
	tests := []struct {
		name           string
		queryParams    string
		mockDoFunc     func(req *http.Request) (*http.Response, error)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:        "Successful fetch of breed images",
			queryParams: "breed_id=beng",
			mockDoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`[{"url": "https://cdn2.thecatapi.com/images/beng.jpg"}]`)),
				}, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"url": "https://cdn2.thecatapi.com/images/beng.jpg"}]`,
		},
		{
			name:           "Missing breed_id parameter",
			queryParams:    "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"error": "breed_id is required"},
		},
		{
			name:        "API request failure",
			queryParams: "breed_id=beng",
			mockDoFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("API request failed")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   map[string]string{"error": "Failed to fetch data"},
		},
		{
			name:        "API response read failure",
			queryParams: "breed_id=beng",
			mockDoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(&FailingReader{}),
				}, nil
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   map[string]string{"error": "Failed to read response"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller and recorder
			controller, w := setupTestBreedsController(tt.queryParams)

			// Setup mock client if mockDoFunc is provided
			if tt.mockDoFunc != nil {
				controller.Client = &MockHTTPClient{DoFunc: tt.mockDoFunc}
			}

			// Execute the test
			controller.GetBreedImages()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code, "Status code mismatch")

			// Parse and verify the response body
			if tt.expectedStatus == http.StatusOK {
				assert.JSONEq(t, tt.expectedBody.(string), w.Body.String(), "Response body mismatch")
			} else {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "Failed to parse response body")
				assert.Equal(t, tt.expectedBody, response, "Response body mismatch")
			}
		})
	}
}

// FailingReader simulates a reader that fails
type FailingReader struct{}

func (r *FailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read failed")
}
