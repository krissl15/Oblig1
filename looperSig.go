package main

import (
	"os"
	"time"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sign := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	start := time.Now()

	//SIGINT gir avslutningsmelding når du trykker CTRL + C. Man kan også bruke CTRL + Break for å pause programmet.
	//SIGTERM gir avslutningsmelding når programmet stopper uventet/crasher.
	//SIGQUIT gir avslutningsmelding når programmet stoppes manuelt.
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sign
		fmt.Println(sig)
		done <- true
	}()

	for i := 1;  ; i++ {
		fmt.Println("o")
		select {
		default: i++

		case <- done:
			elapsed := time.Since(start)
			fmt.Println("Programmet er avsluttet, og kjørte i", elapsed) //Avslutningsmeldingen.

			os.Exit(1)
		}
	}
}
