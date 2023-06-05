package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-base/internal/service"
	"github.com/go-nunu/nunu-layout-base/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	*Handler
	userService *service.UserService
}

func NewUserHandler(handler *Handler, userService *service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (c *UserHandler) GetUserById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := c.userService.GetUserById(params.Id)
	c.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}
func (c *UserHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
