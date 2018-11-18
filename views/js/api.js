
function registerUser() {
    var registerData = {
        "firstName": $("#inputFirstName").val(),
        "lastName": $("#inputLastName").val(),
        "email": $("#inputRegisterEmail").val(),
        "password": $("#inputRegisterPassword").val()
    }
    $.ajax({
        url: "http://localhost:9090/register/",
        method: "POST",
        data: JSON.stringify(registerData),
        dataType: 'json',
        contentType: "application/json",
        success: function (data) {
            if (data.Status == 200) {
                toastr.success(data.Message);
            } else {
                toastr.error(data.Message)
            }

        },
    });
}

function followUser(followerId, action, event) {
    var data = {
        "followerId": parseInt(followerId),
    }
    var url = "http://localhost:9090/"
    url += (action == "Follow") ? "follow/" : "unfollow/";
    $.ajax({
        url: url,
        method: "POST",
        data: JSON.stringify(data),
        dataType: 'json',
        contentType: "application/json",
        success: function (data) {
            action = (action == "Follow") ? "UnFollow" : "Follow";
            event.target.setAttribute("onclick", "followUser('" + followerId + "','" + action + "',event)");
            event.target.innerHTML = action
        },
    });
}

function signIn() {
    var data = {
        "email": document.getElementById("inputLoginEmail").value,
        "password": document.getElementById("inputLoginPassword").value,
    }
    var url = "http://localhost:9090/login/";
    $.ajax({
        url: url,
        method: "POST",
        data: JSON.stringify(data),
        dataType: 'json',
        contentType: "application/json",
        success: function (data) {
            if (data.Status == 200) {
                toastr.success(data.Message)
                location.href = "http://localhost:9090/posts/";
            }
            else {
                toastr.error(data.Message)
            }
        },
    });
}


function addPost() {
    var status = document.getElementById("status").value
    if (!status) {
        toastr.error("Please enter status!!")
        return
    }
    var data = {
        "status": status
    }
    var url = "http://localhost:9090/posts/"

    $.ajax({
        url: url,
        method: "POST",
        data: JSON.stringify(data),
        dataType: 'json',
        contentType: "application/json",
        success: function (data) {
            if (data.Status == 200) {
                toastr.success(data.Message);
            } else {
                toastr.error(data.Message)
            }
        },
    });
}
