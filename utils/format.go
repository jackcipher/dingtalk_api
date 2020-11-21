package utils

import (
	"fmt"
	"github.com/jackcipher/dingtalk_api/structures"
	"regexp"
)

func getMobilesFromContent(content string, atMobiles []string) ([]string, []string) {
	r,_ := regexp.Compile(`@(\d{11})`)
	result := r.FindAllStringSubmatch(content, -1)
	var parsedMobiles []string
	var extendMobiles []string
	for _, atMobile := range atMobiles {
		flag := true
		for _,v := range result {
			mobile := v[1]
			if mobile == atMobile {
				flag = false
			}
		}
		if flag {
			extendMobiles = append(extendMobiles, atMobile)
		}
	}
	for _,v := range result {
		mobile := v[1]
		parsedMobiles = append(parsedMobiles, mobile)
	}
	return parsedMobiles, extendMobiles
}

func FormatMarkDownMessage(title, message string, isAtAll bool, atMobiles []string) *structures.MarkdownMessage {
	atStr := ""
	parsedMobiles, extendMobiles := getMobilesFromContent(message, atMobiles)
	if len(extendMobiles) > 0 {
		for _,v := range extendMobiles {
			atStr += fmt.Sprintf("@%s", v)
		}
	}
	// 普通的at放到消息底部
	if atStr != "" {
		message = fmt.Sprintf("%s\n\n***\n\n%s\n\n", message, atStr)
	}
	for _,v := range parsedMobiles {
		atMobiles = append(atMobiles, v)
	}
	return &structures.MarkdownMessage{
		Msgtype:  "markdown",
		Markdown: structures.MarkdownRow{
			Title: title,
			Text: message,
		},
		At: structures.AtRow{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	}
}