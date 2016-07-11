/**
 * @file trackera.go
 * @brief A simple implementation of type - D tracker.
 * @author Hanlin Shi
 * @version 0.1.0
 */
package main

import (
	"net/http"
	"net/url"
)

/**
 * Global variables.
 */
var idd_max = 0

/**
 * @brief Handling function for listener of type - A tracker.
 *
 *  This function will parest the request, then record the info transtered.
 *
 * @param w HTTP response.
 * @param r HTTP request.
 * @return Void.
 */
func ListenerDHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	m, _ := url.ParseQuery(u.RawQuery)
	id := m["id"][0]
	referer := m["referer"][0]
	RecordRefer(id, referer)
}
