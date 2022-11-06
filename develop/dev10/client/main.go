package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	telnet := NewTelnet("", "", 10)
	if err := telnet.getTimeout(os.Args[2]); err != nil {
		log.Fatalln(err)
	}
	telnet.getSocket(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
	err := telnet.dial()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("connection established")
	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			<-sigCh
			fmt.Println("Closing the connection")
			os.Exit(0)
		}
	}()
	if err := telnet.write(); err != nil {
		log.Fatalln(err)
	}
	wg.Wait()
}
