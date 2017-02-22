package handler

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "github.com/mgonzo/venues/model"
)

func Collection(w http.ResponseWriter, r *http.Request) {
  // connect to db
  db, err := sql.Open("mysql",
  "calendar:hopeitsok@tcp(172.17.0.1:3306)/calendar")

  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // retrieve collection
  // need a package to filter, sort and page results
  // need a real Venue struct in a model file

  var (
    id int
    name string
  )

  rows, err := db.Query("select id, name from venue")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  var list []model.Venue

  for rows.Next() {
    err := rows.Scan(&id, &name)
    if err != nil {
      log.Fatal(err)
    }

    var ven = model.Venue{
      Id: id,
      Name: name,
    }

    list = append(list, ven)

  }

  err = rows.Err()
  if err != nil {
    log.Fatal(err)
  }

  jobj, _ := json.Marshal(list)
  fmt.Fprintln(w, string(jobj))
}
