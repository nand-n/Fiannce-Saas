package utils

import (
    "saas-finance/models"
)

func CalculateBudgetStatus(budgets []models.Budget) map[string]float64 {
    status := make(map[string]float64)
    for _, budget := range budgets {
        status[budget.Name] = budget.Remaining / budget.Amount * 100
    }
    return status
}

func GenerateRecommendations(transactions []models.Transaction, budgets []models.Budget) string {
    // This is a placeholder for a more sophisticated recommendation system
    return "Based on your recent transactions, you might want to review your marketing budget."
}
