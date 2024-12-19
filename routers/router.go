package routers

import (
	"CatAPI/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/voting", &controllers.VotingController{})
	// beego.Router("/api/vote", &controllers.VotingController{}, "post:Vote")
    beego.Router("/breeds", &controllers.BreedsController{}, "get:GetBreeds")
	beego.Router("/favs", &controllers.FavoritesController{}, "get:GetFavorites")
}
