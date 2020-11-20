package main

import "github.com/jackcipher/dingtalk_api"

func main() {
	dingtalk := dingtalk_api.New("token", "secret")
	dingtalk.SendMarkdown("test", "hi", []string{"your_phone_number"}, false)
}
