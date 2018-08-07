package devicebundle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/domgoodwin/go-api/app/bundles/common"
	"github.com/domgoodwin/go-api/app/bundles/db"
	"github.com/domgoodwin/go-api/app/bundles/r53"
)

// DeviceController struct
type DeviceController struct {
	common.Controller
}

// Index func return all devices in database
func (c *DeviceController) Index(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		[]*Device{NewDevice("host", "192.168.0.1", "80.80.80.80")},
		http.StatusOK,
	)
}

// Create func return all devices in database
func (c *DeviceController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(json.NewDecoder(r.Body))
	res := db.PutItem("home-devices", NewDevice("hello", "world", "indeed"))
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}

// Get record sets
func (c *DeviceController) GetRecords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := r53.GetRecordSets(vars["id"])
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}

// Get record sets
func (c *DeviceController) UpdateRecordSet(w http.ResponseWriter, r *http.Request) {
	var rs r53.RecordSet
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	println(rs.HostedZoneId)
	res := r53.UpdateRecordSet(rs)
	c.SendJSON(
		w,
		r,
		res,
		http.StatusOK,
	)
}
