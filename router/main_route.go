package router

import (
	"let-you-cook/config"
	"let-you-cook/handler"
	"let-you-cook/repository"
	"let-you-cook/service"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

func SetupRouter() *gin.Engine {
	route := gin.Default()

	apiV1 := route.Group("/api/v1")

	db := config.ConnectDatabase()

	// index repository
	indexRepo := repository.NewIndexRepo(db)

	// auth route
	authRepo := repository.NewAuthRepository(db, indexRepo)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	// user route
	userRepo := repository.NewUserRepo(db, indexRepo)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// user route
	profileRepo := repository.NewProfileRepo(db, indexRepo)
	profileService := service.NewProfileService(profileRepo, userRepo)
	profileHandler := handler.NewProfileHanlder(profileService)

	// task SetupRouter
	taskRepo := repository.NewTaskRepo(db, indexRepo)
	taskService := service.NewTaskService(taskRepo, userRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// category SetupRouter
	categoryRepo := repository.NewCategoryRepo(db, indexRepo)
	categoryService := service.NewCategoryService(categoryRepo, userRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// session SetupRouter
	sessionRepo := repository.NewSessionRepo(db, indexRepo)
	sessionService := service.NewSessionService(sessionRepo, userRepo)
	sessionHandler := handler.NewSessionHandler(sessionService)

	AuthRoute(apiV1, authHandler)
	UserRoute(apiV1, userHandler)
	ProfileRoute(apiV1, profileHandler)
	TaskRoute(apiV1, taskHandler)
	CategoryRoute(apiV1, categoryHandler)
	SessionRoute(apiV1, sessionHandler)

	return route
}
