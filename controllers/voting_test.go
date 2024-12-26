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
    beego "github.com/beego/beego/v2/server/web"
    beecontext "github.com/beego/beego/v2/server/web/context"
)

type MockHTTPClient struct {
    DoFunc func(req *http.Request) (*http.Response, error)
}

// func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
//     if m.DoFunc != nil {
//         return m.DoFunc(req)
//     }
//     return nil, errors.New("DoFunc not implemented")
// }

func init() {
    beego.BConfig.RunMode = "test"
    beego.BConfig.Log.AccessLogs = false
    beego.BConfig.Log.FileLineNum = false
    beego.BConfig.CopyRequestBody = true
}

func setupTestController(body string) (*VotingController, *httptest.ResponseRecorder) {
    r := httptest.NewRequest("POST", "/api/voting/vote", bytes.NewBufferString(body))
    w := httptest.NewRecorder()

    // Create beego context
    context := beecontext.NewContext()
    context.Reset(w, r)

    // Create and initialize controller
    c := &VotingController{}
    c.Init(context, "", "", nil)

    return c, w
}

func TestSubmitVote(t *testing.T) {
    tests := []struct {
        name           string
        requestBody    string
        mockDoFunc     func(req *http.Request) (*http.Response, error)
        expectedStatus int
        expectedBody   map[string]string
    }{
        {
            name:        "Successful vote submission",
            requestBody: `{"image_id": "4RzEwvyzz", "sub_id": "demo-7t38rnyvj", "value": true}`,
            mockDoFunc: func(req *http.Request) (*http.Response, error) {
                return &http.Response{
                    StatusCode: http.StatusOK,
                    Body:       io.NopCloser(bytes.NewBufferString(`{"message":"SUCCESS"}`)),
                }, nil
            },
            expectedStatus: http.StatusOK,
            expectedBody:   map[string]string{"message": "Vote submitted successfully"},
        },
        {
            name:           "Invalid JSON payload",
            requestBody:    `invalid json`,
            expectedStatus: http.StatusBadRequest,
            expectedBody:   map[string]string{"error": "Invalid payload"},
        },
        {
            name:        "API call failure",
            requestBody: `{"image_id": "abc123", "sub_id": "user123", "value": true}`,
            mockDoFunc: func(req *http.Request) (*http.Response, error) {
                return nil, errors.New("failed to send vote")
            },
            expectedStatus: http.StatusInternalServerError,
            expectedBody:   map[string]string{"error": "Failed to send vote"},
        },
        // {
        //     name:           "Failed to read request body",
        //     requestBody:    `invalid body`, // Simulating invalid body using FailingReader
        //     mockDoFunc:     nil,
        //     expectedStatus: http.StatusInternalServerError,
        //     expectedBody:   map[string]string{"error": "Invalid payload"},
        // },
        // {
        //     name:           "Failed to marshal payload",
        //     requestBody:    `{"key": "value"}`, // Valid JSON but can simulate marshal failure
        //     mockDoFunc:     nil,
        //     expectedStatus: http.StatusInternalServerError,
        //     expectedBody:   map[string]string{"error": "Failed to marshal payload"},
        // },
        // {
        //     name:           "Failed to create HTTP request",
        //     requestBody:    `{"image_id": "abc123", "sub_id": "user123", "value": true}`,
        //     mockDoFunc:     nil,
        //     expectedStatus: http.StatusInternalServerError,
        //     expectedBody:   map[string]string{"error": "Failed to create request"},
        // },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup controller and recorder
            controller, w := setupTestController(tt.requestBody)

            // Setup mock client if mockDoFunc is provided
            if tt.mockDoFunc != nil {
                controller.Client = &MockHTTPClient{DoFunc: tt.mockDoFunc}
            }

            // Execute the test
            controller.SubmitVote()

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