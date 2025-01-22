package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "sync"
)

var mu sync.Mutex

func SaveToCSV(items []Item) error {
    mu.Lock()
    defer mu.Unlock()

    file, err := os.OpenFile("items.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open CSV file: %w", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, item := range items {
        record := []string{
            item.ID,
            item.Title,
            item.URL,
            item.CompanyName,
            fmt.Sprintf("%d", item.FundRaiseAmount),
            fmt.Sprintf("%d", item.RaisedFundAmount),
            fmt.Sprintf("%d", item.InvestorsCount),
            fmt.Sprintf("%d", item.ExpectedReturnRate),
            item.CollateralType,
            item.Duration,
            item.FundRaisingEndDate,
            item.BannerURL,
            item.Platform.ID,
            item.Platform.Name,
            item.Platform.Code,
            item.Platform.Website,
        }
        if err := writer.Write(record); err != nil {
            return fmt.Errorf("failed to write record to CSV: %w", err)
        }
    }

    return nil
}