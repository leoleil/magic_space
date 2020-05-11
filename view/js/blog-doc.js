$("#headline").bind("input propertychange",function(event){
    $(".headline").empty();
    $(".headline").append($("#headline").val());
});
$("#submit").bind("click",function () {
    from = new FormData($("#blog-form")[0]);

    console.log($("#blog-form").serialize());
    /*$.ajax({
        type: "POST",
        url: "http://localhost:4010/blog/upload",
        dataType: "json",
        data:from,
        processData: false,
        contentType: false,
        success: function (data) {
            console.log(data);
        },
        error: function () {
            console.log("no");
        },
    });*/
})