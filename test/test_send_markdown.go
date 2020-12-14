package main

import (
	"fmt"
	"github.com/jackcipher/dingtalk_api"
	"github.com/jackcipher/dingtalk_api/structures"
)

func main() {
	//dingtalk := dingtalk_api.NewGroupBot("", "")
	//rst := dingtalk.SendMarkdown("unitTest", "", []string{""}, false, []structures.DActionCardButton{
	//	{
	//		Title:     "left",
	//		ActionURL: "https://www.baidu.com",
	//	},
	//	{
	//		Title: "right",
	//		ActionURL: "http://www.baidu.com",
	//	},
	//})
	//fmt.Println(rst)
	bot := dingtalk_api.NewWorkBot("", 0)
	err := bot.SendWorkNoticeToPersons([]string{""}, "test", "message", []structures.DActionCardForNoticeButton{
		{
			Title: "test",
			ActionURL: "https://www.baidu.com",
		},
	})
	fmt.Println(err)
}
