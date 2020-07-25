$(function () {
    p = $(".doc").attr("id");
    $.ajax({
        type: "GET",
        url: "/blog/open",
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
                        url: "/blog/delete",
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
            tinymce.init({
                selector: "#mytextarea",
                skin: 'oxide-dark',
                language:'zh_CN',
                plugins: 'image,lists,advlist,autoresize,emoticons,table',
                toolbar: 'undo redo | formatselect fontsizeselect forecolor bold italic | alignleft aligncenter alignright | bullist numlist | image | emoticons | table',
                images_upload_handler: function (blobInfo, succFun, failFun) {
                    var xhr, formData;
                    var file = blobInfo.blob();//转化为易于理解的file对象
                    xhr = new XMLHttpRequest();
                    xhr.withCredentials = false;
                    xhr.open('POST', '/blog/img/upload');
                    xhr.onload = function() {
                        var json;
                        if (xhr.status != 200) {
                            failFun('HTTP Error: ' + xhr.status);
                            return;
                        }
                        json = JSON.parse(xhr.responseText);
                        if (!json || typeof json.location != 'string') {
                            failFun('Invalid JSON: ' + xhr.responseText);
                            return;
                        }
                        succFun("/" + json.location);
                    };
                    formData = new FormData();
                    formData.append('file', file, file.name );
                    xhr.send(formData);
                },
            });
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
