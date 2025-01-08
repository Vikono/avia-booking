package db

import (
	"database/sql"
	"homework/specs"
	"log"

	"github.com/google/uuid"
)

func GetFlights(db *sql.DB, params specs.GetFlightsParams) ([]specs.FlightForList, error) {
	query := `
		SELECT 
			f.id, 
			CONCAT(c.name, ' ', a.airport_code) as origin,
			CONCAT(c2.name, ' ', a2.airport_code) as destination,
			f.departure_date, f.departure_time,
			f.arrival_date, f.arrival_time,
			co.name as company_name 
			FROM flights as f
			JOIN airports as a ON f.origin_id = a.id
			JOIN airports as a2 ON f.destination_id = a2.id
			JOIN cities as c ON a.city_id = c.id 
			JOIN cities as c2 ON a2.city_id = c2.id 
			JOIN companies as co ON f.company_id = co.id 
			WHERE c.name = $1 AND c2.name = $2 AND f.departure_date = $3`

	results, err := db.Query(query, params.Origin, params.Destination, params.DepartureDate.String())
	if err != nil {
		return nil, err

	}
	var flightList []specs.FlightForList
	for results.Next() {
		var flight specs.FlightForList
		// sequence: id origin destination departuredate departuretime  arrivaldate arrivaltime company
		err = results.Scan(&flight.Id, &flight.Origin, &flight.Destination, &flight.DepartureDate, &flight.DepartureTime, &flight.ArrivalDate, &flight.ArrivalTime, &flight.Company)
		if err != nil {
			return nil, err
		}
		minMaxCost, err := GetFlightPrices(db, flight.Id)
		if err != nil {
			return nil, err
		}
		flight.ArrivalDate = flight.ArrivalDate[:10] // ээээээээээээээээээээээээм
		flight.ArrivalTime = flight.ArrivalTime[11:19]
		flight.DepartureDate = flight.DepartureDate[:10]
		flight.DepartureTime = flight.DepartureTime[11:19]
		flight.MaxCost = minMaxCost.maxPrice
		flight.MinCost = minMaxCost.minPrice
		flightList = append(flightList, flight)

	}
	return flightList, nil
}

type MinMaxPrice struct {
	minPrice int64
	maxPrice int64
}

func GetFlightPrices(db *sql.DB, flight_id uuid.UUID) (MinMaxPrice, error) {
	query := `SELECT MIN(cost) from class_data where flight_id = $1`
	results, err := db.Query(query, flight_id)
	if err != nil {
		return MinMaxPrice{}, err
	}
	log.Println(results)
	var minPrice int64
	for results.Next() {
		var price int64
		err = results.Scan(&price)
		log.Println(price)
		if err != nil {
			return MinMaxPrice{}, err
		}
		minPrice = price
	}

	query = `SELECT MAX(cost) from class_data where flight_id = $1`
	results, err = db.Query(query, flight_id)
	if err != nil {
		return MinMaxPrice{}, err
	}
	log.Println(results)
	var maxPrice int64
	for results.Next() {
		var price int64
		err = results.Scan(&price)
		if err != nil {
			return MinMaxPrice{}, err
		}
		log.Println(price)
		maxPrice = price
	}
	log.Println(minPrice, maxPrice)

	return MinMaxPrice{minPrice, maxPrice}, nil
}

func GetSimpleFlight(db *sql.DB, flight_id uuid.UUID) specs.SimpleResponse {
	var ans specs.SimpleResponse
	ans.Id = flight_id

	return ans
}
