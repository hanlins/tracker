package main

import (
	//"github.com/gorilla/mux"
	//"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

/**
 * @brief Check whether corresponding cookie exists in request.
 * @param r HTTP request.
 * @param name Cookie name that we want to check.
 * @return true if exists, false if not.
 */
func CookieExists(r *http.Request, name string) bool {
	_, err := r.Cookie(name)
	if err == nil {
		return true
	} else {
		return false
	}
}

/**
 * @brief Get cookie value.
 * @param r HTTP request.
 * @param name Cookie name that we want to get.
 * @return Cookie's value.
 */
func GetCookie(r *http.Request, name string) string {
	val, _ := r.Cookie(name)
	return val.Value
}

/**
 * @brief Set cookie.
 * @param w HTTP response.
 * @param name Cookie name that we want to set.
 * @param val Cookie value that we want to set.
 * @return Void.
 */
func SetCookie(w http.ResponseWriter, name string, val string) {
	cookie := &http.Cookie{
		Name:  name,
		Value: val,
		//Path:    "/",
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, cookie)
	return
}

/**
 * @brief Generate ID.
 * @param id_max Recorder for maximum id, used for new id generation.
 * @return ID string.
 */
func GenerateID(id_max *int) string {
	identifier := *id_max
	*id_max = *id_max + 1
	return strconv.Itoa(identifier)
}

/**
 * @brief Record user visiting event.
 *
 *  TODO: use database or logs to store the visit records.
 *
 * @param id The identifier of the user.
 * @param url The referer url that user just visited.
 * @return Whether there's error.
 */
func RecordRefer(id string, url string) error {
	fmt.Println(id + " : " + url)
	return nil
}
