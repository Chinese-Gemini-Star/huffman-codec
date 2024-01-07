package huffmanCode

// Message 前端调用接口用消息
type Message struct {
	// Id 主键
	Id uint `gorm:"primaryKey,auto_increment"`
	// Text 文本
	Text string `json:"Text"`
	// CharsetID 字符集ID
	CharsetID string `json:"CharsetID"`
	// Username 用户名
	Username string `json:"Username"`
}
