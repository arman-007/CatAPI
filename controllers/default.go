package controllers

import (
	"CatAPI/utils"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.vip"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }

func (c *MainController) Index() {
    // Create a channel to receive data from goroutines
    ch := make(chan utils.APIResponse, 3) // Use `utils.APIResponse`

    // Launch goroutines to fetch data concurrently
    go utils.FetchData("https://api.thecatapi.com/v1/images/search?limit=1&size=med", "voting", ch)
    go utils.FetchData("https://api.thecatapi.com/v1/breeds", "breeds", ch)
    go utils.FetchData("https://api.thecatapi.com/v1/favourites?limit=20&page=0&order=Desc&sub_id=", "favorites", ch)

    // Collect data from the channel
    responseMap := make(map[string]interface{})
    for i := 0; i < 3; i++ {
        res := <-ch
        if res.Error != nil {
            responseMap[res.Key] = map[string]string{"error": res.Error.Error()}
			fmt.Println(responseMap)
        } else {
            responseMap[res.Key] = res.Data
        }
    }

	// Fetch images for each breed
    breeds, ok := responseMap["breeds"].([]map[string]interface{})
    if ok {
        for i, breed := range breeds {
            breedID := breed["id"].(string)
            images := utils.FetchBreedImages(breedID) // New function to fetch breed images
            breeds[i]["images"] = images
        }
        responseMap["breeds"] = breeds
    }

    // Pass the data to the template
    c.Data["Voting"] = responseMap["voting"]
    c.Data["Breeds"] = responseMap["breeds"]
    c.Data["Favorites"] = responseMap["favorites"]

    // Render the template
    c.TplName = "index.tpl"
}
