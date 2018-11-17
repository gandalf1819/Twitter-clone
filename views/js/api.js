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

});

function addPost(userId ,status){
    var data={
        "userId": userId,
        "status" : status,
    }
    var url ="http://localhost:9090/posts/"

    $.ajax({
        url: url,
        method: "POST",
        data: JSON.stringify(data),
        dataType:'json',
        contentType:"application/json",
        success: function(data) {
            //action = (action == "Follow")? "UnFollow":"Follow";
            //event.target.setAttribute("onclick","followUser("+userId+",'"+followerId+"','"+action+"',event)");    
            //event.target.innerHTML = action
        },
    });
}

function followUser(userId ,followerId, action, event){
    var data={
        "userId": userId,
        "followerId" : parseInt(followerId),
    }
    var url ="http://localhost:9090/"
    url+= (action == "Follow")? "follow/":"unfollow/";
    $.ajax({
        url: url,
        method: "POST",
        data: JSON.stringify(data),
        dataType:'json',
        contentType:"application/json",
        success: function(data) {
            action = (action == "Follow")? "UnFollow":"Follow";
            event.target.setAttribute("onclick","followUser("+userId+",'"+followerId+"','"+action+"',event)");    
            event.target.innerHTML = action
        },
    });
}