package main

import (
	"log"
	"net/http"

	"github.com/domgoodwin/go-api/app/bundles/devicebundle"

	"github.com/gorilla/mux"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Controllers declaration
	kc := &devicebundle.DeviceController{}
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()
	// Routes handling
	s.HandleFunc("/prime", kc.NextPrime).Methods("GET")
	s.HandleFunc("/device", kc.Index).Methods("GET")
	s.HandleFunc("/device", kc.Create).Methods("POST")
	s.HandleFunc("/db/tables", kc.ListTables).Methods("GET")
	s.HandleFunc("/r/records/{id}", kc.GetRecords).Methods("GET")
	s.HandleFunc("/r/records", kc.UpdateRecordSet).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
