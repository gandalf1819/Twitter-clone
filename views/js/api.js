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
                //var msg = $('#inputFirstName').val();
                toatr.success(data.firstName);
            },
        });
    });

});

function followUser(followerId, action, event){
    var data={
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
            event.target.setAttribute("onclick","followUser('"+followerId+"','"+action+"',event)");    
            event.target.innerHTML = action
        },
    });
}

function signIn(){
    var data={
        "email": document.getElementById("inputLoginEmail").value,
        "password" : document.getElementById("inputLoginPassword").value,
    }
    var url ="http://localhost:9090/login/";
    $.ajax({
        url: url,
        method: "POST",
        data: JSON.stringify(data),
        dataType:'json',
        contentType:"application/json",
        success: function(data) {
            if(data.Status == 200){
                location.href="http://localhost:9090/posts/";
            }
        },
    });
}


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
