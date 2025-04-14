package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"smsforwarder-manager/models"
	"strings"
)

func GetMessages(c *gin.Context) {

	recInfo := new(getMessageRequest)
	err := c.ShouldBind(recInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}

	fmt.Println("获取到的手机号：", recInfo.Phone)
	keyword := fmt.Sprintf("alias = \"%s\"", recInfo.Phone)
	phone := models.QueryPhoneDataSign(keyword, 0)
	fmt.Println(phone.IP)
	url := fmt.Sprintf("http://%s:5000/sms/query", phone.IP)

	payload := `{
  "data": {
    "type": 1,
    "page_num": 1,
    "page_size": 5,
    "keyword": ""
  },
  "timestamp": 1652590258638,
  "sign": ""
}`
	req := HttpPost(url, payload)

	var m []Message
	err = json.Unmarshal([]byte(req), &m)
	if err != nil {
		fmt.Println(err)
	}
	var messages []map[string]string
	for _, info := range m {

		message := map[string]string{
			"sender":  info.Number,
			"content": info.Content,
		}
		messages = append(messages, message)
	}
	c.JSON(http.StatusOK, messages)
}

func SendMessages(c *gin.Context) {
	sendMsgData := new(sendMessageData)
	err := c.ShouldBind(sendMsgData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "404",
		})
		fmt.Println(err)
		return
	}
	//fmt.Println("============================")
	//fmt.Println(sendMsgData.Phone)
	//fmt.Println(sendMsgData.Content)
	//fmt.Println(sendMsgData.Num)
	//fmt.Println("开始发短信")
	//fmt.Println("============================")
	keyword := fmt.Sprintf("alias = \"%s\"", sendMsgData.Num)
	phone := models.QueryPhoneDataSign(keyword, 0)
	fmt.Println(phone.IP)
	url := fmt.Sprintf("http://%s:5000/sms/send", phone.IP)
	payload := `{
  "data": {
    "sim_slot": 1,
    "phone_numbers": "[手机号]",
    "msg_content": "[短信内容]"
  },
  "timestamp": 1652590258638,
  "sign": ""
}`
	payload = strings.Replace(strings.Replace(payload, "[手机号]", sendMsgData.Phone, -1), "[短信内容]", sendMsgData.Content, -1)
	req := HttpPost(url, payload)
	if strings.Contains(req, "err") {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  404,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"req":  req,
		})
	}

}

// ForwarderMessage
func ForwarderMessage(c *gin.Context) {
	forwarderMsgData := new(ForwarderMessageData)
	err := c.ShouldBind(forwarderMsgData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}
	// sign 比对
	if forwarderMsgData.Sign != "梅干小小酥饼" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  404,
		})
		return
	}
	code := MessageCodeProcess(forwarderMsgData.Number)
	err = models.SaveMessage(forwarderMsgData.Phone, forwarderMsgData.Number, forwarderMsgData.Content, code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
	
}

func MessageCodeProcess(content string) string {
	re := regexp.MustCompile(`.*?(.{0,15})[随机|验证|登录|授权|动态|校验]码(.{0,10})`)
	match := re.FindAllString(content, -1)
	if len(match) == 0 {
		return "None"
	}
	re = regexp.MustCompile(`\d{4,6}\b`)
	code := re.FindAllString(match[0], -1)

	if len(code) == 0 {
		content = "验证" + strings.Replace(content, match[0], "", -1)
		//re = regexp.MustCompile(`.*?(.{0,15})[随机|验证|登录|授权|动态|校验]码(.{0,10})`)
		match = re.FindAllString(content, -1)
		if len(match) == 0 {
			return "None"
		}
		re = regexp.MustCompile(`\d{4,6}\b`)
		code = re.FindAllString(match[0], -1)
	}

	if len(code) != 0 {
		return code[0]
	}
	return "None"
}
