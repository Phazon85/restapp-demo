package httprouter

import (
	"net/http"

	"github.com/Phazon85/restapp-demo/services/controllers/todos"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func New() *gin.Engine {
	// Create new Gin Engine.
	r := gin.Default()

	// Instantiate Handlers.
	todosController := todos.NewController()

	// Gin Routes.
	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.GET("", todosController.GetTodos)
		}
		v1.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
