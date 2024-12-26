package controllers

import (
	"CatAPI/utils"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

func setupTestMainController(fetchDataFunc FetchDataFunc) (*MainController, *httptest.ResponseRecorder) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Create Beego context
	context := beecontext.NewContext()
	context.Reset(w, r)

	// Create and initialize controller
	c := &MainController{FetchData: fetchDataFunc}
	c.Init(context, "", "", nil)

	return c, w
}

// Mock implementation of FetchData
func mockFetchData(url string, key string, ch chan<- utils.APIResponse, extraParams map[string]string) {
	switch key {
	case "voting":
		ch <- utils.APIResponse{Key: "voting", Data: map[string]interface{}{"id": "123", "image_url": "https://example.com/image.jpg"}}
	case "breeds":
		ch <- utils.APIResponse{Key: "breeds", Data: []map[string]interface{}{{"id": "beng", "name": "Bengal"}}}
	case "favorites":
		ch <- utils.APIResponse{Key: "favorites", Data: []map[string]interface{}{{"id": "fav1", "image_url": "https://example.com/favorite.jpg"}}}
	default:
		ch <- utils.APIResponse{Key: key, Error: errors.New("mock error")}
	}
}

func TestIndex(t *testing.T) {
	controller, w := setupTestMainController(mockFetchData)

	// Execute the method
	controller.Index()

	// Assert status code
	assert.Equal(t, 200, w.Code, "Status code mismatch")

	// Validate the data passed to the template
	voting := controller.Data["Voting"]
	breeds := controller.Data["Breeds"]
	favorites := controller.Data["Favorites"]

	// Assert voting data
	expectedVoting := map[string]interface{}{"id": "123", "image_url": "https://example.com/image.jpg"}
	assert.Equal(t, expectedVoting, voting, "Voting data mismatch")

	// Assert breeds data
	expectedBreeds := []map[string]interface{}{{"id": "beng", "name": "Bengal"}}
	assert.Equal(t, expectedBreeds, breeds, "Breeds data mismatch")

	// Assert favorites data
	expectedFavorites := []map[string]interface{}{{"id": "fav1", "image_url": "https://example.com/favorite.jpg"}}
	assert.Equal(t, expectedFavorites, favorites, "Favorites data mismatch")
}

func TestIndexWithErrors(t *testing.T) {
	mockFetchDataWithError := func(url string, key string, ch chan<- utils.APIResponse, extraParams map[string]string) {
		ch <- utils.APIResponse{Key: key, Error: errors.New("mock error")}
	}

	controller, w := setupTestMainController(mockFetchDataWithError)

	// Execute the method
	controller.Index()

	// Assert status code
	assert.Equal(t, 200, w.Code, "Status code mismatch")

	// Validate the data passed to the template
	voting := controller.Data["Voting"]
	breeds := controller.Data["Breeds"]
	favorites := controller.Data["Favorites"]

	// Assert voting error
	assert.Equal(t, map[string]string{"error": "mock error"}, voting, "Voting error mismatch")

	// Assert breeds error
	assert.Equal(t, map[string]string{"error": "mock error"}, breeds, "Breeds error mismatch")

	// Assert favorites error
	assert.Equal(t, map[string]string{"error": "mock error"}, favorites, "Favorites error mismatch")
}
