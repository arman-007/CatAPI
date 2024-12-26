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
	// beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// type MockHTTPClient struct {
// 	DoFunc func(req *http.Request) (*http.Response, error)
// }

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return nil, errors.New("DoFunc not implemented")
}

func setupTestFavoritesController(body string) (*FavoritesController, *httptest.ResponseRecorder) {
	r := httptest.NewRequest("POST", "/api/favorites/add", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	// Create Beego context
	context := beecontext.NewContext()
	context.Reset(w, r)

	// Create and initialize controller
	c := &FavoritesController{}
	c.Init(context, "", "", nil)

	return c, w
}

func TestAddFavorite(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		mockDoFunc     func(req *http.Request) (*http.Response, error)
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:        "Successful favorite addition",
			requestBody: `{"image_id": "abc123", "sub_id": "user123"}`,
			mockDoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{"message":"SUCCESS"}`)),
				}, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"message": "SUCCESS"},
		},
		{
			name:           "Missing image_id",
			requestBody:    `{"sub_id": "user123"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"error": "image_id is required"},
		},
		{
			name:           "Invalid JSON payload",
			requestBody:    `invalid json`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"error": "Invalid payload"},
		},
		{
			name:        "API call failure",
			requestBody: `{"image_id": "abc123", "sub_id": "user123"}`,
			mockDoFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("API request failed")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   map[string]string{"error": "Failed to add favorite"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup controller and recorder
			controller, w := setupTestFavoritesController(tt.requestBody)

			// Setup mock client if mockDoFunc is provided
			if tt.mockDoFunc != nil {
				controller.Client = &MockHTTPClient{DoFunc: tt.mockDoFunc}
			}

			// Execute the test
			controller.AddFavorite()

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code, "Status code mismatch")

			// Parse response body
			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err, "Failed to parse response body")
			assert.Equal(t, tt.expectedBody, response, "Response body mismatch")
		})
	}
}
