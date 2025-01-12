package v1

import (
	"encoding/json"
	"homework/internal/db"
	"homework/specs"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (a apiServer) GetFlights(w http.ResponseWriter, r *http.Request, params specs.GetFlightsParams) {
	flight_list, err := db.GetFlights(a.DB, params)
	if err != nil {
		log.Printf("Error: %s while fetching the query result", err)
		w.WriteHeader(500)
		return
	}
	log.Println(flight_list)
	// response, err := json.Marshal(flight_list)
	// if err != nil {
	// 	log.Printf("Error while marshalling the result")
	// 	w.WriteHeader(500)
	// 	return
	// }
	// log.Println(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flight_list)

}

func (a apiServer) GetFlight(w http.ResponseWriter, r *http.Request, fligthId uuid.UUID) {
	flightCard, err := db.GetFlightCard(a.DB, fligthId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flightCard)

}

func (a apiServer) GetSimpleFlight(w http.ResponseWriter, r *http.Request, flightId uuid.UUID) {
	// flight_id := db.GetSimpleFlight(a.DB, flightId)

	// response, err := json.Marshal(1)
	// if err != nil {
	// 	w.WriteHeader(500)
	// 	return
	// }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(1)
}
