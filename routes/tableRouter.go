package routes


import(
	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/controllers"

)

func TableRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables" , controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id" , controllers.GetTable())
	incomingRoutes.POST("/tables" , controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:Table_id" , controllers.UpdateTable())

}
