package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// @title          Rest App Demo
// @version        1.0
// @description    This is a sample server rest server using the Gin Router
// @termsOfService http://swagger.io/terms/

// @contact.email justin@phazon.app

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1

const (
	sqlConnectionString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	defaultFileName     = "/etc/dbconfig"
	errorExitCode       = 1
)

type configuration struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	DriverName string `json:"driverName"`
	DBName     string `json:"dbName"`
	SSLMode    string `json:"sslMode"`
}

func main() {
	// Initialize configuration
	cfg := &configuration{}
	data, err := ioutil.ReadFile(defaultFileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatal(err)
	}

	// Establish repository connection.
	sqlConn, err := sql.Open(cfg.DriverName, fmt.Sprintf(sqlConnectionString, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		log.Fatal(err)
	}

	// if err := sqlConn.Ping(); err != nil {
	// 	log.Fatal(err)
	// }

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
