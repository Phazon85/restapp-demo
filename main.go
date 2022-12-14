package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/Phazon85/restapp-demo/docs"
	todoHandler "github.com/Phazon85/restapp-demo/handlers/todos"
	groupcacheService "github.com/Phazon85/restapp-demo/services/groupcache"
	todoService "github.com/Phazon85/restapp-demo/services/todos"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mailgun/groupcache"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.uber.org/zap"
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
	defaultHost         = "localhost"
	defaultPort         = "5432"
	defaultUser         = "postgres"
	defaultPassword     = "changeme"
	defaultDBName       = "restapp-demo"
	defaultSSLMode      = "disable"
	defaultDriverName   = "postgres"
	defaultCacheSize    = 1024 * 1024
	sqlConnectionString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	errorExitCode       = 1
)

var (
	host       = flag.String("host", lookupEnv("HOST", defaultHost), "host")
	port       = flag.String("port", lookupEnv("PORT", defaultPort), "port")
	user       = flag.String("user", lookupEnv("USER", defaultUser), "user")
	password   = flag.String("password", lookupEnv("PASSWORD", defaultPassword), "password")
	dbName     = flag.String("dbname", lookupEnv("DB_NAME", defaultDBName), "dbname")
	driverName = flag.String("driverName", lookupEnv("DRIVER_NAME", defaultDriverName), "driverName")
	sslMode    = flag.String("sslMode", lookupEnv("SSL_MODE", defaultSSLMode), "sslMode")
)

func lookupEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	// Establish repository connection.
	sqlConn, err := sql.Open(*driverName, fmt.Sprintf(sqlConnectionString, *host, *port, *user, *password, *dbName, *sslMode))
	if err != nil {
		logger.Fatal("sql.Open - Failed to create SQL object: ", zap.Error(err))
	}

	// Check is SQL successfully connected.
	if err := sqlConn.Ping(); err != nil {
		logger.Fatal("sql.Ping - Failed to connect to SQL object: ", zap.Error(err))
	}

	// Instantiate services.
	todoServ := todoService.New(logger, sqlConn)

	// Instantiate handlers.
	todoHandler := todoHandler.New(todoServ)

	// Instantiate groupcache service.
	groupcacheService, err := groupcacheService.New(logger)
	if err != nil {
		logger.Fatal("groupcacheService.New - Failed to instantiate HTTPPool: ", zap.Error(err))
	}

	groupcacheService.Pool = groupcache.NewHTTPPoolOpts(groupcacheService.Address, &groupcache.HTTPPoolOptions{})

	//Spawn groupcache peering server
	// Pool keeps track of peers in our cluster and identifies which peer owns a key.
	groupcacheService.SetPeers()

	//Create groupcache groups to serve.
	testGroup := groupcache.NewGroup("test", defaultCacheSize, groupcacheService.NewTodoGetter(todoServ))

	// Create new Gin Engine.
	r := gin.Default()

	// Gin Routes.
	v1Group := r.Group("/api/v1")
	{
		// Todos router group
		todos := v1Group.Group("/todos")
		{
			todos.GET("", todoHandler.Get)
			todos.GET("/:key", todoHandler.TestGetByID(testGroup))
			todos.POST("", todoHandler.Post)
			todos.DELETE("/:id", todoHandler.Delete)
			todos.PUT("/:id", todoHandler.Put)
		}

		// base v1 routes
		v1Group.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}

	// Groupcache router group
	groupcacheGroup := r.Group("/_groupcache")
	{
		groupcacheGroup.GET("", func(c *gin.Context) { groupcacheService.Pool.ServeHTTP(c.Writer, c.Request) })

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run Gin server.
	log.Print("Server Starting")
	r.Run(":8080")
}
