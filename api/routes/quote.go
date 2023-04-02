package routes

import (
	"github.com/alireza-dehghan-nayeri/go-project/api/controller"
	"github.com/alireza-dehghan-nayeri/go-project/api/middlewares"
	"github.com/alireza-dehghan-nayeri/go-project/infrastructure"
)

// QuoteRoute -> Route for question module
type QuoteRoute struct {
	Controller controller.QuoteController
	Handler    infrastructure.GinRouter
}

// NewQuoteRoute -> initializes new choice rouets
func NewQuoteRoute(
	controller controller.QuoteController,
	handler infrastructure.GinRouter,

) QuoteRoute {
	return QuoteRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p QuoteRoute) Setup() {
	quote := p.Handler.Gin.Group("/quotes")
	{
		quote.Use(middlewares.JwtAuthMiddleware())
		quote.GET("/", p.Controller.GetQuotes)
		quote.POST("/", p.Controller.AddQuote)
		quote.GET("/:id", p.Controller.GetQuote)
		quote.DELETE("/:id", p.Controller.DeleteQuote)
		quote.PUT("/:id", p.Controller.UpdateQuote)
	}
}
