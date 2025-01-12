package db

import (
	"database/sql"
	"homework/specs"
	"log"
	"time"

	"github.com/google/uuid"
)

func endOfDate(date string) (time.Time, error) {
	log.Println(date)
	datetime, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", date)
	return datetime.Add(23*time.Hour + 59*time.Minute + 59*time.Second), err
}

func beginingOfDate(date string) (time.Time, error) {
	new_datetime, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", date)
	return new_datetime, err
}

func DBGetFlightsForList(db *sql.DB, origin string, destination string, departure_datetime string) (*sql.Rows, error) {
	query := `
		SELECT 
			f.id, 
			CONCAT(c.name, ' ', a.airport_code) as origin,
			CONCAT(c2.name, ' ', a2.airport_code) as destination,
			f.departure_datetime,
			f.arrival_datetime,
			co.name as company_name 
			FROM flights as f
			JOIN airports as a ON f.origin_id = a.id
			JOIN airports as a2 ON f.destination_id = a2.id
			JOIN cities as c ON a.city_id = c.id 
			JOIN cities as c2 ON a2.city_id = c2.id 
			JOIN companies as co ON f.company_id = co.id 
			WHERE c.name = $1 AND c2.name = $2 AND f.departure_datetime BETWEEN $3 AND $4`

	endTime, err := endOfDate(departure_datetime)
	if err != nil {
		return nil, err
	}
	startTime, err := beginingOfDate(departure_datetime)
	if err != nil {
		return nil, err
	}
	results, err := db.Query(query, origin, destination, startTime, endTime)
	return results, err
}

func DBGetFlightById(db *sql.DB, id uuid.UUID) (*sql.Rows, error) {
	query := `
		SELECT 
			f.id, 
			CONCAT(c.name, ' ', a.airport_code) as origin,
			CONCAT(c2.name, ' ', a2.airport_code) as destination,
			f.departure_datetime,
			f.arrival_datetime,
			co.name as company_name 
			FROM flights as f
			JOIN airports as a ON f.origin_id = a.id
			JOIN airports as a2 ON f.destination_id = a2.id
			JOIN cities as c ON a.city_id = c.id 
			JOIN cities as c2 ON a2.city_id = c2.id 
			JOIN companies as co ON f.company_id = co.id 
			WHERE f.id = $1`

	results, err := db.Query(query, id)
	return results, err
}

func DBGetFlightClassData(db *sql.DB, id uuid.UUID, class_type *string) (*sql.Rows, error) {
	query := `SELECT c.class_type, c.buggage, c.luggage, c.cost FROM class_data as c
			WHERE c.flight_id = $1`

	if class_type != nil {
		query += ` AND c.class_type = $3`
		results, err := db.Query(query, id, &class_type)
		return results, err
	}

	results, err := db.Query(query, id)
	return results, err

}

func GetMinFlightPrice(db *sql.DB, flight_id uuid.UUID) (int64, error) {
	query := `SELECT MIN(cost) from class_data where flight_id = $1`
	results, err := db.Query(query, flight_id)
	if err != nil {
		return -1, err
	}

	var minPrice int64
	for results.Next() {
		var price int64
		err = results.Scan(&price)
		log.Println(price)
		if err != nil {
			return -1, err
		}
		minPrice = price
	}
	return minPrice, nil
}

func GetMaxFlightPrice(db *sql.DB, flight_id uuid.UUID) (int64, error) {

	query := `SELECT MAX(cost) from class_data where flight_id = $1`
	results, err := db.Query(query, flight_id)
	if err != nil {
		return -1, err
	}

	var maxPrice int64
	for results.Next() {
		var price int64
		err = results.Scan(&price)
		if err != nil {
			return -1, err
		}
		log.Println(price)
		maxPrice = price
	}

	return maxPrice, nil
}

func GetFlights(db *sql.DB, params specs.GetFlightsParams) ([]specs.FlightForList, error) {
	results, err := DBGetFlightsForList(db, params.Origin, params.Destination, params.DepartureDate.String())
	if err != nil {
		return nil, err
	}

	var flightList []specs.FlightForList
	for results.Next() {
		var flight specs.FlightForList
		// sequence: id origin destination departuredate departuretime  arrivaldate arrivaltime company
		err = results.Scan(&flight.Id, &flight.Origin, &flight.Destination, &flight.DepartureDatetime, &flight.ArrivalDatetime, &flight.Company)
		if err != nil {
			return nil, err
		}

		flight.MaxCost, err = GetMinFlightPrice(db, flight.Id)
		if err != nil {
			return nil, err
		}
		flight.MinCost, err = GetMaxFlightPrice(db, flight.Id)
		if err != nil {
			return nil, err
		}
		flightList = append(flightList, flight)

	}
	return flightList, nil
}

// type MinMaxPrice struct {
// 	minPrice int64
// 	maxPrice int64
// }

func GetFlightCard(db *sql.DB, flightId uuid.UUID) (specs.FlightCard, error) {
	result, err := DBGetFlightById(db, flightId)
	if err != nil {
		return specs.FlightCard{}, err
	}

	var flightCard specs.FlightCard
	for result.Next() {
		// sequence: id origin destination departuredate departuretime  arrivaldate arrivaltime company
		err = result.Scan(
			&flightCard.Id, &flightCard.Origin, &flightCard.Destination,
			&flightCard.DepartureDatetime, &flightCard.ArrivalDatetime, &flightCard.Company)
		if err != nil {
			return specs.FlightCard{}, err
		}

		class_data, err := DBGetFlightClassData(db, flightId, nil)
		if err != nil {
			return specs.FlightCard{}, err
		}

		for class_data.Next() {
			var classCard specs.ClassData
			err = class_data.Scan(&classCard.Class, &classCard.Buggage, &classCard.Luggage, &classCard.Cost)
			if err != nil {
				return specs.FlightCard{}, err
			}

			flightCard.ClassData = append(flightCard.ClassData, classCard)
		}
	}

	return flightCard, nil

}
