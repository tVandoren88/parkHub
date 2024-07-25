CREATE TABLE IF NOT EXISTS Flights(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Name VARCHAR (127) NOT NULL UNIQUE,
    DepartureTime timestamp NOT NULL,
    ArrivalTime timestamp NOT NULL,
    PRIMARY KEY (ID)
)
