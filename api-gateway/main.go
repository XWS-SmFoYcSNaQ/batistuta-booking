package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gwServer := GetServer()

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err := gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
