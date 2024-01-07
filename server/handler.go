package server

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"huffman-codec/server/huffmanCode"
)

func Handle(app *iris.Application) {
	// 绑定静态目录
	app.HandleDir("/", "./webapp")

	// 绑定后端接口
	mvc.New(app.Party("/")).Handle(new(huffmanCode.Controller))
}
