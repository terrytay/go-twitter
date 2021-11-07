package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terrytay/twitter/backend/entities"
	"github.com/terrytay/twitter/backend/services"
	"github.com/terrytay/twitter/backend/tools"
)

type UserHandler struct {
	UserService services.IUserService
}

func (handler *UserHandler) InitializeRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/", handler.post)
	}
}

func (u *UserHandler) post(c *gin.Context) {
	user := entities.NewUser{}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		tools.ToJSON(gin.H{"success": false, "message": err.Error()}, c.Writer)
		return
	}

	err = u.UserService.Create(c, user)
	if err != nil {
		tools.ToJSON(gin.H{"success": false, "message": err.Error()}, c.Writer)
		return
	}

	c.Status(http.StatusOK)
}
