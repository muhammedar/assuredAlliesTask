package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"dictionaryManager.com/server"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Printf("Assured Allies Server Instance Starting...")

	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
