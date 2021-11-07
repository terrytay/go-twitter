package internal

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
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
	// gin.SetMode(gin.ReleaseMode)

	// gin app
	app := App{router: gin.New()}

	initializeMiddlewares(app.router)

	initializeControllers(app.router, handlers)

	return &app
}

func initializeMiddlewares(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
}

func initializeControllers(router *gin.Engine, handlers *[]Handler) {
	if handlers != nil {
		for _, handler := range *handlers {
			handler.InitializeRoutes(router)
		}
	}
}

func (a *App) Start(port string, l *log.Logger) {

	if os.Getenv("ENV") == "development" {
		log.Printf("server running on http://localhost%s", port)
	}

	s := &http.Server{
		Addr:         port,     // configure the bind address
		Handler:      a.router, // set the default handler
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second, // max time for connections using TCP keep-alive
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Listen to interrupt signal on channel
	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)

	// Shutdown gracefully upon receiving SIGINT
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// timeout context to pass to shutdown process
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(tc)

}
