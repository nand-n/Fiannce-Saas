package main

import (
    "saas-finance/config"
    "saas-finance/routes"
    "saas-finance/database"
    "saas-finance/utils"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Initialize logger
    utils.InitLogger()
    utils.Log.Info("Logger initialized")

    // Initialize configuration
    config.LoadConfig()
    utils.Log.Info("Configuration loaded")

    // Initialize database
    database.ConnectDatabase()
    utils.Log.Info("Database connected")

    // Setup routes
    routes.SetupRoutes(r)
    utils.Log.Info("Routes set up")

    // Run the server
    utils.Log.Info("Starting server on :8080")
    r.Run(":8080")
}
