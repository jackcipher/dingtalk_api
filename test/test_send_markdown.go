package main

import "github.com/jackcipher/dingtalk_api"

func main() {
	dingtalk := dingtalk_api.New("x", "x")
	dingtalk.SendMarkdown("test", "hi", []string{"x"}, true)
}
