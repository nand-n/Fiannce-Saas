package services

import (
    "saas-finance/models"
    "saas-finance/database"
)

func CreateBudget(budget models.Budget) (models.Budget, error) {
    if err := database.DB.Create(&budget).Error; err != nil {
        return models.Budget{}, err
    }
    return budget, nil
}

func GetBudgets() ([]models.Budget, error) {
    var budgets []models.Budget
    if err := database.DB.Find(&budgets).Error; err != nil {
        return nil, err
    }
    return budgets, nil
}
