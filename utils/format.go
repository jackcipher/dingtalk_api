package utils

import (
	"fmt"
	"github.com/jackcipher/dingtalk_api/structures"
)

func FormatMarkDownMessage(title, message string, isAtAll bool, atMobiles []string) *structures.MarkdownMessage {
	atStr := ""
	for _,v := range atMobiles {
		atStr += fmt.Sprintf("@%s", v)
	}
	message = atStr + message
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