package models

import "gorm.io/gorm"

type Transaction struct {
    gorm.Model
    Amount       int64
    Currency     string
    Description  string
    PaymentID    string
}
