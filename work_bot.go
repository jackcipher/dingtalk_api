package dingtalk_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackcipher/dingtalk_api/structures"
	"github.com/jackcipher/quickrequest"
	"net/http"
	"strings"
)

type WorkBotConfig struct {
	AccessToken string
	AgentId int
}

type WorkNoticePersonsParams struct {
	AgentId int `json:"agent_id"`
	UseridList string `json:"userid_list"`
	Msg interface{} `json:"msg"`
}

type WorkNoticeDeptsParams struct {
	AgentId int `json:"agent_id"`
	DeptIdList string `json:"dept_id_list"`
	Msg structures.NoticeMarkdownMessage `json:"msg"`
}

func NewWorkBot(accessToken string, agentId int) *WorkBotConfig {
	return &WorkBotConfig{
		AccessToken: accessToken,
		AgentId:     agentId,
	}
}

func (p *WorkBotConfig) getWorkNoticeUrl() string {
	return fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", p.AccessToken)
}

func (p *WorkBotConfig)formatWorkNoticeRequestParams(userIdList []string, title, message string, btns []structures.DActionCardForNoticeButton) (error, []byte) {
	if len(userIdList) ==0  {
		return errors.New("用户ID不能为空"), nil
	}
	message = strings.Trim(message, "")
	if len(message) == 0 {
		return errors.New("消息数据不能为空"), nil
	}
	strUserIdList := strings.Join(userIdList, ",")
	var config WorkNoticePersonsParams
	if len(btns) > 0 {
		params := structures.DActionCardForNoticeParams {
			Title:   title,
			Message: message,
			Buttons: btns,
		}
		actionCard := params.Format()
		config = WorkNoticePersonsParams{
			AgentId:    p.AgentId,
			UseridList: strUserIdList,
			Msg:        actionCard,
		}
	} else {
		config = WorkNoticePersonsParams{
			AgentId:    p.AgentId,
			UseridList: strUserIdList,
			Msg: structures.NoticeMarkdownMessage{
				Msgtype:  "markdown",
				Markdown: structures.MarkdownRow{
					Title: title,
					Text: message,
				},
			},
		}
	}
	var err error
	var byteJson []byte
	if byteJson,err = json.Marshal(config); err != nil {
		return err, nil
	}
	return nil, byteJson
}

func (p *WorkBotConfig) SendWorkNoticeToPersons(userIdList []string, title, message string, btns []structures.DActionCardForNoticeButton) error {
	err, byteJson := p.formatWorkNoticeRequestParams(userIdList, title, message, btns)
	if err != nil {
		return err
	}
	url := p.getWorkNoticeUrl()
	fmt.Println(url, string(byteJson))
	rawResult, statusCode := quickrequest.PostJson(url, byteJson, nil)
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("网络错误，状态码:%d", statusCode))
	}
	var response = &structures.DingtalkResponse{}
	if err := json.Unmarshal(rawResult, response); err!=nil {
		return errors.New("JSON解析失败")
	}
	if response.ErrCode != 0 {
		return errors.New(fmt.Sprintf("推送工作通知失败, 钉钉API返回:%s", response.ErrMsg))
	}
	return nil
}