package main

import (
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
)

type Person struct {
    Id int32
    Name string
}

func main() {
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        panic(err.Error())
    }
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
    t := dbmap.AddTableWithName(Person{}, "person").SetKeys(true, "Id")
    t.ColMap("Id").Rename("id")
    t.ColMap("Name").Rename("name")
    dbmap.DropTables()
    err = dbmap.CreateTables()
    if err != nil {
        panic(err.Error())
    }

    tx, _ := dbmap.Begin()
    for i := 0; i < 100; i++ {
        tx.Insert(&Person{0, fmt.Sprintf("mattn%03d", i)})
    }
    tx.Commit()

    list, _ := dbmap.Select(Person{}, "select * from person")
    for _, l := range list {
        p := l.(*Person)
        fmt.Printf("%d, %s\n", p.Id, p.Name)
    }
}
