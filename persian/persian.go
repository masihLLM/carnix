package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "time"

    persian "github.com/yaa110/go-persian-calendar"
)

type Item struct {
    ID                   string  `json:"_id"`
    Title                string  `json:"title"`
    URL                  string  `json:"url"`
    CompanyName          string  `json:"companyName"`
    FundRaiseAmount      int64   `json:"fundRaiseAmount"`
    RaisedFundAmount     int64   `json:"raisedFundAmount"`
    InvestorsCount       int     `json:"investorsCount"`
    ExpectedReturnRate   float64 `json:"expectedReturnRate"`
    CollateralType       string  `json:"collateralType"`
    Duration             string  `json:"duration"`
    FundRaisingEndDate   string  `json:"fundRaisingEndDate"`
    PlatformID           string  `json:"platformID"`
    PlatformName         string  `json:"platformName"`
    PlatformCode         string  `json:"platformCode"`
    BannerURL            string  `json:"bannerURL"`
    PlatformWebsite      string  `json:"platformWebsite"`
}

func ConvertUnixToPersian(unixTimestamp int64) string {
    t := time.Unix(unixTimestamp, 0)
    p := persian.New(t)
    return fmt.Sprintf("%d/%02d/%02d", p.Year(), p.Month(), p.Day())
}

func ReadCSVAndConvert(filePath string) ([]Item, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var items []Item
    for _, record := range records[1:] { // Skip header
        unixTimestamp, err := strconv.ParseInt(record[9], 10, 64) // Assuming the Unix timestamp is in the 10th column
        if err != nil {
            return nil, err
        }
        persianDate := ConvertUnixToPersian(unixTimestamp)

        item := Item{
            ID:                 record[0],
            Title:              record[1],
            CompanyName:        record[2],
            FundRaiseAmount:    parseInt64(record[3]),
            RaisedFundAmount:   parseInt64(record[4]),
            InvestorsCount:     parseInt(record[5]),
            ExpectedReturnRate: parseFloat64(record[6]),
            CollateralType:     record[7],
            Duration:           record[8],
            FundRaisingEndDate: persianDate,
            PlatformID:         record[10],
            PlatformName:       record[11],
            PlatformCode:       record[12],
            URL:                record[13],
            BannerURL:          record[14],
            PlatformWebsite:    record[15],
        }
        items = append(items, item)
    }

    return items, nil
}

func parseInt64(s string) int64 {
    i, _ := strconv.ParseInt(s, 10, 64)
    return i
}

func parseInt(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func parseFloat64(s string) float64 {
    f, _ := strconv.ParseFloat(s, 64)
    return f
}

func WriteCSV(filePath string, items []Item) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write header
    writer.Write([]string{"ID", "Title", "Company Name", "Fund Raise Amount", "Raised Fund Amount", "Investors Count", "Expected Return Rate", "Collateral Type", "Duration", "Fund Raising End Date", "Platform ID", "Platform Name", "Platform Code", "URL", "Banner URL", "Platform Website"})

    // Write records
    for _, item := range items {
        record := []string{
            item.ID,
            item.Title,
            item.CompanyName,
            strconv.FormatInt(item.FundRaiseAmount, 10),
            strconv.FormatInt(item.RaisedFundAmount, 10),
            strconv.Itoa(item.InvestorsCount),
            strconv.FormatFloat(item.ExpectedReturnRate, 'f', 6, 64),
            item.CollateralType,
            item.Duration,
            item.FundRaisingEndDate,
            item.PlatformID,
            item.PlatformName,
            item.PlatformCode,
            item.URL,
            item.BannerURL,
            item.PlatformWebsite,
        }
        writer.Write(record)
    }

    return nil
}

func main() {
    items, err := ReadCSVAndConvert("../items927.csv")
    if err != nil {
        fmt.Printf("Error reading CSV: %v\n", err)
        return
    }

    err = WriteCSV("items927_converted.csv", items)
    if err != nil {
        fmt.Printf("Error writing CSV: %v\n", err)
        return
    }

    fmt.Println("CSV conversion completed successfully.")
}