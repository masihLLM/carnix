package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Item struct {
    ID                   string `json:"_id"`
    Title                string `json:"title"`
    URL                  string `json:"url"`
    CompanyName          string `json:"companyName"`
    FundRaiseAmount      int64  `json:"fundRaiseAmount"`
    RaisedFundAmount     int64  `json:"raisedFundAmount"`
    InvestorsCount       int    `json:"investorsCount"`
    ExpectedReturnRate   int    `json:"expectedReturnRate"`
    CollateralType       string `json:"collateralType"`
    Duration             string `json:"duration"`
    FundRaisingEndDate   string `json:"fundRaisingEndDate"`
    BannerURL            string `json:"bannerURL"`
    Platform             Platform `json:"platform"`
}

type Platform struct {
    ID      string `json:"_id"`
    Name    string `json:"name"`
    Code    string `json:"code"`
    Website string `json:"website"`
}

type Response struct {
    Items []Item `json:"items"`
    Count int    `json:"count"`
    Total int    `json:"total"`
}

func FetchData(url string) ([]Item, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch data: %w", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    var response Response
    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    return response.Items, nil
}