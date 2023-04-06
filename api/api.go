package api

import (
	"github.com/gin-gonic/gin"
	"github.com/InYuusha/api/v1"
)

func ApplyRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}