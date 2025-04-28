package main

import (
	"fmt"
	"log"

	"github.com/hmdnu/autocheck/service"
)

func main() {
	cookie, err := service.GetCookie()
	if err != nil {
		log.Fatalln("error getting the cookie", err)
	}
	responses, err := service.CheckIn(cookie.Token, cookie.Uid)
	if err != nil {
		log.Fatalln("error when checked in", err)
	}

	fmt.Println(responses)
}
