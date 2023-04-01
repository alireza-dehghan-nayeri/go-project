package main

import (
	"github.com/alireza-dehghan-nayeri/go-project/api/controller"
	"github.com/alireza-dehghan-nayeri/go-project/api/repository"
	"github.com/alireza-dehghan-nayeri/go-project/api/routes"
	"github.com/alireza-dehghan-nayeri/go-project/api/service"
	"github.com/alireza-dehghan-nayeri/go-project/infrastructure"
	"github.com/alireza-dehghan-nayeri/go-project/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()

	quoteRepository := repository.NewQuoteRepository(db)
	quoteService := service.NewQuoteService(quoteRepository)
	quoteController := controller.NewQuoteController(quoteService)
	quoteRoute := routes.NewQuoteRoute(quoteController, router)
	quoteRoute.Setup()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.Quote{}, &models.User{})

	router.Gin.Run(":8000")
}
