package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/controllers"

)

func FoodRoutes(incominRoutes *gin.Engine){
	incominRoutes.GET("/foods" , controllers.GetFoods())
	incominRoutes.GET("/foods/:food_id" , controllers.GetFood())
	incominRoutes.POST("/foods" , controllers.CreateFood())
	incominRoutes.PATCH("/foods/:food_id" , controllers.UpdateFood())

}