package devicebundle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/domgoodwin/go-api/app/bundles/common"
	"github.com/domgoodwin/go-api/app/bundles/db"
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
