package services

import (
    "saas-finance/config"
    "saas-finance/models"
    "saas-finance/database"
    "github.com/stripe/stripe-go/v72"
    "github.com/stripe/stripe-go/v72/charge"
)

func init() {
    config.LoadConfig()
    stripe.Key = config.GetStripeSecretKey()
}

func CreateStripeCharge(amount int64, currency, description, source string) (*stripe.Charge, error) {
    params := &stripe.ChargeParams{
        Amount:       stripe.Int64(amount),
        Currency:     stripe.String(currency),
        Description:  stripe.String(description),
    }
    params.SetSource(source)

    ch, err := charge.New(params)
    if err != nil {
        return nil, err
    }

    return ch, nil
}

func SaveTransaction(transaction *models.Transaction) {
    database.DB.Create(&transaction)
}
