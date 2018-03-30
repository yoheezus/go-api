package devicebundle

import (
	"net/http"

	"github.com/domgoodwin/go-api/app/bundles/common"
)

// KittiesController struct
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
