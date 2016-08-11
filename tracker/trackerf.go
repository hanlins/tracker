/**
 * @file trackerf.go
 * @brief A simple implementation of type - F tracker, track using flash cookie
 * @author Hanlin Shi
 * @version 0.0.1
 */
package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

/**
 * Global variables.
 */
var idf_max = 0
var script_file_F = []byte("")

/**
 * @brief Initialize tracker A.
 * @return Void.
 */
func InitTrackerF() {
	script_file_F, _ = ioutil.ReadFile("../trackerf/script.js")
}

/**
 * @brief Handler for tracker type - F that records browser user visits.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerFHandler(w http.ResponseWriter, r *http.Request) {
	file := mux.Vars(r)["file"]
	// if the script file that set tracker id, feed proposal id
	if file == "script.js" {
		// generate id
		id := GenerateID(&idf_max)
		// serve the script file
		w.Write(script_file_F)
		// @TODO This is questionable, it also happens in type - A tracker, check
		//       and fix it later.
		http.ServeFile(w, r, "../trackerf/script.js")
		// execute functions
		w.Write([]byte("type_f(\"" + "IDF" + "\", " + "\"" + id + "\"" + ");"))
	} else {
		// otherwise, serve the flieV first
		http.ServeFile(w, r, "../trackerf/"+file)
	}
}
