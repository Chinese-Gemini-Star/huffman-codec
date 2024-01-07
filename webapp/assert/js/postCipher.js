$(function () {
    $("#postCipher").click(function (e) {
        err = false;
        // 获取当前字符集ID
        const charsetId = $("#currentCharsetId").text();

        // 无字符集ID无法译码
        if (charsetId == "") {
            $("#originalText").val("");
            alert("译码时必须给出确定的字符集,在此前没有编码的情况下不可根据原文自动生成");
            return;
        }

        // 生成消息JSON
        const message = {
            Text: $("#cipherText").val(),
            CharsetID: charsetId,
        };
        const jsonMessage = JSON.stringify(message);
        console.log(jsonMessage);

        // ajax调用后端接口
        $.ajax({
            type: "POST",
            url: "/cipher",
            contentType: "application/json",
            data: jsonMessage,
            dataType: "text",
            async: false,
            success: function (result, status, xhr) {
                result = JSON.parse(result);
                $("#originalText").val(result.Text);
                $("#currentCharsetId").text(result.CharsetID);
                $("p:has(#currentCharsetId)").attr("hidden", false);
            },
            error: function (xhr, status, error) {
                $("#originalText").val("");
                alert(status + ":" + xhr.responseText);
            }
        });
    });
})