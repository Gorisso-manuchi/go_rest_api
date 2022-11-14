package Routes

import (
	"restapi/restApi/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("/page-api")
	{
		grpl.GET("page", Controllers.GetUsers)
		grpl.POST("page", Controllers.CreateUser)
		grpl.GET("page/:id", Controllers.GetUserByID)
		grpl.PUT("page/:id", Controllers.UpdateUser)
		grpl.DELETE("page/:id", Controllers.DeleteUser)

	}
	return r
}
