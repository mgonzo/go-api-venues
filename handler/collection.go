package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mgonzo/venues/model"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Collection(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open(viper.GetString("sqltype"), connect())

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// retrieve collection
	// need a package to filter, sort and page results
	var (
		id   int
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
			Id:   id,
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
