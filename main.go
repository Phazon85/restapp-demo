package main

import (
	_ "github.com/Phazon85/restapp-demo/docs"
	"github.com/Phazon85/restapp-demo/services/httprouter"
	_ "github.com/swaggo/files" // swagger embed files
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
	//Create Gin Router with Handlers attached
	r := httprouter.New()

	//Run Gin server
	r.Run(":8080")
}
