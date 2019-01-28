package events

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/viile/rssbot/models"
	"github.com/viile/rssbot/pkg/xorm"
	"log"
	"net/http"
	"sort"
	"strings"
)

var (
	orm        *gorm.DB
)

type (
	Events struct{}
)

type Message struct {
	ToUserName string
	FromUserName string
	CreateTime int64
	MsgType string
	Event string
	Content string
	MsgId int64
}

func InitEvents(g *gin.Engine) (err error) {
	if orm, err = xorm.NewClient(); err != nil {
		log.Println("Init orm init failed ", err)
		return
	}

	events := Events{}

	g.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	g.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	g.GET("/signature",events.Signature)

	g.POST("/",events.Event)

	return
}

func (e *Events) Event(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	log.Println(string(buf[0:n]))

	mess := &Message{

	}

	err := xml.Unmarshal(buf[0:n],mess)
	if err != nil {
		c.String(http.StatusOK,"parser error")
		return
	}

	if mess.MsgType == "event" {
		if mess.Event == "subscribe" {
			e.Subscribe(mess)
		} else if mess.Event == "unsubscribe" {
			e.Unsubscribe(mess)
		} else {
		}
	} else if mess.MsgType == "text" {
		/*
		if mess.Content == "list" {

		} else if
		*/
	} else {

	}

	c.String(http.StatusOK,"")
	return
}

func (e *Events) Signature(c *gin.Context) {
	token := "viilebot"
	timestamp, err := c.GetQuery("timestamp")
	if !err || len(timestamp) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "args error",
		})
	}
	nonce, err := c.GetQuery("nonce")
	if !err || len(nonce) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "args error",
		})
	}
	signature, err := c.GetQuery("signature")
	if !err || len(signature) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "args error",
		})
	}
	echostr, err := c.GetQuery("echostr")
	if !err || len(echostr) == 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "args error",
		})
	}

	args := []string{token, timestamp, nonce}
	log.Println(args)
	sort.Sort(sort.StringSlice(args))
	sign := strings.Join(args, "")
	log.Println(sign)
	h := sha1.New()
	h.Write([]byte(sign))
	sh := h.Sum(nil)
	bs := fmt.Sprintf("%x", sh)
	log.Println(bs)
	log.Println(signature)
	if bs == signature {
		c.String(200, echostr)
	} else {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "args error",
		})
	}
}

func(e *Events) Subscribe(mess *Message) error {
	log.Println(mess)
	user := &models.User{
		OpenID:mess.FromUserName,
	}
	orm.Create(&user)
	return nil
}

func(e *Events) Unsubscribe(mess *Message) error {
	user := &models.User{
		OpenID:mess.FromUserName,
	}
	if orm.First(&user).RecordNotFound() {
		return nil
	}
	orm.Delete(user)
	orm.Delete(models.Sub{}, "user_id = ?", user.ID)

	return nil
}
