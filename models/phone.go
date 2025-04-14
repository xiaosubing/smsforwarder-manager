package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

func InsertPhoneData(data string, ip string) string {
	var phoneData PhoneBase
	err := json.Unmarshal([]byte(data), &phoneData)
	if err != nil {
		return ""
	}

	e := insertPhoneDataSub(phoneData.Data.SimInfoList.Num0.Number, ip, phoneData.Data.ExtraSim1)
	if e != "ok" {
		return e
	}
	e = insertPhoneDataSub(phoneData.Data.SimInfoList.Num1.Number, ip, phoneData.Data.ExtraSim2)
	if e != "ok" {
		return e
	}
	return "ok"
}

func insertPhoneDataSub(phoneNumber, ip, alias string) string {

	if phoneNumber == "" {
		return "phone is null"
	}
	phoneNumber = strings.Replace(phoneNumber, "+86", "", -1)

	keyword := fmt.Sprintf("phone = \"%s\"", phoneNumber)

	phones := QueryPhoneData(keyword, 0)
	fmt.Println(phones)

	if len(phones) != 0 {
		fmt.Println()
		fmt.Println("已经存在，无需重新插入")
		return "phone exist"
	} else {
		fmt.Println()
		fmt.Println("检查后仍需要插入")
		fmt.Println(keyword)
		fmt.Println(phoneNumber)
		fmt.Println("======================")
	}

	var p Phone

	p.IP = ip
	p.Phone = strings.Replace(phoneNumber, "+86", "", -1)
	p.Alias = alias

	// save
	err := InsertData(&p)
	if err != nil {
		return "insert phone data error"
	}
	return "ok"
}

func QueryPhoneData(keyword string, size int) []Phone {
	var p []Phone
	var param QueryParams
	if keyword != "" {
		param.Keyword = keyword
	}
	if size != 0 {
		param.PageSize = size
	}

	err := QueryData(&p, param)
	if err != nil {
		return nil
	}

	return p

}

func QueryPhoneDataSign(keyword string, size int) *Phone {
	var p Phone
	var param QueryParams
	if keyword != "" {
		param.Keyword = keyword
	}
	if size != 0 {
		param.PageSize = size
	}

	err := QueryData(&p, param)
	if err != nil {
		return nil
	}

	return &p

}
