/**
 * @file trackerg.go
 * @brief A simple implementation of type - G tracker.
 * @author Hanlin Shi
 * @version 0.1.0
 */
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * Global variables.
 */
var idg_max = 0

func TrackerGHandler(w http.ResponseWriter, r *http.Request) {
	cookie_name := "IDG"
	// referer_url := r.Header.Get("Referer")
	if CookieExists(r, cookie_name) {
		// record event
		id := GetCookie(r, cookie_name)
		RecordRefer("tkg", id, "visited trackerg before")
	} else {
		TransferSetCookie(w, cookie_name)
		id := GenerateID(&idg_max)
		SetCookie(w, cookie_name, id)
	}
	file := mux.Vars(r)["file"]
	ServeFileWrapper(w, r, "../trackerg/"+file)
}
