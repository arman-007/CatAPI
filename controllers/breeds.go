package controllers

import (
	// "encoding/json"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	beego.Controller
}

// Fetch the list of breeds
func (c *BreedsController) GetBreeds() {
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/breeds", nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Set("x-api-key", "DEMO-API-KEY")

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

// Fetch images for a specific breed
func (c *BreedsController) GetBreedImages() {
	breedID := c.GetString("breed_id")
	if breedID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "breed_id is required"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/images/search?limit=8&size=med&breed_id=" + breedID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Set("x-api-key", "DEMO-API-KEY")

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
