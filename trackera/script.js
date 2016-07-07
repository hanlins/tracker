/**
 * @file script.js
 * @description This file contains the script to be sent as response for type A
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
  var parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}

/**
 * @description This function set cookie in the first - party page, it will
 *              1). examine whether the cookie has already been set.
 *              2). set a new id in cookie if has not been set.
 *              3). transfet the identifier together with other information
 *                  back.
 * @param id_name The identifier's name stored in the cookie.
 * @param id_proposal The proposaled id value, will be used if the cookie has
 *        not been set.
 */
function type_a(id_name, id_proposal) {
  /** examine whether cookie value exists. */
  var id_value = getCookie(id_name, document);
  /** set the cookie value if it has not been set.*/
  if (id_value === null) {
    setCookie(id_name, id_proposal);
    id_value = id_proposal;
  }
  /** transmit the id together with other information back to tracker. */
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (xhttp.readyState == 4 && xhttp.status == 200) {
      console.log("tracker A: request sent, with response: " + xhttp.responseText);
    }
  };
  var tracker_receiver = "http://trackera.com/listenera";
  xhttp.open("GET", constructTrackerRequest(id_value, tracker_receiver), true);
  xhttp.send();
}

/**
 * @description This function helps to set the cookie value.
 * @param id_name The identifier's name stored in the cookie.
 * @param id_proposal The proposaled id value to be set.
 */
function setCookie(id_name, id_proposal) {
  /** calculate expiration, currently set to 1 hour */
  var d = new Date();
  d.setTime(d.getTime() + (60*60*1000));
  var expires = "expires="+ d.toUTCString();
  document.cookie=id_name + "=" + id_proposal + "; " + expires;
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
