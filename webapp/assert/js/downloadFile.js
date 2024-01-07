// 下载文件
$(function () {
    // 译码并下载原文
    $("#downloadOriginal").click(function () {
        $("#postCipher").click();
        const content = $("#originalText").val();
        if (content == "") return;

        // 创建链接及对应文件
        const blob = new Blob([content], { type: "text/plain" });
        const link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);
        const filename = prompt("文件名(无需后缀):", "原文");
        if (filename == null) return;
        link.download = filename + ".txt";

        // 下载文件
        link.click();
    });

    $("#downloadCipher").click(function () {
        $("#postOriginal").click();
        const content = $("#cipherText").val();
        if (content == "") return;


        // 创建二进制流,写入密文
        const buffer = new Uint8Array(content.length / 2);
        for (let i = 0; i < content.length / 2; i++) {
            buffer[i] = parseInt(content.substring(2 * i, 2 * i + 2), 16); // 每次取出两位
        }

        // 创建链接及对应文件
        var blob = new Blob([buffer], { type: "application/octet-stream" });
        var link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);
        const filename = prompt("文件名(无需后缀):", "密文");
        if (filename == null) return;
        link.download = filename + ".hex";

        // 下载文件
        link.click();
    });
});