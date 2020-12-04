package main

import (
	"fmt"
	"github.com/jackcipher/dingtalk_api"
)

func main() {
	//dingtalk := dingtalk_api.NewGroupBot("", "")
	//rst := dingtalk.SendMarkdown("unitTest", "test", []string{}, false)
	//fmt.Println(rst)
	bot := dingtalk_api.NewWorkBot("", 0)
	err := bot.SendWorkNoticeToPersons([]int{0}, "test", "message")
	fmt.Println(err)
}
