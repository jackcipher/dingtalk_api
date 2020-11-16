package structures

type MarkdownRow struct {
	Title string	`json:"title"`
	Text  string	`json:"text"`
}

type AtRow struct {
	AtMobiles []string	`json:"atMobiles"`
	IsAtAll bool		`json:"isAtAll"`
}

type MarkdownMessage struct {
	Msgtype string	`json:"msgtype"`
	Markdown MarkdownRow	`json:"markdown"`
	At AtRow	`json:"at"`
}
