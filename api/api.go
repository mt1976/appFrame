package api

import (
	"fmt"
	"net/http"

	xlogs "github.com/mt1976/appFrame/logs"
)

// GetURLparam returns a selected parmeter value from the a given URI
func getURLparam(r *http.Request, paramID string) string {
	//fmt.Println(paramID)
	//fmt.Println(r.URL)
	key := r.FormValue(paramID)
	//log.Printf("URL Parameter : Key=%q Value=%q", paramID, string(key))
	xlogs.Information("URL Parameter :", fmt.Sprintf("Key=%q Value=%q", paramID, string(key)))
	return key
}
