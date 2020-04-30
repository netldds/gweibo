package gweibo

type Parser interface {
	ParseBody() map[string]interface{}
}

type HtmlParse struct {
}

func NewHtmlParse() *HtmlParse {
	return &HtmlParse{}
}
