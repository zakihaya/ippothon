package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
    "strconv"
    //"strings"
)

type Stock struct {
    Id int32
    CompanyName string
    Price string
}

func main() {
    db, err := sql.Open("sqlite3", "./scrape.db")
    if err != nil {
        panic(err.Error())
    }
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
    t := dbmap.AddTableWithName(Stock{}, "stock").SetKeys(true, "Id")
    t.ColMap("Id").Rename("id")
    dbmap.DropTables()
    err = dbmap.CreateTables()
    if err != nil {
        panic(err.Error())
    }

    tx, _ := dbmap.Begin()

    baseUrl := "http://stocks.finance.yahoo.co.jp/stocks/qi/?&p="
    for i := 1; i <= 179; i++ {
      pageString := strconv.Itoa(i)
      pageUrl := baseUrl + pageString
      fmt.Println(pageUrl)
      doc, _ := goquery.NewDocument(pageUrl)
      doc.Find("table.yjS tr.yjM").Each(func(_ int, s *goquery.Selection) {
          companyName := s.Find(".yjMt").Text()
          price := s.Find(".price font").Text()
          if price == "" {
              price = s.Find(".price").Text()
          }
          tx.Insert(&Stock{0, companyName, price})
      })
    }

    tx.Commit()

    list, _ := dbmap.Select(Stock{}, "select * from stock")
    fmt.Println(len(list))
}
