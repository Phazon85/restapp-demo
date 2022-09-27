package httprouter

import (
	"net/http"

	"github.com/Phazon85/restapp-demo/services/controllers/todos"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	todosController := todos.NewController()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todos", todosController.GetTodos)

	return r
}
