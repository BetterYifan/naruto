package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	server := &http.Server{
		Addr:         ":4001",
		WriteTimeout: time.Second * 2,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// 自定义默认的mux
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/say", sayBye)

	server.Handler = mux

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("err ")
		}
	}()

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello,this is version 1"))
	//})
	//http.HandleFunc("/say", sayBye)

	//http.ListenAndServe(":4000", mux)
	server.ListenAndServe()
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}

func sayBye(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("bye-bye"))
}
