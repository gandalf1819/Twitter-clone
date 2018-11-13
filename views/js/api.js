$(document).ready(function() {
    
    $("#register-button").on('click', function() {
        console.log("API CALLED")
        var registerData={
            "firstName": $("#inputFirstName").val(),
            "lastName" : $("#inputLastName").val(),
            "email": $("#inputRegisterEmail").val(),
            "password":$("#inputRegisterPassword").val()
        }
        $.ajax({
            url: "http://localhost:9090/register/",
            method: "POST",
            data: JSON.stringify(registerData),
            dataType:'json',
            contentType:"application/json",
            success: function(data) {
               // $("#response").html(data);
            },
        });
    });
});