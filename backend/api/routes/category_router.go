package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xStrato/pluralsight-backup/backend/api/controllers"
)

type CategoryRouter struct {
	ctr *controllers.CategoryController
}

func NewCategoryRouter(c *controllers.CategoryController) *CategoryRouter {
	return &CategoryRouter{c}
}

func (cr *CategoryRouter) Configure(gin *gin.Engine) error {
	main := gin.Group("api/v1")
	{
		category := main.Group("video")
		{
			category.POST("/", cr.ctr.Create)
		}
	}
	return nil
}
