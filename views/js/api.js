$(document).ready(function() {
    
    $("#register-button").on('click', function() {
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

    $("#follow-button").on('click', function() {
        var followerData={
            "userId": 1,
            "followerId" : 2,
        }
        $.ajax({
            url: "http://localhost:9090/follow/",
            method: "POST",
            data: JSON.stringify(followerData),
            dataType:'json',
            contentType:"application/json",
            success: function(data) {
               // $("#response").html(data);
            },
        });
    });

    $("#unfollow-button").on('click', function() {
        var unfollowerData={
            "userId": 1,
            "followerId" : 2,
        }
        $.ajax({
            url: "http://localhost:9090/unfollow/",
            method: "POST",
            data: JSON.stringify(unfollowerData),
            dataType:'json',
            contentType:"application/json",
            success: function(data) {
               // $("#response").html(data);
            },
        });
    });
});