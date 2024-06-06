package services

import (
    "saas-finance/models"
    "saas-finance/database"
)

func CreateAccount(account models.Account) (models.Account, error) {
    if err := database.DB.Create(&account).Error; err != nil {
        return models.Account{}, err
    }
    return account, nil
}

func GetAccounts() ([]models.Account, error) {
    var accounts []models.Account
    if err := database.DB.Find(&accounts).Error; err != nil {
        return nil, err
    }
    return accounts, nil
}
