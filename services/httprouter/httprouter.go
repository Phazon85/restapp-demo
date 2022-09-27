package httprouter

import (
	"net/http"

	"github.com/Phazon85/restapp-demo/services/handlers/todos"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todos", todos.GetTodoHandler)

	return r
}
