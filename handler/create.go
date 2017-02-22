package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mgonzo/venues/model"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// get post args
	// validation?
	decoder := json.NewDecoder(r.Body)

	var v model.Venue
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	db, err := sql.Open(viper.GetString("sqltype"), connect())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	now := time.Now()

	// get constraints
	// Build the insert string
	table := fmt.Sprintf("insert into %[1]s", "venue")

	cols := fmt.Sprintf(
		"(%[1]s, %[2]s, %[3]s, %[4]s)",
		"created_at",
		"name",
		"timezone_id",
		"city_id",
	)

	vals := fmt.Sprintf(
		"values (%[1]d, '%[2]s', %[3]d, %[4]d)",
		now.Unix(),
		v.Name,
		timezone(db, v.Timezone),
		city(db, v.Zip),
	)

	insert := fmt.Sprintf("%[1]s %[2]s %[3]s", table, cols, vals)
	log.Println(insert)

	// insert a new record
	res, err := db.Exec(insert)
	if err != nil {
		log.Fatal(err)
	}

	// get the record that was just created
	venue_id, err := res.LastInsertId()
	log.Println(venue_id)

	// create the phone
	phone_id := phone(db, venue_id, v.Phone)

	// return 201 with the composite
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, venue_id)
	fmt.Fprintln(w, phone_id)
}

func timezone(db *sql.DB, location string) int {
	tzq := fmt.Sprintf("select id from timezone where location='%[1]s'", location)
	tz, err := db.Query(tzq)
	if err != nil {
		log.Fatal(err)
	}
	defer tz.Close()

	var id int
	for tz.Next() {
		err := tz.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return id
}

func city(db *sql.DB, zip int) int {
	cityq := fmt.Sprintf("select id from city where zip='%[1]d'", zip)
	city, err := db.Query(cityq)
	if err != nil {
		log.Fatal(err)
	}
	defer city.Close()

	var id int
	for city.Next() {
		err := city.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return id
}

func phone(db *sql.DB, venue_id int64, p model.Phone) int64 {
	table := fmt.Sprintf("insert into %[1]s", "phone")

	cols := fmt.Sprintf(
		"(%[1]s, %[2]s, %[3]s, %[4]s)",
		"venue_id",
		"country_code",
		"area_code",
		"number",
	)

	vals := fmt.Sprintf(
		"values (%[1]d, '%[2]s' ,'%[3]s','%[4]s')",
		venue_id,
		"1",
		p.Area,
		p.Number,
	)

	insert := fmt.Sprintf("%[1]s %[2]s %[3]s", table, cols, vals)
	log.Println(insert)

	// insert a new record
	res, err := db.Exec(insert)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	return id
}
