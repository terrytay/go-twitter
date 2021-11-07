package internal

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/terrytay/twitter/backend/handlers"
	"github.com/terrytay/twitter/backend/infrastructures/db"
	"github.com/terrytay/twitter/backend/repositories"
	"github.com/terrytay/twitter/backend/services"
)

type Repositories struct {
	userRepository repositories.IUserRepository
}

type Services struct {
	userService services.IUserService
}

type Handler interface {
	InitializeRoutes(router *gin.Engine)
}

func getRepositories(db *db.Database) *Repositories {
	return &Repositories{
		userRepository: &repositories.UserRepository{db},
	}
}

func getServices(repos *Repositories) *Services {
	return &Services{
		userService: &services.UserService{repos.userRepository},
	}
}

func getHandlers(services *Services) *[]Handler {
	return &[]Handler{
		&handlers.UserHandler{UserService: services.userService},
	}
}

func InitApp(db *db.Database) {
	repos := getRepositories(db)
	services := getServices(repos)
	handlers := getHandlers(services)

	app := Initialize(handlers)
	port := os.Getenv("PORT")

	app.Start(port)
}
