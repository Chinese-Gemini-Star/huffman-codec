$(function () {
    // 发送密文
    $("#send").click(function () {
        $("#postOriginal").click();
        const content = $("#cipherText").val();
        if (content == "") return;

        const username = prompt("注意,本功能为实验性功能.\n目标用户名:", "");
        if (username == null) return;

        // 生成消息JSON
        const message = {
            Text: $("#cipherText").val(),
            CharsetID: $("#currentCharsetId").text(),
            Username: username,
        };
        const jsonMessage = JSON.stringify(message);

        // ajax调用后端接口
        $.ajax({
            type: "POST",
            url: "/send",
            contentType: "application/json",
            data: jsonMessage,
            dataType: "text",
            async: false,
            success: function (result, status, xhr) {
                alert(result);
            },
            error: function (xhr, status, error) {
                alert(status + ":" + xhr.responseText);
            }
        });
    });

    // 接收密文
    $("#receive").click(function () {
        const username = prompt("注意,本功能为实验性功能.\n目前只能获取最新接收到的密文.\n用户名:", "");
        if (username == null) return;

        // ajax调用后端接口
        $.ajax({
            type: "GET",
            url: "/receive/" + username,
            dataType: "text",
            async: false,
            success: function (result, status, xhr) {
                console.log(result);
                result = JSON.parse(result);

                // 不存在此密文
                if (result.CharsetID == "") {
                    return;
                }
                // 显示密文
                $("#cipherText").val(result.Text);
                // 字符集更改为已生成字符集
                $("#existed").attr("checked", true);
                $("input[name='charset']").click();
                $("#charsetId").val(result.CharsetID);
                $("#charsetId").change();
            },
            error: function (xhr, status, error) {
                alert(status + ":" + xhr.responseText);
            }
        });

        if (charset != "existed") {
            alert("此用户没有收到密文");
            return;
        }
        $("#postCipher").click();
    });
});