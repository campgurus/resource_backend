package Routes

import (
	"github.com/gin-gonic/gin"
	"resource-api/Controllers"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/clients")
	{
		grp1.GET("client", Controllers.GetClients)
		grp1.POST("client", Controllers.CreateClient)
		grp1.GET("client/:id", Controllers.GetClientByID)
		grp1.PUT("client/:id", Controllers.UpdateClient)
		grp1.DELETE("client/:id", Controllers.DeleteClient)
	}
	return r
}
