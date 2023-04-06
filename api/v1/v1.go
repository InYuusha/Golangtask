package v1

import (
	"github.com/InYuusha/api/v1/cmd"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r*gin.RouterGroup){
	
	v1:=r.Group("/v1")
	{
		cmd.ApplyRoutes(v1)
	}
}