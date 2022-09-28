package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/Phazon85/restapp-demo/docs"
	todoHandler "github.com/Phazon85/restapp-demo/handlers/todos"
	todoService "github.com/Phazon85/restapp-demo/services/todos"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

const (
	sqlConnectionString = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	defaultSQLDriver    = "postgres"
	defaultHost         = "localhost"
	defaultPort         = 5432
	defaultUser         = "postgres"
	defaultPassword     = "Voltage13-2"
	defaultDBName       = "restapp-demo"
	errorExitCode       = 1
)

var host = flag.String("host", "test", "host")

// @title          Rest App Demo
// @version        1.0
// @description    This is a sample server rest server using the Gin Router
// @termsOfService http://swagger.io/terms/

// @contact.email justin@phazon.app

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1

func main() {
	// Establish repository connection.
	sqlConn, err := sql.Open(defaultSQLDriver, fmt.Sprintf(sqlConnectionString, defaultHost, defaultPort, defaultUser, defaultPassword, defaultDBName))
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate services.
	todoService := todoService.New(sqlConn)

	// Instantiate Handlers.
	todoHandler := todoHandler.New(todoService)

	// Create new Gin Engine.
	r := gin.Default()

	// Gin Routes.
	v1 := r.Group("/api/v1")
	{
		// Todos router group
		todos := v1.Group("/todos")
		{
			todos.GET("", todoHandler.Get)
			todos.POST("", todoHandler.Post)
			todos.DELETE("/:id", todoHandler.Delete)
			todos.PUT("/:id", todoHandler.Put)
		}

		// base v1 routes
		v1.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run Gin server.
	r.Run(":8080")
}
