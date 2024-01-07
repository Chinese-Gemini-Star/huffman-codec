package huffmanCode

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Controller struct {
	Ctx iris.Context
}

// returnTextResponse 返回文本响应
func (c *Controller) returnTextResponse(code int, text string) mvc.Response {
	fmt.Println(text)
	return mvc.Response{
		Code: code,
		Text: text,
	}
}
