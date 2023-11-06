package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"github.com/gocolly/colly"
)

type Stock struct{
company, price, change string
}

func main(){
	ticker := []string{
		"MSFT",
		"IBM",
		"GE",
		"UNP",
		"COST",
		"MCD",
		"V",
		"WMT",
		"DIS",
		"MMM",
		"INTC",
  		"AXP",
    	"AAPL",
    	"BA",
    	"CSCO",
    	"GS",
    	"JPM",
    	"CRM",
    	"VZ",
   
	}

	stocks := []Stock{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) { 
		log.Println("Something went wrong: ", err) 
	}) 

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		
		stock := Stock{}

		stock.company = e.ChildText("h1")
		fmt.Println("Company:", stock.company)
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price:", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change:", stock.change)

		stocks = append(stocks, stock)
	})
	c.Wait()

	for _, t := range ticker {
	c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}

	fmt.Println(stocks)

	file, err := os.Create("stocks.csv") 
	if err != nil { 
		log.Fatalln("Failed to create output CSV file", err) 
	} 
	defer file.Close() 
	writer := csv.NewWriter(file)
	headers := []string{
		"company", 
		"price", 
		"change", 
	} 
	writer.Write(headers) 
	for _, stock := range stocks { 
		record := []string{ 
			stock.company,
			stock.price,
			stock.change,
		} 
		writer.Write(record) 
		}
		defer writer.Flush()
}