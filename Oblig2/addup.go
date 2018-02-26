package main

import (
"os"
"os/signal"
"syscall"
"fmt"
)

 func main() {
	sign := make(chan os.Signal, 1)
	done := make(chan bool, 1)
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
			fmt.Println("Programmet er avsluttet, og kjørte i") //Avslutningsmeldingen.

			os.Exit(1)
		}
	}
}

