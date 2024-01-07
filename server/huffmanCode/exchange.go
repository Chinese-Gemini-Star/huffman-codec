package huffmanCode

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"huffman-codec/server/database"
)

// PostSend 发送密文至指定用户
func (c *Controller) PostSend(message Message) mvc.Result {
	fmt.Println("向", message.Username, "发送密文", message.Text, ",字符集ID:", message.CharsetID)
	if db, err := database.GetDB(); err == nil {
		db.AutoMigrate(&message)
		db.Create(&message)
		return mvc.Response{
			Code: iris.StatusOK,
			Text: "密文发送成功",
		}
	} else {
		fmt.Println("数据库链接失败")
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Text: "数据库链接失败",
		}
	}
}

// GetReceiveBy 接收指定用户的密文
func (c *Controller) GetReceiveBy(username string) mvc.Result {
	fmt.Println("获取", username, "所收到的密文")
	if db, err := database.GetDB(); err == nil {
		var message Message
		db.AutoMigrate(&message)

		db.Last(&message, "username = ?", username)
		if res, err := json.Marshal(message); err == nil {
			return mvc.Response{
				Code:        iris.StatusOK,
				ContentType: "application/json",
				Content:     res,
			}
		} else {
			fmt.Println("结果序列化失败")
			return mvc.Response{
				Code: iris.StatusInternalServerError,
				Text: "结果序列化失败",
			}
		}
	} else {
		fmt.Println("数据库链接失败")
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Text: "数据库链接失败",
		}
	}
}
