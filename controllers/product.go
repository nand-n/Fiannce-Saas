package controllers

import (
    "saas-finance/models"
    "saas-finance/services"
    "saas-finance/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    product, err := services.CreateProduct(input)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
    products, err := services.GetProducts()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, products)
}
