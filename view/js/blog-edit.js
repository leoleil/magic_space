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
                $(".jumbotron").append('<bottom class="btn btn-danger" id="delete">delete</bottom>');
                $("#delete").click(function () {
                    p = $(".doc").attr("id");
                    $.ajax({
                        type: "POST",
                        url: "http://localhost:4010/blog/delete",
                        dataType: "json",
                        data:{
                            "id":p,
                        },
                        success: function (data) {
                            alert("删除成功");
                            console.log(data);
                            window.location.href = "/index";
                        },
                        error: function () {
                            console.log("no");
                            alert("删除出错");
                        },
                    });
                });
            }
        },
        error: function () {
            console.log("no");
        },
    });
})
function addBlog(entity) {
    $(".headline").append(entity.title);
    $("#mytextarea").val(entity.body);
    //$("#time").append(entity.gmt_modified);
    $("#headline").val(entity.title);
    $(".blog-id").val(entity.id);
}
$("#headline").bind("input propertychange",function(event){
    $(".headline").empty();
    $(".headline").append($("#headline").val());
});
