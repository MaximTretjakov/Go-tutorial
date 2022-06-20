package main

import "fmt"

func ping(ping chan<- string, mes string) {
	fmt.Println(mes)
	ping <- mes
}

func pong(ping <-chan string, pong chan<- string) {
	mes := <-ping

	if mes == "ping" {
		mes = "pong"
	} else {
		mes = "unknown command"
	}

	pong <- mes
}

func main() {
	pingCh := make(chan string)
	pongCh := make(chan string)

	go ping(pingCh, "ping")
	go pong(pingCh, pongCh)

	status := <-pongCh
	fmt.Println(status)
}
