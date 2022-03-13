package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/karim-w/techunicorn-task/service"
	"go.uber.org/fx"
)

type DocController struct {
	service *service.DocService
}

func (d *DocController) GetDocs(c *gin.Context) {
	if doc, err := d.service.GetAll(10, 0, ""); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, doc)
	}
}

func (d *DocController) GetOne(c *gin.Context) {
	if doc, err := d.service.GetOne(c.Param("id")); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, doc)
	}
}

func (a *DocController) SetupRoutes(rg *gin.RouterGroup) {
	rg.GET("/", a.GetDocs)
	rg.GET("/:id", a.GetOne)
}

func NewDocController(service *service.DocService) *DocController {
	return &DocController{
		service: service,
	}
}

var DocControllerModule = fx.Option(fx.Provide(NewDocController))
