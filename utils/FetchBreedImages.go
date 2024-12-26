package utils

import(
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchBreedImages(breedID string) []map[string]string {
    url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_id=%s&limit=5", breedID)
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("x-api-key", "your-api-key")
	req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Error fetching breed images:", err)
        return nil
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    var images []map[string]string
    json.Unmarshal(body, &images)

    return images
}
