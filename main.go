package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() { 
    for page := 1; page <= 48; page++ {
       
            url := fmt.Sprintf("https://carinx.ir/api/fund-raises?page=%d&platforms=halalfund%%2Cdongi%%2Ckarencrowd%%2Czeema%%2Cifund%%2Chamafarin%%2Cnovincrowd%%2Ccharisma%%2Cibcrowd%%2Chamashena%%2Csepehrino%%2Cpulsar&activeOpportunitiesOnly=false", page)
            items, err := FetchData(url)
            if err != nil {
                log.Printf("Error fetching data for page %d: %v", page, err)
                return
            }
            err = saveToCSV(items)
            if err != nil {
                log.Printf("Error saving data for page %d: %v", page, err)
            }

            
            sleepDuration := time.Duration(3+rand.Intn(8)) * time.Second
            fmt.Printf("Sleeping for %v at %v (page %d)\n", sleepDuration, time.Now().Format(time.RFC3339), page)

            time.Sleep(sleepDuration)

    
    } 
    fmt.Println("Data fetching and saving completed.")
}