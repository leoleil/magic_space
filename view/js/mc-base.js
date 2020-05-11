$(function () {
    // 验证用户登录情况
    $.ajax({
        type: "POST",
        url: "/asd/check",
        dataType: "json",
        success: function (data) {
            console.log(data);
            $(".navbar-right").append(
                '<li class="dropdown">\n' +
                '<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Add <span class="caret"></span></a>\n' +
                '<ul class="dropdown-menu">\n' +
                '<li><a href="/blog/doc">Blog</a></li>\n' +
                '<li><a href="#">Video</a></li>\n' +
                '<li><a href="#">File</a></li>\n' +
                '</ul>\n' +
                '</li>' +
                '<li><a>' + data.username + '</a></li>'
            );
        },
        error: function () {
            console.log("no");
            $(".navbar-right").append(
                '<li><a href="/join">Join Us</a></li>'
            );
        },
    });
});