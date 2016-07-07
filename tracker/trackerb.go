package main

import (
	//"github.com/gorilla/mux"
	//"errors"
	"fmt"
	"net/http"
)

// identifier
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
	// check whether cookie has been set
	id := ""
	cookie_name := "IDB"
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
	//script := mux.Vars(r)["script"]
	//http.ServeFile(w, r, "../trackera/"+script+".js")
}
