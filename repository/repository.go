package repository

import (
	"net/http"

	"github.com/jasonkofo/servicestarter/servicestarter"
	"github.com/julienschmidt/httprouter"
)

// Ping HTTP handler
func Ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	servicestarter.SendPong(w)
}
