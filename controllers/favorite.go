package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"bytes"
    "fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type FavoritesController struct {
	beego.Controller
}

func (c *FavoritesController) AddFavorite() {
    // Log the raw request body
    body, err := io.ReadAll(c.Ctx.Request.Body)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        fmt.Println("ERROR: Failed to read request body:", err)
        c.Data["json"] = map[string]string{"error": "Failed to read request body"}
        c.ServeJSON()
        return
    }
    // fmt.Println("RAW REQUEST BODY:", string(body))

    // Parse the JSON
    var payload map[string]interface{}
    // fmt.Println(payload)
    if err := json.Unmarshal(body, &payload); err != nil {
        c.Ctx.Output.SetStatus(400)
        fmt.Println("ERROR: Failed to parse JSON:", err)
        c.Data["json"] = map[string]string{"error": "Invalid payload"}
        c.ServeJSON()
        return
    }

    // Log the parsed payload
    // fmt.Println("PARSED PAYLOAD:", payload)

    // Proceed with the rest of your logic (e.g., forwarding to The Cat API)
    c.Data["json"] = map[string]string{"message": "Payload received successfully"}
    c.ServeJSON()

    reqBody, err := json.Marshal(payload)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to marshal payload"}
        c.ServeJSON()
        return
    }

    req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/favourites", bytes.NewReader(reqBody))
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to create request"}
        c.ServeJSON()
        return
    }
    req.Header.Set("x-api-key", "live_GQGS0iyuOQPXMeMpC7aTQle8rd1Go6WB3rmtDNBNxSg3xeK1INujU9tRhtZdH8v3")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to add favorite"}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    // body, err := io.ReadAll(resp.Body)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to read response"}
        c.ServeJSON()
        return
    }

    c.Ctx.Output.SetStatus(resp.StatusCode)
    c.Ctx.Output.Body(body)
}

// Fetch user's favorites
func (c *FavoritesController) GetFavorites() {
	subID := c.GetString("sub_id")
    // fmt.Println(subID)
	if subID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "sub_id is required"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/favourites?limit=20&page=0&order=Desc&sub_id=" + subID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Set("x-api-key", "live_GQGS0iyuOQPXMeMpC7aTQle8rd1Go6WB3rmtDNBNxSg3xeK1INujU9tRhtZdH8v3")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to fetch data"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to read response"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body(body)
}
