package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	addre := "//wx4.sinaimg.cn/orj360/006yududgy1g5iabhjai3j310c0rkgob.jpg"
	addre = "https:" + addre
	addre = strings.Replace(addre, "orj360", "mw690", 1)
	resp, err := http.Get(addre)
	if err != nil {
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	picName := filepath.Base(addre)
	os.Mkdir("dir", 0775)
	p := path.Join("dir", picName)
	err = ioutil.WriteFile(p, b, 0775)
	if err != nil {
		fmt.Println(err)
	}
}
