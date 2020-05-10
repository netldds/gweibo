package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type ImgStore interface {
	SaveImg(url string)
	GetReferrer() string
}
type StandardOutputStore struct {
	t        *time.Time
	context  []byte
	imgUrl   string
	ImgSaver ImgStore
}

var DefaultStore = &StandardOutputStore{}

const (
	PREVIEW_SIZE = "orj360" //小图
	FULL_SIZE    = "mw690"  //大图
	PIC_DIR      = "PICS"
)

type DefaultImgStore struct {
	imgUrl  string
	ImgName string
}

func (s *DefaultImgStore) SaveImg(url string) {
	if url == "" {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in parsing error is %v: \n", err)
		}
	}()
	s.imgUrl = url
	s.imgUrl = "https:"
	s.imgUrl = strings.Replace(s.imgUrl, PREVIEW_SIZE, FULL_SIZE, 1)
	resp, err := http.Get(s.imgUrl)
	if err != nil {
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	s.ImgName = filepath.Base(s.imgUrl)
	os.Mkdir(PIC_DIR, 0775)
	p := path.Join(PIC_DIR, s.ImgName)
	err = ioutil.WriteFile(p, b, 0775)
	if err != nil {
		fmt.Println(err)
	}
}
func (s *DefaultImgStore) GetReferrer() string {
	return s.imgUrl
}

var DefaultImgSaverHandler = &DefaultImgStore{}

func (s *StandardOutputStore) SaveContext(timestamp time.Time, context []byte, imgUrl string) {
	if s.ImgSaver == nil {
		s.ImgSaver = DefaultImgSaverHandler
	}
	s.ImgSaver.SaveImg(imgUrl)
	str := fmt.Sprintf("%v %v \n %v", timestamp.Format("2006-01-02 15:04"), string(context), s.ImgSaver.GetReferrer())
	fmt.Println(str)
}
