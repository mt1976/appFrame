package AppFrame

import (
	"fmt"
	"net/http"

	logs "github.com/mt1976/AppFrame/logs"
)

// GetURLparam returns a selected parmeter value from the a given URI
func GetURLparam(r *http.Request, paramID string) string {
	//fmt.Println(paramID)
	//fmt.Println(r.URL)
	key := r.FormValue(paramID)
	//log.Printf("URL Parameter : Key=%q Value=%q", paramID, string(key))
	logs.Information("URL Parameter :", fmt.Sprintf("Key=%q Value=%q", paramID, string(key)))
	return key
}
