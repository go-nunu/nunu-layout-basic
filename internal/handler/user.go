package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout/internal/model"
	"github.com/go-nunu/nunu-layout/internal/service"
	"github.com/go-nunu/nunu-layout/pkg/log"
	"github.com/go-nunu/nunu-layout/pkg/resp"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
	log         *log.Logger
}

func NewUserHandler(log *log.Logger, sf *sonyflake.Sonyflake, userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         log,
	}
}

func (c *UserHandler) CreateUser(ctx *gin.Context) {

	var params struct {
		Username string `json:"username" binding:"required,min=2,max=20"`
		Email    string `json:"email" binding:"required,email"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := c.userService.CreateUser(&model.User{
		Username: params.Username,
		Email:    params.Email,
	})
	c.log.Logger.Info("CreateUser", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
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
	c.log.Logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}
func (c *UserHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}