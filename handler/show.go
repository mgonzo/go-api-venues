package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/mgonzo/venues/model"
)

func Show(w http.ResponseWriter, r *http.Request) {
	// validate?
	vars := mux.Vars(r)
	qualifier := vars["qualifier"]

	table := table()
	db := dbconn()
	defer db.Close()

	var (
		id   int
		name string
	)

	query := fmt.Sprintf(
		"select id, name from %[1]s where %[2]s=%[3]s",
		table,
		viper.GetString("qualifier"),
		qualifier,
	)

	rows, err := db.Query(query)
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
