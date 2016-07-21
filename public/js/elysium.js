/* elysium.js */

console.log('Elysium!');


//  Constants
const AUTOFILL_MIN_CHAR = 3;

//  AUTH
/*  Check Registered User */
$(document).ready(function() {
  $('#user_name, #user_email').on('keyup blur change', function() {
    var target = this;
    if ($(this).val().length > AUTOFILL_MIN_CHAR) {
      var param = [];
      if ($('#user_name').val()) {
        param.push('user_name='+$('#user_name').val());
      }
      if ($('#user_email').val()) {
        param.push('user_email='+$('#user_email').val());
      }
      param = param.join('&');
      $.getJSON('/api/v1/users?'+param, function(d) {
        if (d.users && d.users.length > 0) {
          var invalidUser = false;
          var invalidEmail = false;

          for(var i=0;i<d.users.length;i++) {
            var ref = d.users[i];
            console.log(ref);
            if (ref.user_name === $('#user_name').val()) {
              invalidUser = true;
            }
            if (ref.user_email === $('#user_email').val()) {
              invalidEmail = true;
            }
          }
          if (invalidUser) {
            $('#user_name').addClass('elysium-textfield__error');
          } else {
            $('#user_name').removeClass('elysium-textfield__error');
          }

          if (invalidEmail) {
            $('#user_email').addClass('elysium-textfield__error');
          } else {
            $('#user_email').removeClass('elysium-textfield__error');
          }

          $('#register_submit').attr('disabled','disabled');
        } else {
          $(target).removeClass('elysium-textfield__error');
          $('#register_submit').removeAttr('disabled');
        }
      });
    }
  });  
});
