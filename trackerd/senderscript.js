/**
 * @file senderscript.js
 * @description This file contains the script to be sent request to type - D
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
function getCookie(name, doc) {
  var value = "; " + doc.cookie;
  value = value.replace(/\s/g, '');// remove space
  var parts = value.split(";" + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}

/**
 * @description This function set cookie in the first - party page, it will
 *              1). examine whether the cookie has already been set.
 *              2). transfet the identifier together with other information to
 *                  type - D tracker if cookie has been set.
 * @param id_name The identifier's name stored in the cookie.
 */
function type_d(id_name) {
  /** examine whether cookie value exists. */
  var id_value = getCookie(id_name, document);
  /** set the cookie value if it has not been set. */
  if (id_value === undefined) {
    return;
  }
  /** transmit the id together with other information back to tracker. */
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (xhttp.readyState == 4 && xhttp.status == 200) {
      console.log("tracker D: request sent, with response: " + xhttp.responseText);
    }
  };
  var tracker_receiver = "http://trackerd.com/listenerd";
  xhttp.open("GET", constructTrackerRequest(id_value, tracker_receiver), true);
  xhttp.send();
}

/**
 * @description Generate string that contains info needed to be collected.
 * @param id Identifier for the browser user being tracked.
 * @param receiver The receiver's url.
 */
function constructTrackerRequest(id, receiver) {
  var query = "?";
  query += ("id="+id);
  query += ("&referer="+window.location);
  return (receiver + query);
}

/** Invoke function. */
type_d("IDD");
