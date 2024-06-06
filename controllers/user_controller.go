package controllers

import (
    "saas-finance/models"
    "saas-finance/services"
    "saas-finance/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    user, err := services.CreateUser(input)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, user)
}
