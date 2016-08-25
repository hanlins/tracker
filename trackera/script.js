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
  value = value.replace(/\s/g, '');// remove space
  var parts = value.split(";" + name + "=");
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
 * @return Void.
 */
function type_a(id_name, id_proposal) {
  /** examine whether cookie value exists. */
  var id_value = getCookie(id_name, document);
  /** set the cookie value if it has not been set. */
  if (id_value === undefined) {
    setCookie(id_name, id_proposal);
    id_value = id_proposal;
  }
  /** transmit the id together with other information back to tracker. */
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (xhttp.readyState == 4 && xhttp.status == 200) {
      console.log("tracker A: request sent, with response: " +
		  xhttp.responseText);
    }
  };
  /** test and set footprint. */
  if (!footprintExists()) {
    // notice the hack here, id_proposal instead of id_value is used here since
    // id_proposal is basiclly guaranteed to be different for each request of
    // this script, so it's ideal for a random seed.
    injectFootprint(generateFootprint(id_proposal));
  }
  var tracker_receiver = "http://trackera.com/listenera";
  xhttp.open("GET", constructTrackerRequest(id_value, tracker_receiver), true);
  xhttp.send();
}

/**
 * @description This function helps to set the cookie value.
 * @param id_name The identifier's name stored in the cookie.
 * @param id_proposal The proposaled id value to be set.
 * @return Void.
 */
function setCookie(id_name, id_proposal) {
  /** calculate expiration, currently set to 1 hour. */
  var d = new Date();
  d.setTime(d.getTime() + (60*60*1000));
  var expires = "expires="+ d.toUTCString();
  document.cookie=id_name + "=" + id_proposal + "; " + expires;
}

/**
 * @description Generate string that contains info needed to be collected.
 * @param id Identifier for the browser user being tracked.
 * @param receiver The receiver's url.
 * @return Void.
 */
function constructTrackerRequest(id, receiver) {
  var query = "?";
  query += ("id=" + id);
  query += ("&referer=" + window.location);
  query += ("&footprint=" + extractFootprint());
  return (receiver + query);
}

/**
 * @description Inject cross-site footprint by modifying window.name attribute.
 *              When user redirected to other location using same window, then
 *              this value will be kept.
 * @param footprint Footprint identifier string.
 * @return Void.
 */
function injectFootprint(footprint) {
  window.name = footprint;
}

/**
 * @description Check whether cross-site footprint exists.
 * @return True if footprint exists, false if not exists.
 */
function footprintExists() {
  if (window.name === undefined || window.name === "") {
    return false;
  } else {
    return true;
  }
}

/**
 * @description Get the cross-site footprint, provided that it do exists.
 * @return Footprint string.
 */
function extractFootprint() {
  return window.name;
}

/**
 * @description Generate cross-site footprint.
 * @param seed The seed string used to generate footprint string.
 * @return Newly generated footprint string.
 */
function generateFootprint(seed) {
  //return Base64.encode(seed);
  return randomString(15, seed);
}

/**
 * @description Generate random string.
 * @param len Length of the random string.
 * @param seed Used as seed / naunce.
 * @return Newly generated random string.
 */
function randomString(len, seed) {
  var rand = "";
  var dic = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + seed;
  for( var i = 0; i < len; i++) {
    rand += dic.charAt(Math.floor(Math.random() * dic.length));
  }
  return rand;
}
