package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
)

type Reservation struct {
	Category string
	Date     string
	Period   string
}

func main() {
	r := mux.NewRouter()

	fmt.Printf("Starting server at port 8081\n")
	r.HandleFunc("/reservation", ReservationHandler).Methods("POST")

	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatal(err)
	}
}

func ReservationHandler(w http.ResponseWriter, r *http.Request) {
	data := new(Reservation)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(data.Category, data.Date, data.Period)

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("data/data.json", file, 0644)
}
