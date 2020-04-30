package services

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
)

type ResponseHandler interface {
	Value
	Parse()
}
type Value interface {
	GetValue() string
}
type Reserver interface {
	Save()
}
type HomePageResponse struct {
}

func (s *HomePageResponse) GetValue() string {

}
func (s *HomePageResponse) Parse() {
	htmlR := bytes.NewBufferString(content)
	body, _ := html.Parse(htmlR)
	fmt.Println(body.Type)
}

var content = `<!DOCTYPE html>
<html>
    <head>
        <title>
            Title of the document
        </title>
    </head>
    <body>
        body content 
        <p>more content</p>
    </body>
</html> `
