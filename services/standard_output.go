package services

import (
	"fmt"
	"time"
)

type ImgStore interface {
	SaveImg(url string)
	GetUrl() string
}
type StandardOutputStore struct {
	t        *time.Time
	context  []byte
	imgUrl   string
	ImgSaver ImgStore
}

var DefaultStore = &StandardOutputStore{}

type DefaultImgStore struct {
	imgUrl string
}

func (s DefaultImgStore) SaveImg(url string) {
	s.imgUrl = url
}
func (s DefaultImgStore) GetUrl() string {
	return s.imgUrl
}

var DefaultImgSaverHandler = DefaultImgStore{}

func (s *StandardOutputStore) SaveContext(timestamp time.Time, context []byte, imgUrl string) {
	if s.ImgSaver == nil {
		s.ImgSaver = DefaultImgSaverHandler
	}
	s.ImgSaver.SaveImg(imgUrl)
	str := fmt.Sprintf("%v %v \n %v", timestamp.Format("2006-01-02 15:04"), string(context), s.ImgSaver.GetUrl())
	fmt.Println(str)
}
