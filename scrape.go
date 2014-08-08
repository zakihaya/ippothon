package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    //"strings"
)

func main() {
    doc, _ := goquery.NewDocument("http://stocks.finance.yahoo.co.jp/stocks/qi/?js=%E3%81%82")
    doc.Find("table.yjS tr.yjM").Each(func(_ int, s *goquery.Selection) {
        fmt.Println(s.Find(".yjMt").Text())
        price := s.Find(".price font").Text()
        if price == "" {
            price = s.Find(".price").Text()
        }
        fmt.Println(price)
    })
}
