package main

import(
	"os"
	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/database"
	"github.com/rahel-yab/Restuarant-Management/middleware"
	"github.com/rahel-yab/Restuarant-Management/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func main(){

	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	
}
