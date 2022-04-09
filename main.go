package main

import (
	"log"
	"e2e/msg"
)

func main() {
	log.Println("Message", msg.Msg("service B"))
}