/**
 * @file script.js
 * @description This file contains the script to be sent as response for type F
 *              tracker.
 * @author Hanlin Shi (hanlins)
 * @version 0.1.0
 */

/**
 * @description Initialize the local storage module.
 * @return Void.
 */
function init() {
  if (!store.enabled) {
    console.log('Local storage is not supported.');
  } else {
    console.log('Local storage is supported, module initialization success.');
  }
}

/**
 * @description Check whether tracker id exists.
 * @param id_name Name of the tracker id.
 * @return True if exists, false if not.
 */
function checkTrackerIdExists(id_name) {
  if (store.get(id_name)) {
    return true;
  }
  return false;
}

/**
 * @description Set flash cookie id.
 * @param id_name Name of the tracker id.
 * @param id_val Tracker id value.
 * @return Void.
 */
function setTrackerId(id_name, id_val) {
  store.set(id_name, id_val);
}

/**
 * @description Business logic for type - F tracker.
 * @param id_name The identifier's name stored in the cookie.
 * @param id_proposal The proposaled id value, will be used if the cookie has
 *        not been set.
 * @return Void.
 */
function type_f(id_name, id_proposal) {
  if (!checkTrackerIdExists(id_name)) {
    console.log("Tracker ID not set yet, get a new one.");
    setTrackerId(id_name, id_proposal);
  }
  console.log("Tracker ID has been set, which is: " + store.get(id_name));
}

// initialize the module
init();
