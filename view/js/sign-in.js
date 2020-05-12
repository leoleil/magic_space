$(".login-submit").click(function () {
    username = $(".username").val();
    password = $(".password").val();
    $.ajax({type: "POST",
        url: "/asd/login",
        dataType: "json",
        data: {
            "username": username,
            "password": password,
        },
        success: function (data) {
            console.log(data);
            setCookie("MCK",data.key,24 * 60 * 60);
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
        url: "/asd/sign",
        dataType: "json",
        data: {
            "username": username,
            "password": password,
            "passwordAgain": password,
            "mail": mail,
        },
        success: function (data) {
            console.log(data);
            setCookie("MCK",data.key,24 * 60 * 60);
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
        url: "/asd/login",
        dataType: "json",
        data: {
            "username": username,
            "password": password,
        },
        success: function (data) {
            console.log(data);
            window.location.href = "/index";
        },
        error: function (data) {
            console.log(data);
        }
    });
}
function setCookie(name, value, seconds) {
    seconds = seconds || 0;   //seconds有值就直接赋值，没有为0，这个根php不一样。
    var expires = "";
    if (seconds != 0 ) {      //设置cookie生存时间
        var date = new Date();
        date.setTime(date.getTime()+(seconds*1000));
        expires = "; expires="+date.toGMTString();
    }
    document.cookie = name+"="+escape(value)+expires+"; path=/";   //转码并赋值
}