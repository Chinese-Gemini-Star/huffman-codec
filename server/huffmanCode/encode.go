package huffmanCode

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
	"strings"
)

// PostOriginal 前端发送待压缩原文
func (c *Controller) PostOriginal(message Message) mvc.Result {
	fmt.Println("编码字符串:", message.Text, ",字符集:", message.CharsetID)
	if message.CharsetID == "" {
		var err error
		if message.CharsetID, err = getCharsetID(message.Text); err != nil {
			return mvc.Response{
				Code: iris.StatusBadRequest,
				Text: "字符集保存失败",
			}
		}
	}
	if huffmanTree, err := ReadHuffmanTree(message.CharsetID); err == nil {
		// 获取霍夫曼编码
		huffmanCode := huffmanTree.GetHuffmanCode()

		// 编码原文
		res := strings.Builder{}
		for _, c := range message.Text {
			if code, ok := huffmanCode[string(c)]; ok {
				res.WriteString(code)
			} else {
				fmt.Println("原文与字符集不匹配")
				return mvc.Response{
					Code: iris.StatusBadRequest,
					Text: "原文与字符集不匹配",
				}
			}
		}
		// 编码终止符
		res.WriteString(huffmanCode["\u0000"])
		// 补齐字节
		if res.Len()%8 != 0 {
			diff := 8 - res.Len()%8
			for i := 0; i < diff; i++ {
				res.WriteString("0")
			}
		}
		fmt.Println("霍夫曼编码结果:", res.String())

		// 转为十六进制
		fmt.Println("十六进制结果为:", strings.ToUpper(bin2hex(res.String())))
		if result, err := json.Marshal(Message{
			Text:      strings.ToUpper(bin2hex(res.String())),
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

func bin2hex(bin string) string {
	tmp := strings.Builder{}
	res := strings.Builder{}
	for _, c := range bin {
		tmp.WriteString(string(c))
		if tmp.Len() == 4 {
			base, _ := strconv.ParseUint(tmp.String(), 2, 0)
			res.WriteString(strconv.FormatUint(base, 16))
			tmp.Reset()
		}
	}
	return res.String()
}

// getCharsetID 根据原文生成霍夫曼树
func getCharsetID(text string) (string, error) {
	charset := make(map[string]int)
	for _, c := range text {
		charset[string(c)]++
	}

	// 生成霍夫曼树
	charset["\u0000"] = 1 // 终止符
	huffmanTree := BuildHuffmanTree(charset)
	if err := huffmanTree.SaveHuffmanTree(); err != nil {
		fmt.Println("霍夫曼树保存失败")
		return "", err
	}
	fmt.Println("生成的霍夫曼树id:", huffmanTree.id)
	return huffmanTree.id, nil
}
