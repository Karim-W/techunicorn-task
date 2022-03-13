package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/karim-w/techunicorn-task/models"
	"github.com/karim-w/techunicorn-task/service"
	"go.uber.org/fx"
)

type AuthController struct {
	service *service.AuthService
}

func (a *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if u, token, err := a.service.Register(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"details": u, "token": token})
	}
}

func (a *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if u, err := a.service.Login(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{
			"aceess_token": u,
		})
	}
}

func (a *AuthController) SetupRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", a.Register)
	rg.POST("/login", a.Login)
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

var AuthControllerModule = fx.Option(fx.Provide(NewAuthController))
