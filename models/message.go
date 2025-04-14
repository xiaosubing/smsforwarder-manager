package models

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
