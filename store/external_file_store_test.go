package store

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestPxyPut(t *testing.T) {
	data, _ := ioutil.ReadFile("/home/me/Figure_1.png")
	log.Println(len(data))

	r := NewClosingBuffer(data)
	extfs := GetExtFS()
	uname, err := extfs.PutFileByReaderVimcn(r, "Figure_1.png", true)
	log.Println(uname, err)
}
