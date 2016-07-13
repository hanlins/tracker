/**
 * @file script.js
 * @description This file contains the script to be sent as response for type C
 *              tracker.
 * @author Hanlin Shi (hanlins)
 * @version 0.1.0
 */

/**
 * @description This function set cookie in the first - party page, it will
 *              1). examine whether the cookie has already been set.
 *              2). set a new id in first-party popup if has not been set, then
 *                  reload the page.
 * @param id_name The identifier's name stored in the cookie.
 * @param set_cookie_url The url that open in popup and set cookie for tracker.
 * @param real_page The real content we want to display.
 * @return Void.
 */
function type_c(id_name, set_cookie_url, real_page) {
  /** set the cookie value in popup. */
  console.log("create new cookie");
  setCookiePopup(set_cookie_url);
  /** reload the page after 10 second */
  //deferReloadPage(10, real_page);
}

/**
 * @description In this function, a pop-up page will be opened and a cookie
 *              will be set in that page.
 * @param set_cookie_url URL of the pop-up that directs the browser to set cookie.
 * @return Void.
 */
function setCookiePopup(set_cookie_url) {
  // check whether popup is disabled,
  // if not, open the popup
  // if disabled, wait for user to click on anything in the page
  var newWin = window.open(set_cookie_url);
  if(!newWin || newWin.closed || typeof newWin.closed=='undefined') { 
    $("body").click(function(e) {
      e.preventDefault();
      window.open(set_cookie_url, '_blank');
    });
  }
}

/**
 * @description Reload the page / iframe after 'delay' seconds.
 * @param delay Waiting time for reloading the page in unit of seconds.
 * @param real_page The real content we want to display.
 * @return Void.
 */
function deferReloadPage(delay, real_page) {
  setTimeout(function(){window.location = real_page;}, 1000 * delay);
}

/** invoke type - C tracker's script. */
type_c("IDC", "http://trackerc.com/c/setcookiec", "http://trackerc.com/c/base.html");
