package main

import (
	"github.com/kataras/iris/v12"
	"huffman-codec/server"
)

func main() {
	app := iris.Default()
	server.Handle(app)

	if err := app.Listen(":24434"); err != nil {
		println("服务器成功启动")
	}
}
