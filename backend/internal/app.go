package internal

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func Initialize(handlers *[]Handler) *App {

	// Set up gin logger
	f, _ := os.Create("log.txt")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Non-debug mode
	//gin.SetMode(gin.ReleaseMode)

	// gin app
	app := App{router: gin.New()}

	initializeMiddlewares(app.router)

	initializeControllers(app.router, handlers)

	return &app
}

func initializeMiddlewares(router *gin.Engine) {
}

func initializeControllers(router *gin.Engine, handlers *[]Handler) {
	if handlers != nil {
		for _, handler := range *handlers {
			handler.InitializeRoutes(router)
		}
	}
}

func (a *App) Start(port string) {

	if os.Getenv("ENV") == "development" {
		log.Printf("server running on http://localhost%s", port)
	}

	err := a.router.Run(port)
	if err != nil {
		log.Fatal(err)
	}

}
