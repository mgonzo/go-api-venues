package handler

import (
	"github.com/mgonzo/go-api-config"
	"github.com/mgonzo/go-api-venues/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Collection(w http.ResponseWriter, r *http.Request) {
	table := config.Table()
	db := config.Dbconn()
	defer db.Close()

	// retrieve collection
	// need a package to filter, sort and page results
	query := fmt.Sprintf("select id, name from %[1]s", table)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id   int
		name string
		phone model.Phone
		city model.City
	)

	var list []model.Venue

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		var ven = model.Venue{
			Id:   id,
			Name: name,
			Phone: phone,
			City: city,
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
