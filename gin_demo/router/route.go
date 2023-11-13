package router

import (
	ginUtil "gin_demo/gin_util"
	"gin_demo/logic"
	"github.com/gin-gonic/gin"
)

func SetUrl(r *gin.Engine, e *logic.LogicEntry) {
	apiHandle := r.Group("buisV3")
	{
		apiHandle.POST("test3", ginUtil.WrapBusinessCall(e.GetInfo))
	}
}
