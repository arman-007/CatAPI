package main

import (
	"encoding/json"
	_ "CatAPI/routers"
	"log"
	"github.com/beego/beego/v2/server/web/filter/cors"
	beego "github.com/beego/beego/v2/server/web"
)


// jsonFunc serializes data into a JSON string
func jsonFunc(data interface{}) string {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Println("Error serializing data to JSON:", err)
        return ""
    }
    return string(jsonData)
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:8080"}, // Replace with your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Register the json function for templates
    beego.AddFuncMap("json", jsonFunc)
	beego.Run()
}

