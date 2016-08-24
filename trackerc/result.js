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

var tracker_name = "tkc";
var tracker_id = getCookie("IDC", document);