package services

import (
	"fmt"
	"log"
	"time"
)

type Store interface {
	SaveContext(timestamp time.Time, context []byte, img string)
}
type ImgSaver interface {
	SaveImg(url string)
	GetUrl() string
}
type StandardOutput struct {
	t        *time.Time
	context  []byte
	imgUrl   string
	ImgSaver ImgSaver
}

type DefautlImgSaver struct {
	imgUrl string
}

func (s DefautlImgSaver) SaveImg(url string) {
	s.imgUrl = url
}
func (s DefautlImgSaver) GetUrl() string {
	return s.imgUrl
}

var DefaultImgSaverHandler = DefautlImgSaver{}

func (s *StandardOutput) SaveContext(timestamp time.Time, context []byte, imgUrl string) {
	if s.ImgSaver == nil {
		s.ImgSaver = DefaultImgSaverHandler
	}
	s.ImgSaver.SaveImg(imgUrl)
	str := fmt.Sprintf("%v %v \n %v", timestamp, string(context), s.ImgSaver.GetUrl())
	log.Println(str)
}
