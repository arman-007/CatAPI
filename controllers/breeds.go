package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	web.Controller
}

// Breeds Page
func (c *BreedsController) GetBreeds() {
	// Fetch the list of breeds from The Cat API
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/breeds", nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Set("x-api-key", "DEMO-API-KEY") // Use your API key

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to fetch breeds"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to read response body"}
		c.ServeJSON()
		return
	}

	var breeds []map[string]interface{}
	if err := json.Unmarshal(body, &breeds); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to parse JSON"}
		c.ServeJSON()
		return
	}

	// Pass the breeds to the template
	c.Data["Breeds"] = breeds
	c.TplName = "breeds.tpl"
}
