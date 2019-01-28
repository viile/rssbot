package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/viile/rssbot/events"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt, syscall.SIGINT)

	g := gin.New()
	g.Use(gin.Recovery())

	// 调用主函数逻辑
	if err := events.InitEvents(g); err != nil {
		log.Fatal(err, "InitApi Addr:", "0.0.0.0:8818")
	}

	s := &http.Server{
		Addr:        "0.0.0.0:8818" ,
		Handler:      g,
		ReadTimeout:  1000 * time.Microsecond,
		WriteTimeout: 1000 * time.Microsecond,

		MaxHeaderBytes: 1 << 20,
	}

	// 启动 GIN 主函数
	go func() {
		s.ListenAndServe()

		log.Fatal(s.ListenAndServe())

		quit <- syscall.SIGINT
	}()

	<-quit
	log.Println("Shutting down server...")

	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatal("could not shutdown: ", err)
	}


	log.Println("Server exiting")
}
