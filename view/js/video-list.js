$(function () {
    $.ajax({
        type: "GET",
        url: "/video/list",
        dataType: "json",
        success: function (data) {
            console.log(data);
            addVideoList(data.list)
        },

        error: function () {
            console.log("no");
        },
    });
})
function addVideoList(list) {
    console.log(list);
    for(var i in list) {
        $("#video_list").append(
            '<a href="/video/open?video='+ list[i] +'" class="list-group-item">\n' +
            '<p class="list-group-item-heading">' + list[i] + '</p>' +
            '</a>'
        );
    }
}