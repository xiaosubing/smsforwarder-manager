package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Phone   string `gorm:"size:20"`
	Number  string `gorm:"size:128"`
	Content string `gorm:"type:text"`
	Code    string `gorm:"size:10"`
	SmsType string `gorm:"size:10"`
	//CreatedAt string `gorm:"size:20"`
}

type User struct {
	gorm.Model
	//ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(20);" json:"username"`
	Password string `gorm:"type:varchar(20);" json:"password"`
	//Phone    string `gorm:"type:varchar(20);" json:"phone"`
}

type Phone struct {
	gorm.Model
	Phone string `gorm:"type:varchar(20);" json:"phone"`
	IP    string `gorm:"type:varchar(20);" json:"ip"`
	Alias string `gorm:"type:varchar(20);" json:"alias"`
}

// 基础信息
type PhoneBase struct {
	Timestamp int64         `json:"timestamp"`
	Code      int           `json:"code"`
	Msg       string        `json:"msg"`
	Data      PhoneBaseData `json:"data"`
}
type Num0 struct {
	CarrierName    string `json:"carrier_name"`
	CountryIso     string `json:"country_iso"`
	IccID          string `json:"icc_id"`
	Number         string `json:"number"`
	SimSlotIndex   int    `json:"sim_slot_index"`
	SubscriptionID int    `json:"subscription_id"`
}
type Num1 struct {
	CarrierName    string `json:"carrier_name"`
	CountryIso     string `json:"country_iso"`
	IccID          string `json:"icc_id"`
	Number         string `json:"number"`
	SimSlotIndex   int    `json:"sim_slot_index"`
	SubscriptionID int    `json:"subscription_id"`
}
type SimInfoList struct {
	Num0 Num0 `json:"0"`
	Num1 Num1 `json:"1"`
}
type PhoneBaseData struct {
	EnableAPIBatteryQuery bool        `json:"enable_api_battery_query"`
	EnableAPICallQuery    bool        `json:"enable_api_call_query"`
	EnableAPIClone        bool        `json:"enable_api_clone"`
	EnableAPIContactQuery bool        `json:"enable_api_contact_query"`
	EnableAPISmsQuery     bool        `json:"enable_api_sms_query"`
	EnableAPISmsSend      bool        `json:"enable_api_sms_send"`
	EnableAPIWol          bool        `json:"enable_api_wol"`
	ExtraDeviceMark       string      `json:"extra_device_mark"`
	ExtraSim1             string      `json:"extra_sim1"`
	ExtraSim2             string      `json:"extra_sim2"`
	SimInfoList           SimInfoList `json:"sim_info_list"`
}

// QueryParams 查询参数结构体
//
//	type PhoneQueryParams struct {
//		Phone string `json:"phone"`
//		IP    string `json:"ip"`
//		Alias string `json:"alias"`
//	}
//
// QueryParams 查询参数结构体
type QueryParams struct {
	Type     int    `json:"type"`      // 类型
	PageNum  int    `json:"page_num"`  // 页码
	PageSize int    `json:"page_size"` // 每页数量
	Keyword  string `json:"keyword"`   // 关键字
}

//type UserQueryParams struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

func (table *User) TableName() string {
	return "user"
}

func (table *Message) TableName() string {
	return "message"
}

func (table *Phone) TableName() string {
	return "phone"
}
