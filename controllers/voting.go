package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "bytes"

	"github.com/beego/beego/v2/server/web"
)

type VotingController struct {
	web.Controller
}

func (c *VotingController) Get() {
	// Prepare the request
	fmt.Println("voting api called")
	url := "https://api.thecatapi.com/v1/images/search?limit=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = "Failed to create request: " + err.Error()
		c.TplName = "error.tpl"
		return
	}

	// Add headers
	req.Header.Set("x-api-key", "live_GQGS0iyuOQPXMeMpC7aTQle8rd1Go6WB3rmtDNBNxSg3xeK1INujU9tRhtZdH8v3")

	// Make the HTTP call
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = "Failed to fetch data: " + err.Error()
		c.TplName = "error.tpl"
		return
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = "Failed to read response body: " + err.Error()
		c.TplName = "error.tpl"
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = "Failed to parse JSON: " + err.Error()
		c.TplName = "error.tpl"
		return
	}

	// Pass the data to the template
	// fmt.Println(result)
	c.Data["Votes"] = result
	c.TplName = "voting.tpl"
}

// func (c *VotingController) Vote() {
// 	var payload struct {
// 		ImageID string `json:"image_id"`
// 		SubID   string `json:"sub_id"`
// 		Value   bool   `json:"value"` // Use bool for true/false
// 	}

// 	// Parse the incoming JSON payload
// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = map[string]string{"error": "Invalid payload"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Validate payload
// 	if payload.ImageID == "" {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = map[string]string{"error": "Invalid payload values"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Forward the payload to The Cat API
// 	reqBody, err := json.Marshal(payload)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(500)
// 		c.Data["json"] = map[string]string{"error": "Failed to marshal payload"}
// 		c.ServeJSON()
// 		return
// 	}

// 	req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/votes", bytes.NewReader(reqBody))
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(500)
// 		c.Data["json"] = map[string]string{"error": "Failed to create request"}
// 		c.ServeJSON()
// 		return
// 	}
// 	// req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("x-api-key", "live_GQGS0iyuOQPXMeMpC7aTQle8rd1Go6WB3rmtDNBNxSg3xeK1INujU9tRhtZdH8v3")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(500)
// 		c.Data["json"] = map[string]string{"error": "Failed to send vote"}
// 		c.ServeJSON()
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(500)
// 		c.Data["json"] = map[string]string{"error": "Failed to read response"}
// 		c.ServeJSON()
// 		return
// 	}

// 	c.Ctx.Output.SetStatus(resp.StatusCode)
// 	c.Ctx.Output.Body(body)
// }
