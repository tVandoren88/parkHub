package flights

import (
	"log"

	database "github.com/tvandoren88/parkHub/internal/pkg/db/migrations/mysql"
)

type Flight struct {
	ID            string
	Name          string
	DepartureTime string
	ArrivalTime   string
}

// Insert new data
func (flight Flight) Save() int64 {
	//set up insert query
	stmt, err := database.Db.Prepare("INSERT INTO Flights(Name, departureTime, arrivalTime) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	// Excute query with params from the flight
	res, err := stmt.Exec(flight.Name, flight.DepartureTime, flight.ArrivalTime)
	if err != nil {
		log.Fatal(err)
	}
	//Find newest Id inserted  to return success
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

// Get all flights availiable
func GetAll() []Flight {
	// Set up query
	stmt, err := database.Db.Prepare("select id, name, departureTime, arrivalTime from Flights")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//excute query and return data set or error to pass back
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var flights []Flight
	// Go through data set and send back only needed data
	for rows.Next() {
		var flight Flight
		err := rows.Scan(&flight.ID, &flight.Name, &flight.DepartureTime, &flight.ArrivalTime)
		if err != nil {
			log.Fatal(err)
		}
		flights = append(flights, flight)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return flights
}

// Get flight based on ID
func GetByID(id int) *Flight {
	stmt, err := database.Db.Prepare("select id, name, departureTime, arrivalTime from Flights Where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	var flight Flight
	err = row.Scan(&flight.ID, &flight.Name, &flight.DepartureTime, &flight.ArrivalTime)
	if err != nil {

		return nil
	}

	return &flight
}

// Update Flight details
func (flight Flight) Update() int64 {
	stmt, err := database.Db.Prepare("Update Flights SET Name = ?, departureTime = ?, arrivalTime = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(flight.Name, flight.DepartureTime, flight.ArrivalTime, flight.ID)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row Updated!")
	return id
}
