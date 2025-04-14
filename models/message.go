package models

import "fmt"

func SaveMessage(phone, number, content, code string) error {
	var m Message
	m.Phone = phone
	m.Number = number
	m.Content = content
	m.Code = code

	err := InsertData(&m)
	if err != nil {
		return err
	}
	return err
}

func GetMessageCode(param QueryParams) *Message {
	var m Message

	err := QueryData(&m, param)
	if err != nil {
		fmt.Println("执行出错拉！ ")
		fmt.Println(err)
		return nil
	}

	return &m
}
