package api

import (
	"net/http"
)

// GetURLparam returns a selected parmeter value from the a given URI
// The function GetURLparam retrieves a specific parameter from a given HTTP request.
func GetURLparam(r *http.Request, paramID string) string {
	return getURLparam(r, paramID)
}
