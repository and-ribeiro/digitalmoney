package routes

import (
	"dhmoney/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/cleanup", handlers.ExitApplication)

	apiGroup := router.Group("/api")

	userGroup := apiGroup.Group("/users")
	{
		userGroup.POST("/", handlers.RegisterUser)
		userGroup.POST("/login", handlers.LoginUser)
		userGroup.POST("/logout", handlers.LogoutUser)
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.PATCH("/:id", handlers.UpdateUser)
	}

	accountGroup := apiGroup.Group("/account")
	{
		accountGroup.GET("/:id", handlers.GetAccount)
		//accountGroup.PATCH("/:id)")
	}

	cardsGroup := apiGroup.Group("/accounts")
	{
		cardsGroup.GET("/:account_id/cards", handlers.GetAllCards)
		cardsGroup.POST("/:account_id/cards", handlers.CreateCard)
		cardsGroup.GET("/:account_id/cards/:card_id", handlers.GetCard)
		cardsGroup.DELETE("/:account_id/cards/:card_id", handlers.DeleteCard)
	}

	transactionGroup := apiGroup.Group("/transaction")
	{
		transactionGroup.POST("/:account_id", handlers.Deposit)
		transactionGroup.GET("/:account_id", handlers.RetrieveTransactions)
	}

}
