package main

import (
	"log"

	"github.com/kkjoker/recipe/handler"
	"github.com/kkjoker/recipe/infra"
	"github.com/kkjoker/recipe/service"


	"github.com/gin-gonic/gin"
)

func main() {
	// 1.engineを作成します
	engine := infra.DBInit()

	// 2.サービスを作成します
	factory := service.NewService(engine)

	// 3.最後にengineを閉めます
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()

	// 4.Ginを作成します
	g := gin.Default()

	// 5.サービス層のMiddlewareを作成します
	g.Use(service.ServiceFactoryMiddleware(factory))

	// 6.v1というグループを作成します
	routes := g.Group("/v1")

	// 7.ルーティングです
	{
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)
	}

	g.Run(":3000")
}