# parkHub
ParkHub Interview

INTRO
First I would like to say I could of approached this task in a different way.  I chose to challenge myself and teach myself GO and Graphql with go.
I have been using this week to learn go and graphql before trying this challenge.  I wanted to show my ability to catch on to a new code and concepts quickly.


Steps to set up enviroment

Assumptions
    Go is already installed.


Run this to set up the db
    docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=flightInfo -d mysql:latest
    docker exc -it mysql bash
    mysql -u root -p
    CREATE DATABASE flightInfo

to see in a database ide
Port 3306
DB Name flightInfo
db password: dbpass

Run the code
go mod tidy
go install
go run .\server.go


If you have issues with migrations I did have to do these steps locally to get it installed.
go to this path
 %GOPATH%\pkg\mod\github.com\golang-migrate\migrate\v4@v4.17.1\cmd\migrate
 run go install in cmd line

Flights with multiple names are
FlightID 1,5,10

Things I would like to add is more error checking with arrival and departure times
I would add created at dates in the flight names and use that for a chronological history

Interacting with the API

CREATE A FLIGHT
    mutation createFlight{
        createFlight(input: { name: "test", arrivalTime: "2012-12-12", departureTime: "2012-12-11"}){
            id
            name
            arrivalTime
            departureTime
        }
    }

FIND ALL FLIGHTS

query findFligths {
  flights{
    id
    name
    departureTime
    arrivalTime
  }
}

UPDATE FLIGHT

mutation updateFlight{
  updateFlight(id: "5", edits: {name: "new3"}){
    id
    name
    arrivalTime
    departureTime
  }

}

INSERT MULTIPLE FLIGHTS
mutation {
  insertFlights(
    objects: [
      {name: "flight10", departureTime: "2012-12-20", arrivalTime: "2012-12-21"},

      {name: "flight20", departureTime: "2012-12-20", arrivalTime: "2012-12-21"},
    ]
  ){
    id
    name
    arrivalTime
    departureTime
  }
}

FIND ALL NAMES USED
query findNames {
  flightNames {
    id
    flightId
    name
  }
}

FIND FLIGHT BY NAME
mutation FlightDetails{
  getFlightDetails(findName: "test"){
    id
    name
    arrivalTime
    departureTime
  }
}
