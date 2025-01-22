package main

import (
    "fmt"
    "log"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for page := 1; page <= 1; page++ {
        wg.Add(1)
        go func(page int) {
            defer wg.Done()
            url := fmt.Sprintf("https://carinx.ir/api/fund-raises?page=%d&platforms=halalfund%%2Cdongi%%2Ckarencrowd%%2Czeema%%2Cifund%%2Chamafarin%%2Cnovincrowd%%2Ccharisma%%2Cibcrowd%%2Chamashena%%2Csepehrino%%2Cpulsar&activeOpportunitiesOnly=false", page)
            items, err := FetchData(url)
            if err != nil {
                log.Printf("Error fetching data for page %d: %v", page, err)
                return
            }
            err = SaveToCSV(items)
            if err != nil {
                log.Printf("Error saving data for page %d: %v", page, err)
            }
        }(page)
    }
    wg.Wait()
    fmt.Println("Data fetching and saving completed.")
}