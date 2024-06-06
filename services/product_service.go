package services

import (
    "saas-finance/models"
    "saas-finance/database"
)

func CreateProduct(product models.Product) (models.Product, error) {
    if err := database.DB.Create(&product).Error; err != nil {
        return models.Product{}, err
    }
    return product, nil
}

func GetProducts() ([]models.Product, error) {
    var products []models.Product
    if err := database.DB.Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}
