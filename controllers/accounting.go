package controllers

import (
    "saas-finance/models"
    "saas-finance/services"
    "saas-finance/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateAccount(c *gin.Context) {
    var input models.Account
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    account, err := services.CreateAccount(input)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, account)
}

func GetAccounts(c *gin.Context) {
    accounts, err := services.GetAccounts()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, accounts)
}
