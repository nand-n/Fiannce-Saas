package routes

import (
    "saas-finance/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/charge", controllers.CreateCharge)
    r.POST("/budget", controllers.CreateBudget)
    r.GET("/budgets", controllers.GetBudgets)
    r.POST("/product", controllers.CreateProduct)
    r.GET("/products", controllers.GetProducts)
    r.GET("/transaction/report", controllers.GenerateTransactionReport)
    r.POST("/account", controllers.CreateAccount)
    r.GET("/accounts", controllers.GetAccounts)
    r.POST("/register", controllers.RegisterUser)
}
