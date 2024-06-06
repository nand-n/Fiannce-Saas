package controllers

import (
    "saas-finance/models"
    "saas-finance/database"
    "saas-finance/utils"
    "github.com/gin-gonic/gin"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/charge"
    "net/http"
)

func CreateCharge(c *gin.Context) {
    var input struct {
        Amount      int64  `json:"amount" binding:"required"`
        Currency    string `json:"currency" binding:"required"`
        Description string `json:"description" binding:"required"`
        Source      string `json:"source" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        utils.Log.Error("Invalid input: ", err)
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
        return
    }

    chargeParams := &stripe.ChargeParams{
        Amount:      stripe.Int64(input.Amount),
        Currency:    stripe.String(input.Currency),
        Description: stripe.String(input.Description),
    }
    chargeParams.SetSource(input.Source)

    ch, err := charge.New(chargeParams)
    if err != nil {
        utils.Log.Error("Failed to create charge: ", err)
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create charge")
        return
    }

    transaction := models.Transaction{
        Amount:      ch.Amount,
        Currency:    string(ch.Currency),
        Description: ch.Description,
        PaymentID:   ch.ID,
    }

    if err := database.DB.Create(&transaction).Error; err != nil {
        utils.Log.Error("Failed to save transaction: ", err)
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to save transaction")
        return
    }

    utils.Log.Info("Charge created: ", transaction)
    c.JSON(http.StatusOK, transaction)
}
