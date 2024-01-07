package huffmanCode

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liyue201/gostl/ds/priorityqueue"
	"io"
	"os"
	"sort"
	"strconv"
)

type HuffmanTree struct {
	id     string
	Graph  *cgraph.Graph // 伪装饰者模式
	Weight int           `json:"weight"`
	Char   string        `json:"char"`
	Left   *HuffmanTree  `json:"left"`
	Right  *HuffmanTree  `json:"right"`
	deep   int
}

// BuildHuffmanTree 构建霍夫曼树
func BuildHuffmanTree(charset map[string]int) *HuffmanTree {
	// 利用小顶堆快速取出当前两个权值最小的元素
	pq := priorityqueue.New[*HuffmanTree](func(a, b *HuffmanTree) int {
		return a.Weight - b.Weight
	}, priorityqueue.WithGoroutineSafe())

	// 将字符集作为叶子节点入队
	keys := make([]string, 0)
	for key, _ := range charset {
		keys = append(keys, key)
	}
	sort.Strings(keys) // 确保每次生成的霍夫曼树一致
	for _, key := range keys {
		node := HuffmanTree{
			Weight: charset[key],
			Char:   key,
			deep:   0,
		}
		pq.Push(&node)
	}

	// 构建霍夫曼树
	for pq.Size() > 1 {
		left := pq.Top()
		pq.Pop()
		right := pq.Top()
		pq.Pop()
		node := HuffmanTree{
			Weight: left.Weight + right.Weight,
			Char:   "",
			Left:   left,
			Right:  right,
			deep:   min(left.deep, right.deep) + 1,
		}
		pq.Push(&node)
	}

	return pq.Top()
}

//// ReadHuffmanTree 从本地读取已存在的霍夫曼树
//func ReadHuffmanTree(id string) (*HuffmanTree, error) {
//	if _, err := os.Stat("./tmp/huffmanTree/" + id + ".tree"); err == nil {
//		if f, err := os.OpenFile("./tmp/huffmanTree/"+id+".tree", os.O_RDONLY, 0600); err == nil {
//			defer f.Close()
//			if content, err := io.ReadAll(f); err == nil {
//				huffmanTree := new(HuffmanTree)
//				if err := json.Unmarshal(content, &huffmanTree); err == nil {
//					return huffmanTree, nil
//				} else {
//					return nil, err
//				}
//			} else {
//				return nil, err
//			}
//		} else {
//			return nil, err
//		}
//	} else {
//		return nil, err
//	}
//}

