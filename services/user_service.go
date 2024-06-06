package services

import (
    "saas-finance/models"
    "saas-finance/database"
    "golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (models.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return models.User{}, err
    }
    user.Password = string(hashedPassword)
    
    if err := database.DB.Create(&user).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}
