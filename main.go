package main

import (
	"log"
	"net/http"

	"github.com/domgoodwin/go-api/app/bundles/devicebundle"

	"github.com/gorilla/mux"

	"github.com/domgoodwin/go-api/app/bundles/db"
)

func main() {
	d := &db.Db{}
	d.List()
	// Controllers declaration
	kc := &devicebundle.DeviceController{}
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()
	// Routes handling
	s.HandleFunc("/device", kc.Index).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
