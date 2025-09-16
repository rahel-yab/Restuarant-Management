package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/controllers"

)

func MenuRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus" , controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id" , controllers.GetMenu())
	incomingRoutes.POST("/menus" , controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id" , controllers.UpdateMenu())

}
