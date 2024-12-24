package routers

import (
	"CatAPI/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")

    beego.Router("/api/voting/cat", &controllers.VotingController{}, "get:GetCat")
	beego.Router("/api/voting/vote", &controllers.VotingController{}, "post:SubmitVote")

    beego.Router("/api/breeds", &controllers.BreedsController{}, "get:GetBreedsAndImages")
	beego.Router("/api/breeds/images", &controllers.BreedsController{}, "get:GetBreedImages")

	beego.Router("/api/favorites", &controllers.FavoritesController{}, "post:AddFavorite")
	beego.Router("/api/favorites", &controllers.FavoritesController{}, "get:GetFavorites")
}
