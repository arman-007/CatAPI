package controllers

import (
	// "encoding/json"
	"io"
	"net/http"
	"CatAPI/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	beego.Controller
}

// Fetch the list of breeds
func (c *BreedsController) GetBreedsAndImages() {
	ch := make(chan utils.APIResponse, 2)

	// Fetch breed list
	go utils.FetchData("https://api.thecatapi.com/v1/breeds", "breeds", ch, nil)

	// Fetch images for a specific breed (for example, "abys")
	go utils.FetchData("https://api.thecatapi.com/v1/images/search", "breed_images", ch, map[string]string{
		"breed_id": "abys", // Replace with the actual breed ID
		"limit":    "8",
		"size":     "med",
	})

	// Collect results
	responseMap := make(map[string]interface{})
	for i := 0; i < 2; i++ {
		res := <-ch
		if res.Error != nil {
			responseMap[res.Key] = map[string]string{"error": res.Error.Error()}
		} else {
			responseMap[res.Key] = res.Data
		}
	}

	// Serve JSON response
	c.Data["json"] = responseMap
	c.ServeJSON()
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
