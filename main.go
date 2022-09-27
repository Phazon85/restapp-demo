package main

import (
	"database/sql"
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
	sqlConnectionString string = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	errorExitCode              = 1
)

// @title           Rest App Demo
// @version         1.0
// @description     This is a sample server rest server using the Gin Router
// @termsOfService  http://swagger.io/terms/

// @contact.email  justin@phazon.app

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	// Setup interrupt handler to gracefully shut down server.
	// errc := make(chan error)
	// go func() {
	// 	c := make(chan os.Signal, errorExitCode)
	// 	errc <- fmt.Errorf("%s", <-c)
	// 	log.Println("channel error")
	// }()

	//Establish repository connection.
	sqlConn, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Voltage13-2 dbname=restapp-demo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//Create services. //TODO: setup business logic for each handler
	todoService := todoService.New(sqlConn)

	// Instantiate Handlers.
	todoHandler := todoHandler.New(todoService)

	// Create new Gin Engine.
	r := gin.Default()

	// Gin Routes.
	v1 := r.Group("/api/v1")
	{
		//Todos router group
		todos := v1.Group("/todos")
		{
			todos.GET("", todoHandler.GetTodos)
		}

		//base v1 routes
		v1.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Run Gin server.
	r.Run(":8080")
}
