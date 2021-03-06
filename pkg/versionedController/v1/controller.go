package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
	Patch(c *gin.Context)
	Register(router *gin.Engine)
}

func RegisterController(router *gin.Engine, controller Controller, routePath string) {
	router.GET(routePath, controller.Get)
	router.POST(routePath, controller.Post)
	router.PUT(routePath, controller.Put)
	router.DELETE(routePath, controller.Delete)
	router.PATCH(routePath, controller.Patch)
}

type GenericController struct {
	Controller
}

func (genericController *GenericController) Get(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNotFound)
}

func (genericController *GenericController) Post(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNotFound)
}

func (genericController *GenericController) Put(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNotFound)
}

func (genericController *GenericController) Delete(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNotFound)
}

func (genericController *GenericController) Patch(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNotFound)
}

func (genericController *GenericController) ResponseJSON(c *gin.Context, status int, response interface{}) {
	c.JSON(status, response)
}

func (genericController *GenericController) ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, map[string]string{"error": message})
}
