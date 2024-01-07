$(function () {
    $("#postOriginal").click(function (e) {
        // 获取当前字符集ID
        const charsetId = charset == "auto" ? "" : $("#currentCharsetId").text();
        // 生成消息JSON
        const message = {
            Text: $("#originalText").val(),
            CharsetID: charsetId,
        };
        const jsonMessage = JSON.stringify(message);
        console.log(jsonMessage);
        // ajax调用后端接口
        $.ajax({
            type: "POST",
            url: "/original",
            contentType: "application/json",
            data: jsonMessage,
            dataType: "text",
            async: false,
            success: function (result, status, xhr) {
                console.log(result);
                result = JSON.parse(result);
                $("#cipherText").val(result.Text);
                $("#currentCharsetId").text(result.CharsetID);
                $("p:has(#currentCharsetId)").attr("hidden", false);
            },
            error: function (xhr, status, error) {
                $("#cipherText").val("");
                alert(status + ":" + xhr.responseText);
            }
        });
    });
})