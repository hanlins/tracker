/**
 * @file trackere.go
 * @brief A simple implementation of type - E tracker.
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
var ide_max = 0

/**
 * @brief Handler for type - E tracker that serve homepage (first-party).
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerEHomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie_name := "IDE"
	if !CookieExists(r, cookie_name) {
		SetCookie(w, cookie_name, GenerateID(&ide_max))
	}
	// for tracker D
	cookie_name_master := "IDD"
	if !CookieExists(r, cookie_name_master) {
		SetCookie(w, cookie_name_master, GenerateID(&idd_max))
	}
	http.ServeFile(w, r, "../trackere/home.html")
}

/**
 * @brief Handler for tracker type - E that records browser user visits.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerEHandler(w http.ResponseWriter, r *http.Request) {
	cookie_name := "IDE"
	if CookieExists(r, cookie_name) {
		id := GetCookie(r, cookie_name)
		referer_url := r.Header.Get("Referer")
		RecordRefer(id, referer_url)
	}
	file := mux.Vars(r)["file"]
	http.ServeFile(w, r, "../trackere/"+file)
}
