package main

import (
	"log"
	"net/http"

	"github.com/domgoodwin/go-api/app/bundles/devicebundle"

	"github.com/gorilla/mux"
)

func main() {
	// Controllers declaration
	kc := &devicebundle.DeviceController{}
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()
	// Routes handling
	s.HandleFunc("/device", kc.Index).Methods("GET")
	s.HandleFunc("/device", kc.Create).Methods("POST")
	s.HandleFunc("/db/tables", kc.ListTables).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
