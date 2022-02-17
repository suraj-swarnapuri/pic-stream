package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	router := gin.Default()
	router.GET("/", index)
	return router
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "hello-world!")
}
