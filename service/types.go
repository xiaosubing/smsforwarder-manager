package service

// LoginRequest 接受登录参数结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 登录返回
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// sendMessageData
type sendMessageData struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
	Num     string `json:"num"`
}

// ForwarderMessage
type ForwarderMessageData struct {
	Phone   string `json:"phone"`
	Number  string `json:"number"`
	Content string `json:"message"`
	Sign    string `json:"sign"`
}

// 定义消息模型
type Message struct {
	ID      uint   `gorm:"primaryKey"`
	Phone   string `gorm:"size:20"`
	Number  string `gorm:"size:128"`
	Content string `gorm:"type:text"`
	Code    string `gorm:"size:10"`
}

// 发短信
type SendMessage struct {
	Data      SendMessageData `json:"data"`
	Timestamp int64           `json:"timestamp"`
	Sign      string          `json:"sign"`
}
type SendMessageData struct {
	SimSlot      int    `json:"sim_slot"`
	PhoneNumbers string `json:"phone_numbers"`
	MsgContent   string `json:"msg_content"`
}

type getMessageRequest struct {
	Phone string `json:"phone"`
}
