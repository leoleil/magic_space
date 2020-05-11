$(".login-submit").click(function () {
    username = $(".username").val();
    password = $(".password").val();
    $.ajax({type: "POST",
        url: "http://localhost:4010/asd/login",
        dataType: "json",
        data: {
            "username": username,
            "password": password,
        },
        success: function (data) {
            window.location.href = "/index";
        },
        error: function (data) {
            console.log(data);
        }
    });
});
$(".signin-submit").click(function () {
    username = $(".username").val();
    password = $(".password").val();
    mail = $(".email").val();
    $.ajax({type: "POST",
        url: "http://localhost:4010/asd/sign",
        dataType: "json",
        data: {
            "username": username,
            "password": password,
            "passwordAgain": password,
            "mail": mail,
        },
        success: function (data) {
            console.log(data);
            window.location.href = "/index";
        },
        error: function (data) {
            console.log(data);
        }
    });
});
function login(){
    username = $(".username").val();
    password = $(".password").val();
    $.ajax({type: "POST",
        url: "http://localhost:4010/asd/login",
        dataType: "json",
        data: {
            "username": username,
            "password": password,
        },
        success: function (data) {
            window.location.href = "/index";
        },
        error: function (data) {
            console.log(data);
        }
    });
}