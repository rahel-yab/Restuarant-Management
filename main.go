package main

import(
	"os"
	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/database"
	"github.com/rahel-yab/Restuarant-Management/middleware"
	"github.com/rahel-yab/Restuarant-Management/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.client, "food")
func main(){

	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRouters(router)
	routes.InvoiceRouters(router)
	routes.MenuRouters(router)
	routes.OrderRouters(router)
	routes.OrderItemRouters(router)
	routes.TableRouters(router)

	router.Run(":" + port)
}
