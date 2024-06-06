package models

import "gorm.io/gorm"

type Budget struct {
    gorm.Model
    Name        string
    Amount      float64
    Spent       float64
    Remaining   float64
    Description string
}
