package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/router"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/router/middleware"
)

func InitServer() *gin.Engine {
	r := gin.New()
	if r == nil {
		panic("load gin error")
	}

	gin.SetMode(gin.DebugMode)

	// 自定义错误处理
	//r.Use(middleware.CustomError)
	// NoCache is a middleware function that appends headers
	r.Use(middleware.NoCache)
	// 跨域处理
	r.Use(middleware.Options)

	// the jwt middleware
	authMid := middleware.JWTAuth()

	// 注册业务路由
	router.ExecRouter(r, authMid)

	return r
}
