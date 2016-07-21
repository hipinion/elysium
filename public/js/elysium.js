/* elysium.js */

console.log('Elysium!');


//  Constants
const AUTOFILL_MIN_CHAR = 3;

//  AUTH
/*  Check Registered User */
$(document).ready(function() {
  $('#user_name').on('keyup blur change', function() {
    if ($(this).val().length > AUTOFILL_MIN_CHAR) {
      console.log('checking',$(this).val());
    }
  });  
});
