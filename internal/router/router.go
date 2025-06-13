package router

import (
	"github.com/MukizuL/hezzl-test/internal/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func newRouter(c *controller.Controller) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/ping", c.Ping)

	router.POST("/good/create", c.CreateGoods)
	router.PATCH("/good/update", c.UpdateGoods)
	router.DELETE("/good/remove", c.RemoveGoods)
	router.GET("/good/list", c.GetGoods)
	router.PATCH("/good/reprioritize", c.ReprioritizeGoods)

	return router
}

func Provide() fx.Option {
	return fx.Provide(newRouter)
}
