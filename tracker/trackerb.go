/**
 * @file trackerb.go
 * @brief A simple implementation of type - B tracker.
 * @author Hanlin Shi
 * @version 0.1.0
 */
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * Global variables.
 */
var idb_max = 0

/**
 * @brief Handling function for type - B tracker.
 *
 *  This tracker will serve a page to the first - party page in which it will
 *  set an id in cookie of that first - party page.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerBHandler(w http.ResponseWriter, r *http.Request) {
	TrackerBCookieAndRecord(w, r, "IDB", &idb_max, "tkb")
	TrackerBCookieAndRecord(w, r, "IDD", &idd_max, "tkd")
	// serve the file
	file := mux.Vars(r)["file"]
	http.ServeFile(w, r, "../trackerb/"+file)
}

/**
 * @brief Handling function for type - B tracker.
 *
 *  This tracker will set cookie and record the referer.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @param cookie_name Name of the cookie to be test and set.
 * @param id_pt Pointer to id counter that used to assign unique id.
 * @param tracker Name of the tracker to be recorded.
 * @return Void.
 */
func TrackerBCookieAndRecord(w http.ResponseWriter, r *http.Request,
	cookie_name string, id_pt *int, tracker string) {
	// check whether cookie has been set
	id := ""
	if CookieExists(r, cookie_name) {
		id = GetCookie(r, cookie_name)
	} else {
		// set cookie
		id = GenerateID(id_pt)
		SetCookie(w, cookie_name, id)
	}
	// get referrer header
	referer_url := r.Header.Get("Referer")
	// record event
	if referer_url != "" {
		err := RecordRefer(tracker, id, referer_url)
		if err != nil {
			fmt.Println("referer url record error")
			return
		}
	}
}
