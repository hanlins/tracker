/**
 * @file script.js
 * @description This file contains the script to be sent as response for type F
 *              tracker.
 * @author Hanlin Shi (hanlins)
 * @version 0.1.0
 */

/**
 * @description Initialize the local storage module.
 * @param id_name The identifier's name stored in the cookie.
 * @param id_proposal The proposaled id value, will be used if the cookie has
 *        not been set.
 * @return Storage object.
 */
function createStorage(id_name, id_proposal) {
  var newStore = new SwfStore({
    namespace: "trackerf",
    swf_url: "//trackerf.com/f/storage.swf",
    debug: true,
    onready: function() {
      // if tracker id doesn't exists, assign a new one
      if (!newStore.get(id_name)) {
        newStore.set(id_name, id_proposal);
        console.log("assign new tracker id: " + id_proposal);
      } else {
        console.log(id_name + " id value exists: " + newStore.get(id_name));
      }
    },
    onerror: function(err) {
      console.error(err.message);
    }
  });
  return newStore;
}

/**
 * @description Business logic for type - F tracker.
 * @param id_name The identifier's name stored in the cookie.
 * @param id_proposal The proposaled id value, will be used if the cookie has
 *        not been set.
 * @return Void.
 */
function type_f(id_name, id_proposal) {
  // create new store object
  store = createStorage(id_name, id_proposal);
}

