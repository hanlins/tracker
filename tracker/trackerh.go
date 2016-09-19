/**
 * @file trackerh.go
 * @brief A simple implementation of type - H tracker. Using tab state.
 * @author Hanlin Shi
 * @version 0.1.0
 */
package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

/**
 * @brief Handling function for type - A tracker.
 *
 *  This tracker will serve a script to the first - party page in which it will
 *  set an id in cookie of that first - party page.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerHHandler(w http.ResponseWriter, r *http.Request) {
	file := mux.Vars(r)["script"]
	ServeFileWrapper(w, r, "../trackerh/"+file+".js")
}

/**
 * @brief Handling function for listener of type - A tracker.
 *
 *  This function will parest the request, then record the info transtered.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func ListenerHHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	m, _ := url.ParseQuery(u.RawQuery)
	referer := m["referer"][0]
	footprint := m["footprint"][0]
	RecordRefer("tkh", footprint, referer)
	// MapFootprintId(footprint, id)
}
