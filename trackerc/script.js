/**
 * @file script.js
 * @description This file contains the script to be sent as response for type C
 *              tracker.
 * @author Hanlin Shi (hanlins)
 * @version 0.1.0
 */

/**
 * @description This function returns the cookie value with its name.
 * @param name name of the cookie.
 * @param d document (from child).
 * @return cookie value.
 */
function getCookie(name, cookie) {
  var value = "; " + cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}

/**
 * @description This function set cookie in the first - party page, it will
 *              1). examine whether the cookie has already been set.
 *              2). set a new id in first-party popup if has not been set, then
 *                  reload the page.
 * @param id_name The identifier's name stored in the cookie.
 * @param set_cookie_url The url that open in popup and set cookie for tracker.
 * @return Void.
 */
function type_c(id_name, set_cookie_url) {
  /** examine whether cookie value exists. */
  var id_value = getCookie(id_name, window.document.cookie);
  /** set the cookie value in popup if it has not been set. */
  if (id_value === undefined) {
    setCookiePopup(set_cookie_url);
    /** reload the page after 5 second */
    deferReloadPage(5);
  }
}

/**
 * @description In this function, a pop-up page will be opened and a cookie
 *              will be set in that page.
 * @param set_cookie_url URL of the pop-up that directs the browser to set cookie.
 * @return Void.
 */
function setCookiePopup(set_cookie_url) {
  (window.open(set_cookie_url, '_blank')).focus();
}

/**
 * @description Reload the page / iframe after 'delay' seconds.
 * @param delay Waiting time for reloading the page in unit of seconds.
 * @return Void.
 */
function deferReloadPage(delay) {
  setTimeout(function(){location.reload();}, 1000 * delay);
}

/** invoke type - C tracker's script. */
type_c("IDC", "http://trackerc.com/c/setcookiec");
