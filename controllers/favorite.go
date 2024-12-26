package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"bytes"
    "fmt"

	beego "github.com/beego/beego/v2/server/web"
)


// type HTTPClient interface {
//     Do(req *http.Request) (*http.Response, error)
// }

type FavoritesController struct {
	beego.Controller
    Client HTTPClient
}

func (c *FavoritesController) AddFavorite() {
    // Read the raw request body
    body, err := io.ReadAll(c.Ctx.Request.Body)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        fmt.Println("ERROR: Failed to read request body:", err)
        c.Data["json"] = map[string]string{"error": "Failed to read request body"}
        c.ServeJSON()
        return
    }

    // Parse the JSON from the request
    var payload map[string]interface{}
    if err := json.Unmarshal(body, &payload); err != nil {
        c.Ctx.Output.SetStatus(400)
        fmt.Println("ERROR: Failed to parse JSON:", err)
        c.Data["json"] = map[string]string{"error": "Invalid payload"}
        c.ServeJSON()
        return
    }

    // Log the parsed payload for debugging
    fmt.Println("PARSED PAYLOAD:", payload)

    // Check for required fields
    imageID, ok := payload["image_id"].(string)
    if !ok || imageID == "" {
        c.Ctx.Output.SetStatus(400)
        c.Data["json"] = map[string]string{"error": "image_id is required"}
        c.ServeJSON()
        return
    }

    subID, _ := payload["sub_id"].(string) // sub_id is optional

    // Prepare the payload for The Cat API
    catAPIPayload := map[string]string{
        "image_id": imageID,
    }
    if subID != "" {
        catAPIPayload["sub_id"] = subID
    }

    reqBody, err := json.Marshal(catAPIPayload)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to marshal payload"}
        c.ServeJSON()
        return
    }

    // Make the request to The Cat API
    req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/favourites", bytes.NewReader(reqBody))
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to create request"}
        c.ServeJSON()
        return
    }
    apiKey := beego.AppConfig.DefaultString("X-API-KEY", "DEMO-API-KEY")
	req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := c.Client
    if client == nil {
        client = &http.Client{}
    }

    // client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to add favorite"}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    // Forward the response from The Cat API back to the client
    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to read response"}
        c.ServeJSON()
        return
    }

    c.Ctx.Output.SetStatus(resp.StatusCode)
    c.Ctx.Output.Body(respBody)
}