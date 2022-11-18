package api

import (
	"net/http"
)

// GetURLparam returns a selected parmeter value from the a given URI
func GetURLparam(r *http.Request, paramID string) string {
	return getURLparam(r, paramID)
}
