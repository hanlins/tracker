package main

import (
	"github.com/gorilla/mux"
	//"errors"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// identifier
var ida_max = 0
var script_file = []byte("")

/**
 * @brief Initialize tracker A.
 */
func InitTrackerA() {
	script_file, _ = ioutil.ReadFile("../trackera/script.js")
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
	w.Write(script_file)
	http.ServeFile(w, r, "../trackera/"+script+".js")
	w.Write([]byte("type_a(\"" + cookie_name + "\", " + id + ");"))
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
	RecordRefer(id, referer)
}
