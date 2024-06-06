package services

import (
    "saas-finance/models"
    "saas-finance/database"
)

func GenerateTransactionReport() ([]models.Transaction, error) {
    var transactions []models.Transaction
    if err := database.DB.Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}
