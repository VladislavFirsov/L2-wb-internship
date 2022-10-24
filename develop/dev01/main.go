package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func main() {
	showCurrentTime()
}

func showCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Println("something went wrong", err)
	}
	fmt.Println(time)
}
