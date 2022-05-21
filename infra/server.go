package infra

import(
	"github.com/kkjoker/recipe/controller"

	"github.com/gin-gonic/gin"
)

func Init() {

	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()


	routes := r.Group("/v1")
	{
		ctrl := controller.Controller{}
        routes.GET("/user", ctrl.Index)
        routes.GET("/user/:id", ctrl.Show)
        routes.POST("/user", ctrl.Create)
        routes.PUT("/user/:id", ctrl.Update)
        routes.DELETE("/user/:id", ctrl.Delete)
	}

	return r
}