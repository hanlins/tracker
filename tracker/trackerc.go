/**
 * @file trackera.go
 * @brief A simple implementation of type - C tracker.
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
var idc_max = 0

/**
 * @brief Handling function for type - C tracker.
 *
 *  This tracker is based on type - B checker, in which it will FORCE third-
 *  party become first-party (redirect, or popup) to bypass blocking for third-
 *  party cookies. It will simply serve a page. In the page, it will
 *  1). Check whether the id is set.
 *  2). If not set, transfer control to new page with first-class identity and
 *      set id in cookie.
 *  3). Track the browser user using tracker-owned cookie set in first-party.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerCHandler(w http.ResponseWriter, r *http.Request) {
	// check whether cookie has been set
	id := ""
	cookie_name := "IDC"
	if CookieExists(r, cookie_name) {
		id = GetCookie(r, cookie_name)
	} else {
		// set cookie
		id = GenerateID(&idb_max)
		SetCookie(w, cookie_name, id)
	}
	// get referrer header
	referer_url := r.Header.Get("Referer")
	// record event
	if referer_url != "" {
		err := RecordRefer(id, referer_url)
		if err != nil {
			fmt.Println("referer url record error")
			return
		}
	}
	// serve the file
	file := mux.Vars(r)["file"]
	http.ServeFile(w, r, "../trackerc/"+file)
}

/**
 * @brief Transfer control to first-party page of tracker that set cookie using
 *        popup page.
 * @param w HTTP response.
 * @param cookie_name Cookie name.
 * @return Void.
 */
func TransferSetCookie(w http.ResponseWriter, cookie_name string) {
	id := GenerateID(&idc_max)
	SetCookie(w, cookie_name, id)
	return
}

/**
 * @brief Set Cookie in the new popup for type - C tracker.
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerCSetCookie(w http.ResponseWriter, r *http.Request) {
	cookie_name := "IDC"
	TransferSetCookie(w, cookie_name)
	return
}
