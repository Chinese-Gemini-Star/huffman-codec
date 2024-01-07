package huffmanCode

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
	"strings"
)

// PostCipher 解压缩霍夫曼编码
func (c *Controller) PostCipher(message Message) mvc.Result {
	fmt.Println("解码字符串:", message.Text, ",字符集:", message.CharsetID)

	// 转为二进制
	message.Text = hex2bin(message.Text)
	fmt.Println("二进制密文为:", message.Text)

	if huffmanTree, err := ReadHuffmanTree(message.CharsetID); err == nil {
		huffmanCode := huffmanTree.GetHuffmanCode()

		// 翻转键值,用以解码
		huffmanDecode := make(map[string]string)
		for key, value := range huffmanCode {
			huffmanDecode[value] = key
		}

		// 解码密文
		res := strings.Builder{}
		tmp := strings.Builder{}
		isEOF := false
		for _, c := range message.Text {
			// 并不是01串
			if string(c) != "0" && string(c) != "1" {
				fmt.Println("密文不正确")
				return mvc.Response{
					Code: iris.StatusBadRequest,
					Text: "密文不正确",
				}
			}
			tmp.WriteString(string(c))
			if word, ok := huffmanDecode[tmp.String()]; ok {
				// 遇到中止符,结束解码
				if word == "\u0000" {
					isEOF = true
					break
				}
				// 当前对应一个字符,存入结果
				res.WriteString(word)
				tmp.Reset()
			}
		}
		fmt.Println(res.String())
		// 没有终止符
		if !isEOF {
			fmt.Println("密文不正确")
			return mvc.Response{
				Code: iris.StatusBadRequest,
				Text: "密文不正确",
			}
		}
		fmt.Println("霍夫曼解码结果:", res.String())
		if result, err := json.Marshal(Message{
			Text:      res.String(),
			CharsetID: message.CharsetID,
		}); err != nil {
			fmt.Println("响应结果JSON编码失败")
			return mvc.Response{
				Code: iris.StatusInternalServerError,
				Text: "响应结果JSON编码失败",
			}
		} else {
			return mvc.Response{
				Code:        iris.StatusOK,
				ContentType: "application/json",
				Content:     result,
			}
		}
	} else {
		fmt.Println("霍夫曼编码读取失败")
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Text: "霍夫曼编码读取失败",
		}
	}
}

func hex2bin(hex string) string {
	res := strings.Builder{}
	for _, c := range hex {
		base, _ := strconv.ParseUint(string(c), 16, 0)
		bin := strconv.FormatUint(base, 2)

		// 补齐4位
		diff := 4 - len(bin)
		for i := 0; i < diff; i++ {
			res.WriteString("0")
		}
		res.WriteString(bin)
	}
	return res.String()
}