// ReadHuffmanTree 从本地读取已存在的霍夫曼树
func ReadHuffmanTree(id string) (*HuffmanTree, error) {
	if _, err := os.Stat("./tmp/huffmanTree/" + id + ".tree"); err == nil {
		if f, err := os.OpenFile("./tmp/huffmanTree/"+id+".tree", os.O_RDONLY, 0600); err == nil {
			defer f.Close()
			if content, err := io.ReadAll(f); err == nil {
				if graph, err := graphviz.ParseBytes(content); err == nil {
					// 伪装饰者模式,封装成自己的类
					return &HuffmanTree{
						Graph: graph,
						id:    id,
					}, nil
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

//// getTreeJson 获取霍夫曼树的序列化结果
//func (tree *HuffmanTree) getTreeJson() ([]byte, error) {
//	fmt.Println("序列化霍夫曼树")
//	// 序列化霍夫曼树
//	if treeJson, err := json.Marshal(*tree); err == nil {
//		return treeJson, nil
//	} else {
//		return nil, err
//	}
//}

// getTreeDot 获取霍夫曼树的序列化结果
func (tree *HuffmanTree) getTreeDot() ([]byte, error) {
	fmt.Println("序列化霍夫曼树")
	g := graphviz.New()
	if graph, err := g.Graph(); err == nil {
		// 遍历霍夫曼树
		var dfs func(root *HuffmanTree) *cgraph.Node
		id := 0
		dfs = func(root *HuffmanTree) *cgraph.Node {
			// 叶子节点
			if root.Char != "" {
				n, _ := graph.CreateNode(strconv.Itoa(id))
				id++
				if root.Char != "\u0000" {
					n.SetLabel(fmt.Sprintf(`'%s' %d`, root.Char, root.Weight))
				} else {
					n.SetLabel(fmt.Sprintf("[EOF] %d", root.Weight))
				}
				return n
			} else {
				n, _ := graph.CreateNode(strconv.Itoa(id))
				id++
				n.SetLabel(strconv.Itoa(root.Weight))
				// 左
				left := dfs(root.Left)
				graph.CreateEdge("", n, left)
				// 右
				right := dfs(root.Right)
				graph.CreateEdge("", n, right)
				return n
			}
		}
		dfs(tree)
		tree.Graph = graph
		var buf bytes.Buffer
		// 序列化
		if err := g.Render(graph, "dot", &buf); err == nil {
			return buf.Bytes(), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// GetHuffmanCode 获取霍夫曼编码
func (tree *HuffmanTree) GetHuffmanCode() map[string]string {
	// 如果本身已经存在本地,则直接反序列化已经编码的结果
	if tree.id != "" {
		if _, err := os.Stat("./tmp/huffmanTree/" + tree.id + ".code"); err == nil {
			if f, err := os.OpenFile("./tmp/huffmanTree/"+tree.id+".code", os.O_RDONLY, 0600); err == nil {
				defer f.Close()
				if content, err := io.ReadAll(f); err == nil {
					huffmanCode := make(map[string]string)
					if err := json.Unmarshal(content, &huffmanCode); err == nil {
						return huffmanCode
					} else {
						fmt.Println("反序列化霍夫曼编码失败")
					}
				} else {
					fmt.Println("读取霍夫曼编码文件失败")
				}
			} else {
				fmt.Println("打开霍夫曼编码文件失败")
			}
		} else if !os.IsNotExist(err) {
			fmt.Println("读取霍夫曼树文件失败")
		}
	}

	// 遍历霍夫曼树用闭包递归函数,左0右1编码
	huffmanCode := make(map[string]string)
	var getCharCode func(root *HuffmanTree, code *[]byte)
	getCharCode = func(root *HuffmanTree, code *[]byte) {
		if root.Char != "" {
			huffmanCode[root.Char] = string(*code)
		} else {
			// 左0
			left := append(*code, '0')
			getCharCode(root.Left, &left)
			// 右1
			right := append(*code, '1')
			getCharCode(root.Right, &right)
		}
	}

	code := make([]byte, 0, tree.deep+1)
	if tree.deep == 0 {
		// 单节点,特判
		code = append(code, '0')
		huffmanCode[tree.Char] = string(code)
	} else {
		// 遍历霍夫曼树,获取字符集
		getCharCode(tree, &code)
	}
	return huffmanCode
}

// SaveHuffmanTree 序列化霍夫曼树以及编码到本地
func (tree *HuffmanTree) SaveHuffmanTree() error {
	var name string
	// 序列化霍夫曼树
	if treeDot, err := tree.getTreeDot(); err == nil {
		// 转化为md5值,以此命名文件名,防止重复
		name = getMd5(treeDot)
		if f, err := os.Create("./tmp/huffmanTree/" + name + ".tree"); err == nil {
			defer f.Close()
			if _, err := f.Write(treeDot); err != nil {
				fmt.Println("写入霍夫曼树文件失败")
				return err
			}
			tree.id = name
		} else if os.IsExist(err) {
			fmt.Println("文件已经存在")
		} else {
			fmt.Println("创建霍夫曼树文件失败")
			return err
		}
	} else {
		fmt.Println("序列化霍夫曼树失败")
		return err
	}

	// 序列化霍夫曼编码
	if codeJson, err := json.Marshal(tree.GetHuffmanCode()); err == nil {
		if f, err := os.Create("./tmp/huffmanTree/" + name + ".code"); err == nil {
			defer f.Close()
			if _, err := f.Write(codeJson); err != nil {
				fmt.Println("写入霍夫曼编码失败")
				return err
			}
		} else {
			fmt.Println("创建霍夫曼编码文件失败")
			return err
		}
	} else if os.IsExist(err) {
		fmt.Println("文件已经存在")
	} else {
		fmt.Println("序列化霍夫曼编码失败")
		return err
	}
	return nil
}

// PostCharset 前端上传字符集
func (c *Controller) PostCharset() mvc.Result {
	charset := make(map[string]int)
	if err := c.Ctx.ReadJSON(&charset); err != nil {
		fmt.Println("无法解析前端发送的字符集,json格式不正确")
		return mvc.Response{
			Code: iris.StatusBadRequest,
			Text: "无法解析字符集,json格式不正确",
		}
	}
	fmt.Println("获取到字符集数据:", charset)

	// 检查字符集数据是否合法
	for c, _ := range charset {
		// 字符长度超过1
		if len(c) != 1 {
			fmt.Println("字符长度超过1")
			return mvc.Response{
				Code: iris.StatusBadRequest,
				Text: "字符长度只允许为1",
			}
		}
	}

	// 生成霍夫曼树
	charset["\u0000"] = 1 // 终止符
	huffmanTree := BuildHuffmanTree(charset)
	if err := huffmanTree.SaveHuffmanTree(); err != nil {
		fmt.Println("霍夫曼树保存失败")
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Text: "霍夫曼树保存失败",
		}
	}
	return mvc.Response{
		Code: iris.StatusOK,
		Text: huffmanTree.id,
	}
}

// GetTreePngBy 获取霍夫曼树图形
func (c *Controller) GetTreePngBy(id string) mvc.Result {
	// 从本地反序列化霍夫曼树
	if huffmanTree, err := ReadHuffmanTree(id); err == nil {
		g := graphviz.New()
		var image bytes.Buffer
		// 写入字节缓冲流中
		if err := g.Render(huffmanTree.Graph, graphviz.PNG, &image); err == nil {
			fmt.Println("成功生成霍夫曼树", id, "的图形")
			return mvc.Response{
				Code:        iris.StatusOK,
				ContentType: "image/png",
				Content:     image.Bytes(),
			}
		} else {
			fmt.Println("霍夫曼树图形生成失败")
			return mvc.Response{
				Code: iris.StatusInternalServerError,
				Text: "霍夫曼树图形生成失败",
			}
		}
	} else {
		fmt.Println("反序列化霍夫曼树失败")
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Text: "反序列化霍夫曼树失败",
		}
	}
}

// getMd5 获取md5码值
func getMd5(str []byte) string {
	hash := md5.Sum(str)
	return hex.EncodeToString(hash[:])
}
