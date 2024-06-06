package utils

import (
    "saas-finance/models"
    "encoding/csv"
    "os"
)

func GenerateTransactionCSV(transactions []models.Transaction, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write CSV header
    writer.Write([]string{"ID", "Amount", "Currency", "Description", "PaymentID"})

    // Write CSV records
    for _, transaction := range transactions {
        writer.Write([]string{
            string(transaction.ID),
            string(transaction.Amount),
            transaction.Currency,
            transaction.Description,
            transaction.PaymentID,
        })
    }

    return nil
}
