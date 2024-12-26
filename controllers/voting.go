package controllers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    beego "github.com/beego/beego/v2/server/web"
)

type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

type VotingController struct {
    beego.Controller
    Client HTTPClient
}

func (c *VotingController) SubmitVote() {
    body, err := io.ReadAll(c.Ctx.Request.Body)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        fmt.Println("ERROR: Failed to read request body:", err)
        c.Data["json"] = map[string]string{"error": "Failed to read request body"}
        c.ServeJSON()
        return
    }

    var payload map[string]interface{}
    if err := json.Unmarshal(body, &payload); err != nil {
        c.Ctx.Output.SetStatus(400)
        fmt.Println("ERROR: Failed to parse JSON:", err)
        c.Data["json"] = map[string]string{"error": "Invalid payload"}
        c.ServeJSON()
        return
    }

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
    apiKey := beego.AppConfig.DefaultString("X-API-KEY", "DEMO-API-KEY")
	req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := c.Client
    if client == nil {
        client = &http.Client{}
    }

    resp, err := client.Do(req)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = map[string]string{"error": "Failed to send vote"}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    // Set the response status code
    c.Ctx.Output.SetStatus(resp.StatusCode)

    if resp.StatusCode >= 400 {
        c.Data["json"] = map[string]string{"error": "API request failed"}
    } else {
        c.Data["json"] = map[string]string{"message": "Vote submitted successfully"}
    }
    c.ServeJSON()
}