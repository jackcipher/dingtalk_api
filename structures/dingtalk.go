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

type DingtalkResponse struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

type NoticeMarkdownMessage struct {
	Msgtype string	`json:"msgtype"`
	Markdown MarkdownRow	`json:"markdown"`
}

const (
	DTypeButtonVertical   = "0"
	DTypeButtonHorizontal = "1"
)

const (
	DActionCardNameForGroup  = "actionCard"
	DActionCardNameForNotice = "action_card"
)

type DActionCardButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type DActionCardDetail struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	BtnOrientation string `json:"btnOrientation"`
	Buttons        []DActionCardButton `json:"btns"`
}

type DActionCard struct {
	Msgtype    string `json:"msgtype"`
	Detail DActionCardDetail `json:"actionCard"`
}

type DActionCardParams struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Buttons []DActionCardButton `json:"buttons"`
}

type DActionCardForNoticeButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"action_url"`
}

type DActionCardForNoticeDetail struct {
	Title          string `json:"title"`
	Markdown       string `json:"markdown"`
	BtnOrientation string `json:"btn_orientation"`
	BtnJSONList    []DActionCardForNoticeButton `json:"btn_json_list"`
}

type DActionCardForNotice struct {
	Msgtype string `json:"msgtype"`
	Detail DActionCardForNoticeDetail `json:"action_card"`
}

type DActionCardForNoticeParams struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Buttons []DActionCardForNoticeButton `json:"buttons"`
}

func (p *DActionCardParams) Format() *DActionCard  {
	detail := DActionCardDetail{
		Title:          p.Title,
		Text:           p.Message,
		BtnOrientation: DTypeButtonHorizontal,
		Buttons:        p.Buttons,
	}
	result := DActionCard {
		Msgtype: "actionCard",
		Detail:  detail,
	}
	return &result
}

func (p *DActionCardForNoticeParams) Format() *DActionCardForNotice {
	detail := DActionCardForNoticeDetail{
		Title:          p.Title,
		Markdown:       p.Message,
		BtnOrientation: DTypeButtonHorizontal,
		BtnJSONList:    p.Buttons,
	}
	result := DActionCardForNotice {
		Msgtype: "action_card",
		Detail:  detail,
	}
	return &result
}