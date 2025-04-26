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

// changepwd
type ChangePassword struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
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

// forwarder 消息模型

type AndroidMessage struct {
	Timestamp int64                `json:"timestamp"`
	Code      int                  `json:"code"`
	Msg       string               `json:"msg"`
	Data      []AndroidMessageData `json:"data"`
}
type AndroidMessageData struct {
	Content     string `json:"content"`
	Date        int64  `json:"date"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	SimID       int    `json:"sim_id"`
	SubID       int    `json:"sub_id"`
	Type        int    `json:"type"`
	TypeImageID int    `json:"typeImageId"`
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
