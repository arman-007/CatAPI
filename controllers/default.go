package controllers

import (
	"CatAPI/utils"
    // "fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type FetchDataFunc func(url string, key string, ch chan<- utils.APIResponse, extraParams map[string]string)

type MainController struct {
	beego.Controller
	FetchData FetchDataFunc // Injectable FetchData function
}

func (c *MainController) Index() {
	// Default to utils.FetchData if no custom FetchData function is provided
	if c.FetchData == nil {
		c.FetchData = utils.FetchData
	}

    
	// fmt.Println("API KEY IS: ", apiKey)

	// Create a channel to receive data from goroutines
	ch := make(chan utils.APIResponse, 3)

	// Launch goroutines to fetch data concurrently
	go c.FetchData("https://api.thecatapi.com/v1/images/search?limit=1&size=med", "voting", ch, nil)
	go c.FetchData("https://api.thecatapi.com/v1/breeds", "breeds", ch, nil)
	go c.FetchData("https://api.thecatapi.com/v1/favourites?limit=20&page=0&order=Desc&sub_id=", "favorites", ch, nil)

	// Collect data from the channel
	responseMap := make(map[string]interface{})
	for i := 0; i < 3; i++ {
		res := <-ch
		if res.Error != nil {
			responseMap[res.Key] = map[string]string{"error": res.Error.Error()}
		} else {
			responseMap[res.Key] = res.Data
		}
	}

	// Pass the data to the template
	c.Data["Voting"] = responseMap["voting"]
	c.Data["Breeds"] = responseMap["breeds"]
	c.Data["Favorites"] = responseMap["favorites"]

	// Render the template
	c.TplName = "index.tpl"
}
