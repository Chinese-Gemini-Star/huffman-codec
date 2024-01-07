// 两个多行文本框高度自动对齐
$(function () {
    $("#originalText").resize(function () {
        $("#cipherText").height($("#originalText").height());
    });
    $("#cipherText").resize(function () {
        $("#originalText").height($("#cipherText").height());
    });
});

let charset = "auto";
// 指定单选框选中后显示额外内容
$(function () {
    $("input[name='charset']").click(function () {
        charset = $("input[name='charset']:checked").val();
        console.log("当前字符集来源为:" + charset);
        $("p:has(#currentCharsetId)").attr("hidden", true);

        $("#customize ~ button").attr("hidden", true);
        $("div:has(>#charsetId)").attr("hidden", true);

        if (charset == "customize") {
            $("#customize ~ button").attr("hidden", false);
        } else if (charset == "existed") {
            $("div:has(>#charsetId)").attr("hidden", false);
        } else if (charset == "default") {
            $("#currentCharsetId").text("19d9059a3d6da446e44d6b811869c766");
            $("p:has(#currentCharsetId)").attr("hidden", false);
        } else {
            $("#currentCharsetId").text("");
            $("p:has(#currentCharsetId)").attr("hidden", true);
        }
    });
});

// 显示与隐藏弹窗
$(function () {
    // 显示自定义字符集输入弹窗
    $("#customize ~ button").click(function () {
        $("#charsetInput").attr("hidden", false);

        // 禁用输入框
        $("#originalText").attr("disabled", true);
        $("#cipherText").attr("disabled", true);
        // 禁用所有按钮
        $(".main :is(input, button)").attr("disabled", true);
    });

    // 隐藏自定义字符集输入弹窗
    $("#charsetInput button").click(function () {
        $("#charsetInput").attr("hidden", true);

        // 启用输入框
        $("#originalText").attr("disabled", false);
        $("#cipherText").attr("disabled", false);
        // 启用所有按钮
        $(".main :is(input, button)").attr("disabled", false);
    });

    // 显示霍夫曼树展示弹窗
    $("#showHuffmanTree").click(function () {
        $("#huffmanTree").attr("hidden", false);
        // 禁用输入框
        $("#originalText").attr("disabled", true);
        $("#cipherText").attr("disabled", true);
        // 禁用所有按钮
        $(".main :is(input, button)").attr("disabled", true);

        if (charset == "auto" && $("#currentCharsetId").text() == "") {
            // 生成消息JSON
            const message = {
                Text: $("#originalText").val(),
                CharsetID: "",
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
                    $("#currentCharsetId").text(result.CharsetID);
                    $("p:has(#currentCharsetId)").attr("hidden", false);
                    // 获取图片
                    $("#huffmanTree img").attr("src", "/tree/png/" + $("#currentCharsetId").text());
                },
                error: function (xhr, status, error) {
                    alert(status + ":" + xhr.responseText);
                }
            });
        } else {
            // 获取图片
            $("#huffmanTree img").attr("src", "/tree/png/" + $("#currentCharsetId").text());
        }
    });

    // 隐藏霍夫曼树展示弹窗
    $("#huffmanTree button").click(function () {
        $("#huffmanTree").attr("hidden", true);

        // 启用输入框
        $("#originalText").attr("disabled", false);
        $("#cipherText").attr("disabled", false);
        // 启用所有按钮
        $(".main :is(input, button)").attr("disabled", false);
    })
});
