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
