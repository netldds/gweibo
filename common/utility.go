package common

import "log"

type Parser interface {
	ParseBody() map[string]interface{}
}

type HtmlParse struct {
}

func NewHtmlParse() *HtmlParse {
	return &HtmlParse{}
}

type DefaultSaver struct {
}

func (s *DefaultSaver) Save(ct string) {
	log.Println(ct)
}
