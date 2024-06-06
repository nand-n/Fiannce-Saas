package config

import (
    "os"
    "github.com/joho/godotenv"
)

func LoadConfig() {
    godotenv.Load()
}

func GetStripeSecretKey() string {
    return os.Getenv("STRIPE_SECRET_KEY")
}
