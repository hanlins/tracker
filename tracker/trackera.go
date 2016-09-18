/**
 * @file trackera.go
 * @brief A simple implementation of type - A tracker.
 * @author Hanlin Shi
 * @version 0.1.0
 */
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
 * Global variables.
 */
var ida_max = 0
var script_file_A = []byte("")

/**
 * @brief Initialize tracker A.
 * @return Void.
 */
func InitTrackerA() {
	script_file_A, _ = ioutil.ReadFile("../trackera/script.js")
}

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
func TrackerAHandler(w http.ResponseWriter, r *http.Request) {
	// check whether cookie has been set
	cookie_name := "IDA"
	id := GenerateID(&ida_max)
	// serve the file
	script := mux.Vars(r)["script"]
	// dummy cookie, just make it httponly
	SetCookie(w, "dummy", "0")
	w.Write(script_file_A)
	// dummy cookie, just make it httponly
	SetCookie(w, "dummy", "0")
	// http.ServeFile(w, r, "../trackera/"+script+".js")
	ServeFileWrapper(w, r, "../trackera/"+script+".js")
	w.Write([]byte("type_a(\"" + cookie_name + "\", " + "\"" + id + "\"" + ");"))
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
func ListenerAHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	m, _ := url.ParseQuery(u.RawQuery)
	id := m["id"][0]
	referer := m["referer"][0]
	footprint := m["footprint"][0]
	RecordRefer("tka", id, referer)
	MapFootprintId(footprint, id)
}

/**
 * @brief Map identifier and footprint received, used for cross-site identifier
 *        referencing.
 * @param footprint Cross-site footprint used for cross-site identifier.
 * @param id Tracker-owned id for the browser user.
 * @return Void.
 */
func MapFootprintId(footprint string, id string) {
	fmt.Println(id + " leaves footprint : " + footprint)
}
