package xorm

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewClient() (client *gorm.DB, err error) {
	ch := make(chan *gorm.DB)
	go connDB(ch, client, "root:123456@tcp(localhost:3307)/bots?charset=utf8&parseTime=True&loc=Local")
	t := time.NewTicker(time.Second * 15)
	select {
	case v := <-ch:
		if v == nil {
			log.Fatal("orm conn error")
		}
		client = v
	case <-t.C:
		log.Fatal("orm conn timeout")
	}
	if err != nil || client == nil {
		err = errors.New("orm init failed")
		return
	}
	log.Println("orm success")
	return
}

func connDB(ch chan *gorm.DB, client *gorm.DB, str string) {
	client, err := gorm.Open("mysql", str)
	if err != nil {
		ch <- nil
	}
	client.LogMode(true)
	ch <- client
}
