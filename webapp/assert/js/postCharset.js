// 发送自定义字符集
$(function () {
    $("#charsetInput .ok").click(function () {
        $.ajax({
            type: "POST",
            url: "/charset",
            contentType: "application/json",
            data: $("#charsetInput textarea").val(),
            dataType: "text",
            success: function (result, status, xhr) {
                $("#currentCharsetId").text(result);
                $("p:has(#currentCharsetId)").attr("hidden", false);
            },
            error: function (xhr, status, error) {
                alert(status + ":" + xhr.responseText);
            }
        });
    });
});

$(function () {
    $("#charsetId").change(function (e) {
        $("#currentCharsetId").text($("#charsetId").val());
        $("p:has(#currentCharsetId").attr("hidden", false);
    });
});