package dingtalk_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackcipher/dingtalk_api/structures"
	"github.com/jackcipher/dingtalk_api/utils"
	"github.com/jackcipher/quickrequest"
	"net/http"
	"time"
)

var currentTimestamp int64
var webhook string
var sign string

type GroupBotConfig struct {
	Token string
	Secret string
}


func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func NewGroupBot(token, secret string) *GroupBotConfig {
	return &GroupBotConfig{
		Token:  token,
		Secret: secret,
	}
}

func (p *GroupBotConfig)reloadWebhook() {
	currentTimestamp = time.Now().Unix()*1000
	sign = utils.UrlEncode(hmacSha256(fmt.Sprintf("%d\n%s", currentTimestamp, p.Secret), p.Secret))
	webhook = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", p.Token, currentTimestamp, sign)
}

func (p *GroupBotConfig)SendMarkdown(title, message string, atMobiles []string, isAtAll bool) error {
	var jsonByte []byte
	var err error
	p.reloadWebhook()
	result := utils.FormatMarkDownMessage(title, message, isAtAll, atMobiles)
	if jsonByte,err = json.Marshal(result); err!=nil {
		return errors.New("消息格式化失败")
	}
	rawResult,statusCode := quickrequest.PostJson(webhook, jsonByte, map[string]string{})
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("网络错误，状态码:%d", statusCode))
	}
	var response = &structures.DingtalkResponse{}
	if err := json.Unmarshal(rawResult, response); err!=nil {
		return errors.New("JSON解析失败")
	}
	if response.ErrCode != 0 {
		return errors.New(fmt.Sprintf("消息推送失败, 钉钉API返回:%s", response.ErrMsg))
	}
	return nil
}
