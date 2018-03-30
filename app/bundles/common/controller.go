package common

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/domgoodwin/go-api/app/bundles/db"
)

// Controller handle all base methods
type Controller struct {
}

// SendJSON marshals v to a json struct and sends appropriate headers to w
func (c *Controller) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(v)
	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

// Index func return all devices in database
func (c *Controller) ListTables(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		db.ListTables(),
		http.StatusOK,
	)
}
