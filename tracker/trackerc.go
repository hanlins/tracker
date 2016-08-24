/**
 * @file trackera.go
 * @brief A simple implementation of type - C tracker.
 * @author Hanlin Shi
 * @version 0.1.1
 */
package main

import (
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
 *  2). If not set, server a file to set the cookie.
 *  3). If set, serve the correct advertisement content.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerCHtmlHandler(w http.ResponseWriter, r *http.Request) {
	cookie_name := "IDC"
	referer_url := r.Header.Get("Referer")
	if CookieExists(r, cookie_name) {
		// record event
		id := GetCookie(r, cookie_name)
		RecordRefer("tkc", id, referer_url)
	}
	file := mux.Vars(r)["file"]
	http.ServeFile(w, r, "../trackerc/"+file+".html")
}

/**
 * @brief Serve script files for type - C tracker.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func TrackerCScriptHandler(w http.ResponseWriter, r *http.Request) {
	script := mux.Vars(r)["script"]
	http.ServeFile(w, r, "../trackerc/"+script+".js")
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
	http.ServeFile(w, r, "../trackerc/close.html")
	return
}
