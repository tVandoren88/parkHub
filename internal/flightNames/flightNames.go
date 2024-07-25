package flightNames

import (
	"log"

	database "github.com/tvandoren88/parkHub/internal/pkg/db/migrations/mysql"
)

type FlightName struct {
	ID       string
	FlightID string
	Name     string
}

// Input new Flight name
func (flightName FlightName) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO FlightNames(FlightID, Name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(flightName.FlightID, flightName.Name)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

// Get all flight names
func GetAll() []FlightName {
	stmt, err := database.Db.Prepare("select id, flightID, name from FlightNames")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var flightNames []FlightName
	for rows.Next() {
		var flightName FlightName
		err := rows.Scan(&flightName.ID, &flightName.FlightID, &flightName.Name)
		if err != nil {
			log.Fatal(err)
		}
		flightNames = append(flightNames, flightName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return flightNames
}

// Get flight name based on flight ID
func GetByFlightID(flightId int) *FlightName {
	stmt, err := database.Db.Prepare("select id, flightID, name from FlightNames Where flightID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(flightId)
	var flightName FlightName
	err = row.Scan(&flightName.ID, &flightName.FlightID, &flightName.Name)
	if err != nil {

		return nil
	}

	return &flightName
}

// Get flight name based on flight name
func GetByFlightName(name string) *FlightName {
	stmt, err := database.Db.Prepare("select id, flightID, name from FlightNames Where name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(name)
	var flightName FlightName
	err = row.Scan(&flightName.ID, &flightName.FlightID, &flightName.Name)
	if err != nil {

		return nil
	}

	return &flightName
}
