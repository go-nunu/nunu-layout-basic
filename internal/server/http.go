package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-base/internal/handler"
	"github.com/go-nunu/nunu-layout-base/internal/middleware"
	"github.com/go-nunu/nunu-layout-base/pkg/helper/resp"
	"github.com/go-nunu/nunu-layout-base/pkg/log"
)

func NewServerHTTP(
	log *log.Logger,
	userHandler *handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})
	r.GET("/user", userHandler.GetUserById)

	return r
}
