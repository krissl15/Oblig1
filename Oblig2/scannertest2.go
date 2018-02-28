package main

import "fmt"

func main() {
	chDone := make(chan bool) //Denne skal sørge for at første go-funksjon får gjøre det den skal før neste begynner.
	ch1 := make(chan int, 2) //Kanalen som tallene skrives inn i.
	ch2 := make(chan int, 2)
	var i int
	var o int
	go func() {
		fmt.Println("Vennligst skriv et tall...")
		fmt.Scan(&i)
		fmt.Println("Setter tallet", i, "inn i kanalen. Skriv et tall til.")
		fmt.Scan(&o)
		fmt.Println("Setter tallet", o, "inn i kanalen.")
		ch1 <- i
		ch1 <- o
		chDone <- true //chDone settes til true, hvilket betyr at neste funksjon kan kjøre.

	}()
	func(){
		<-chDone //Det er denne som gjør at funksjonen venter på "true" signalet.
		fmt.Println(<-ch1)
		fmt.Println(<-ch1)
		ch2 <- i + o //Legger sammen tallene som skrives inn. Hvordan kan disse gjøres til ch1 istedenfor variablene?
		fmt.Println("Dine tall lagt sammen:")
		fmt.Println(<-ch2)

	}()
}