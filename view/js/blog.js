$(function () {
    p = $(".p").attr("id");
    $.ajax({
        type: "GET",
        url: "http://localhost:4010/blog/list",
        dataType: "json",
        data:{
            "page":p,
        },
        success: function (data) {
            console.log(data);
            addBlogList(data.list, data.page, 1, data.endPage, data.limit)
        },
        error: function () {
            console.log("no");
        },
    });
})
function addBlogList(blogList, pageNum, fistPage, endPage, limit) {
    console.log(blogList);
    for(var i in blogList) {
        $("#blog_list").append(
            '<a href="/blog/view?id='+ blogList[i].id +'" class="list-group-item">\n' +
            '<h3 class="list-group-item-heading">' + blogList[i].title + '</h3>\n' +
            '<p class="list-group-item-text"> '+ blogList[i].gmt_modified +'</p>\n' +
            '</a>'
        );
    }
   if(pageNum > fistPage && pageNum < endPage){
       $(".pagination").append(
           '<li>\n' +
           '<a href="/blog?page='+ (pageNum - 1) +'" aria-label="Previous">\n' +
           '<span aria-hidden="true">&laquo;</span>\n' +
           '</a>\n' +
           '</li>'
       );
       for(i = fistPage; i <= endPage;i++){
           if(i == pageNum){
               $(".pagination").append(
                   '<li class="active"><a href="/blog?page=' + i + '">'+ i +'</a></li>\n'
               );
           }else{
               $(".pagination").append(
                   '<li><a href="/blog?page=' + i + '">'+ i +'</a></li>\n'
               );
           }
       }
       $(".pagination").append(
           '<li>\n' +
           '<a href="/blog?page='+ (pageNum + 1) + '" aria-label="Next">\n' +
           '<span aria-hidden="true">&raquo;</span>\n' +
           '</a>\n' +
           '</li>'
       );
   }else if(pageNum == fistPage && pageNum != endPage){
       for(i = fistPage; i <= endPage;i++){
           if(i == pageNum){
               $(".pagination").append(
                   '<li class="active"><a href="/blog?page=' + i + '">'+ i +'</a></li>\n'
               );
           }else{
               $(".pagination").append(
                   '<li><a href="/blog?page=' + i + '">'+ i +'</a></li>\n'
               );
           }
       }
       $(".pagination").append(
           '<li>\n' +
           '<a href="/blog?page='+ (pageNum + 1) + '" aria-label="Next">\n' +
           '<span aria-hidden="true">&raquo;</span>\n' +
           '</a>\n' +
           '</li>'
       );
   }else if(pageNum != fistPage && pageNum == endPage){
       $(".pagination").append(
           '<li>\n' +
           '<a href="/blog?page='+ (pageNum - 1) +'" aria-label="Previous">\n' +
           '<span aria-hidden="true">&laquo;</span>\n' +
           '</a>\n' +
           '</li>'
       );
       for(i = fistPage; i <= endPage;i++){
           if(i == pageNum){
               $(".pagination").append(
                   '<li class="active"><a href="/blog?page=' + i + '">'+ i +'</a></li>\n'
               );
           }else{
               $(".pagination").append(
                   '<li><a href="/blog?page=' + i + '">'+ i +'</a></li>\n'
               );
           }
       }
   }
}
