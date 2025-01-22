package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "sync"
)

var mu sync.Mutex

func saveToCSV(items []Item) error {
    mu.Lock()
    defer mu.Unlock()

    file, err := os.OpenFile("items927.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open CSV file: %w", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Check if the file is empty to write the header
    fileInfo, err := file.Stat()
    if err != nil {
        return fmt.Errorf("failed to get file info: %w", err)
    }
    if fileInfo.Size() == 0 {
        header := []string{
            "ID", "Title", "Company Name", "Fund Raise Amount", "Raised Fund Amount",
            "Investors Count", "Expected Return Rate", "Collateral Type", "Duration",
            "Fund Raising End Date", "Platform ID", "Platform Name", "Platform Code",
            "URL", "Banner URL", "Platform Website",
        }
        if err := writer.Write(header); err != nil {
            return fmt.Errorf("failed to write header to CSV: %w", err)
        }
    }

    for _, item := range items {
        record := []string{
            item.ID,
            item.Title,
            item.CompanyName,
            fmt.Sprintf("%d", item.FundRaiseAmount),
            fmt.Sprintf("%d", item.RaisedFundAmount),
            fmt.Sprintf("%d", item.InvestorsCount),
            fmt.Sprintf("%f", item.ExpectedReturnRate),
            item.CollateralType,
            item.Duration,
            fmt.Sprintf("%d", item.FundRaisingEndDate.Unix()),
            item.Platform.ID,
            item.Platform.Name,
            item.Platform.Code,
            item.URL,
            item.BannerURL,
            item.Platform.Website,
        }
        if err := writer.Write(record); err != nil {
            return fmt.Errorf("failed to write record to CSV: %w", err)
        }
    }

    return nil
}