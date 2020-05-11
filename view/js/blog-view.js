$(function () {
    p = $(".doc").attr("id");
    $.ajax({
        type: "GET",
        url: "http://localhost:4010/blog/open",
        dataType: "json",
        data:{
            "id":p,
        },
        success: function (data) {
            console.log(data);
            addBlog(data.entity);
            if(data.edit){
                console.log("test");
                $(".jumbotron").append('<a class="btn btn-primary" href="/blog/edit?id=' + p + '" role="button">edit</a>');
            }
        },
        error: function () {
            console.log("no");
        },
    });
})
function addBlog(entity) {
    $(".headline").append(entity.title);
    $("#time").append(entity.gmt_modified);
    $("#body").append(entity.body);

}
