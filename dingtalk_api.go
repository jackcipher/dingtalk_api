package dingtalk_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jackcipher/dingtalk_api/utils"
	"github.com/jackcipher/quickrequest"
	"log"
	"time"
)



var currentTimestamp int64
var webhook string
var sign string


type DingtalkConfig struct {
	Token string
	Secret string
}



func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func New(token, secret string) *DingtalkConfig {
	return &DingtalkConfig{
		Token:  token,
		Secret: secret,
	}
}

func (p *DingtalkConfig)reloadWebhook() {
	currentTimestamp = time.Now().Unix()*1000
	sign = utils.UrlEncode(hmacSha256(fmt.Sprintf("%d\n%s", currentTimestamp, p.Secret), p.Secret))
	webhook = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", p.Token, currentTimestamp, sign)
}

func (p *DingtalkConfig)SendMarkdown(title, message string, atMobiles []string, isAtAll bool) {
	var jsonByte []byte
	var err error
	p.reloadWebhook()
	result := utils.FormatMarkDownMessage(title, message, isAtAll, atMobiles)
	if jsonByte,err = json.Marshal(result); err!=nil {
		log.Fatalln(err)
	}
	quickrequest.PostJson(webhook, jsonByte, map[string]string{})
}
