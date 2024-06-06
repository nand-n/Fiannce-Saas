package controllers

import (
    "saas-finance/models"
    "saas-finance/services"
    "saas-finance/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateBudget(c *gin.Context) {
    var input models.Budget
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    budget, err := services.CreateBudget(input)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, budget)
}

func GetBudgets(c *gin.Context) {
    budgets, err := services.GetBudgets()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, budgets)
}
