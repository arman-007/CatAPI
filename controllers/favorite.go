package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type FavoritesController struct {
	web.Controller
}

func (c *FavoritesController) GetFavorites() {
	// Get the sub_id from query parameters or use a default one
	subID := c.GetString("sub_id", "demo-default")

	// Fetch favorites from The Cat API
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/favourites?limit=20&page=0&order=Desc&sub_id="+subID, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Set("x-api-key", "DEMO-API-KEY") // Replace with your API key

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to fetch favorites"}
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

	var favorites []map[string]interface{}
	if err := json.Unmarshal(body, &favorites); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to parse JSON"}
		c.ServeJSON()
		return
	}

	// Pass the favorites to the template
	c.Data["Favorites"] = favorites
	c.TplName = "favs.tpl"
}
