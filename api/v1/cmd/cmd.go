package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/InYuusha/api/v1/cmd/handler"
	m "github.com/InYuusha/api/v1/cmd/middleware"
)

func ApplyRoutes(r *gin.RouterGroup) {

	r.Use(m.ErrorMiddleware())

	cmdRoutes := r.Group("/cmd")
	{
		cmdRoutes.POST("/query", handler.ExecuteQuery)

	}
}
