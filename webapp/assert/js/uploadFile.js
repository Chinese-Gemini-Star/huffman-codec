// 按钮点击转为上传文件,并显示文件名和文件内容
$(function () {
    function getExtension(name) {
        return name.substring(name.lastIndexOf("."));
    }
    // 转为上传文件
    $("#originalFileUpload").click(function () {
        $("#originalFileUpload + input").click();
    });
    $("#cipherFileUpload").click(function () {
        $("#cipherFileUpload + input").click();
    });

    // 显示源码文件名和文件内容
    $("#originalFileUpload + input").change(function () {
        reader = new FileReader();
        const file = $("#originalFileUpload + input").prop('files')[0];

        // 根据后缀名判断文件类型
        if (getExtension(file.name) != ".txt") {
            alert("文件类型需要为纯文本(text/plain),文件后缀名需为.txt");
            return;
        } else {
            $("#originalFileUpload ~ .fileName").text(file.name);
            reader.readAsText(file);
            reader.addEventListener("loadend", function () {
                $("#originalText").val(reader.result);
            });
        }
    });

    // 显示译码文件名和文件内容
    $("#cipherFileUpload + input").change(function () {
        reader = new FileReader();
        const file = $("#cipherFileUpload + input").prop('files')[0];

        // 根据后缀名判断文件类型
        if (getExtension(file.name) == ".txt") {
            // 文本文档
            $("#cipherFileUpload ~ .fileName").text(file.name);
            reader.readAsText(file);
            reader.addEventListener("loadend", function () {
                $("#cipherText").val(reader.result);
            });
        } else if (getExtension(file.name) == ".hex") {
            // 字节文件
            $("#cipherFileUpload ~ .fileName").text(file.name);
            reader.readAsArrayBuffer(file);
            reader.addEventListener("loadend", function () {
                cipher = "";
                new Uint8Array(reader.result).forEach((value, index) => {
                    let num = parseInt(value).toString(16);
                    cipher += num;
                });
                $("#cipherText").val(cipher);
            });
        } else {
            alert("文件类型需要为纯文本(text/plain)或字节文件(application/octet-stream), 文件后缀需对应为.txt或.hex");
            return;
        }
    });
});