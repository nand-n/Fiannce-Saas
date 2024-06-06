package controllers

import (
    "saas-finance/services"
    "saas-finance/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GenerateTransactionReport(c *gin.Context) {
    report, err := services.GenerateTransactionReport()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, report)
}
