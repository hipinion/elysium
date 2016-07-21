/* elysium.js */

console.log('Elysium!');


//  Constants
const AUTOFILL_MIN_CHAR = 3;

//  AUTH
/*  Check Registered User */
$(document).ready(function() {
  $('#user_name').on('keyup blur change', function() {
    var target = this;
    if ($(this).val().length > AUTOFILL_MIN_CHAR) {
      $.getJSON('/api/v1/users?user_name='+$(this).val(), function(d) {
        if (d.users && d.users.length > 0) {
          $(target).addClass('elysium-textfield__error');
          $('#register_submit').attr('disabled','disabled');
        } else {
          $(target).removeClass('elysium-textfield__error');
          $('#register_submit').removeAttr('disabled');
        }
      });
    }
  });  
});
