package utils

import (
	"fmt"
	"github.com/jackcipher/dingtalk_api/structures"
)

func FormatMarkDownMessage(title, message string, isAtAll bool, atMobiles []string, manualAt bool) *structures.MarkdownMessage {
	atStr := ""
	if !manualAt {
		for _,v := range atMobiles {
			atStr += fmt.Sprintf("@%s", v)
		}
	}
	// 普通的at放到消息底部
	message = fmt.Sprintf("%s\n\n***\n\n%s\n\n", message, atStr)
	return &structures.MarkdownMessage{
		Msgtype:  "markdown",
		Markdown: structures.MarkdownRow{
			Title: title,
			Text: message,
		},
		At:	structures.AtRow{
			AtMobiles: atMobiles,
			IsAtAll: isAtAll,
		},
	}
}