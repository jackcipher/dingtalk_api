package main

import (
	"github.com/jackcipher/dingtalk_api"
)

func main() {
	dingtalk := dingtalk_api.New("", "")
	dingtalk.SendMarkdown("unitTest", "", []string{}, false)
}
