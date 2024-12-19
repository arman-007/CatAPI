package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type VotingController struct {
	beego.Controller
}

// Fetch a random cat image for voting
func (c *VotingController) GetCat() {
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/images/search", nil)
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

// Submit a vote
func (c *VotingController) SubmitVote() {
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		fmt.Println("ERROR: Failed to read request body:", err)
		c.Data["json"] = map[string]string{"error": "Failed to read request body"}
		c.ServeJSON()
		return
	}
	fmt.Println("RAW REQUEST BODY:", string(body))

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
        c.Ctx.Output.SetStatus(400)
        fmt.Println("ERROR: Failed to parse JSON:", err)
        c.Data["json"] = map[string]string{"error": "Invalid payload"}
        c.ServeJSON()
        return
    }
	c.Data["json"] = map[string]string{"message": "Payload received successfully"}
    c.ServeJSON()

	reqBody, err := json.Marshal(payload)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to marshal payload"}
		c.ServeJSON()
		return
	}

	req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/votes", bytes.NewReader(reqBody))
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Set("x-api-key", "DEMO-API-KEY")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to send vote"}
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
