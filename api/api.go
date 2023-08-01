package api

import (
	"fmt"
	"net/http"

	xlogger "github.com/mt1976/appFrame/logs"
)

// GetURLparam returns a selected parmeter value from the a given URI
// The function `getURLparam` retrieves a specific parameter value from a given HTTP request URL.
func getURLparam(r *http.Request, paramID string) string {
	xlogs := xlogger.New()
	//fmt.Println(paramID)
	//fmt.Println(r.URL)
	key := r.FormValue(paramID)
	//log.Printf("URL Parameter : Key=%q Value=%q", paramID, string(key))
	xlogs.Info("URL Parameter :", fmt.Sprintf("Key=%q Value=%q", paramID, string(key)))
	return key
}
