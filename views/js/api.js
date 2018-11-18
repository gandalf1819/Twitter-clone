
function registerUser() {

    var firstName = document.getElementById("inputFirstName").value;
    var lastName = document.getElementById("inputLastName").value;
    var email = document.getElementById("inputRegisterEmail").value;
    var password = document.getElementById("inputRegisterPassword").value;

    if (!firstName) {
        toastr.error("Please enter First Name!")
        return
    }
    if(!lastName) {
        toastr.error("Please enter Last Name!")
        return
    }
    if(!email) {
        toastr.error("Please enter Email!")
        return
    }
    if(!password) {
        toastr.error("Please enter Password!")
        return
    }

    var registerData = {
        "firstName": firstName,
        "lastName": lastName,
        "email": email,
        "password": password
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
    var email = document.getElementById("inputLoginEmail").value;
    var password = document.getElementById("inputLoginPassword").value;

    if(!email) {
        toastr.error("Please enter Email!")
        return
    }
    if(!password) {
        toastr.error("Please enter Password!")
        return
    }

    var data = {
        "email": email,
        "password": password,
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

function logout() {
    var url = "http://localhost:9090/logout/"
    $.ajax({
        url: url,
        method: "DELETE",
        dataType: 'json',
        contentType: "application/json",
        success: function (data) {
            if (data.Status == 200) {
                toastr.success(data.Message);
                location.href = "http://localhost:9090/login/";
            } else {
                toastr.error(data.Message)
            }
        },
    });
}
