<!DOCTYPE html>
# OPPGAVE 1. Fyll ut tabellen
<html>

<body>
	
# Oppgave 2 

## Oppgave 2A:
Se fil "sorting.go"

## Oppgave 2B:
Se fil "sorting_test.go"

## Oppgave 2C:
![BenchTestCMD](https://i.imgur.com/gStphWH.png)
<br>
Vi kan se at big-O for alle de tre algoritmene vi har testet så er det QuickSort som kommer best ut i forhold til hastigheten mellom BubbleSort, BubbleSortModified og QuickSort. Vi ser at QuickSort går gjennom ganske mange flere løkker enn BubbleSort med raskere tid. <br>
**Eksempel**: QuickSort på 100 elementer tar 4731 nanosekunder per løkke i forhold til BubbleSort som bruker 26358 nanosekund per løkke på 100 elementer. Jo flere elementer vi får hvor lengere tid vill prosessen ta.
<br>
![benchTestGraph](https://i.imgur.com/s2xL3sv.png)<br>
Hvis vi går inn på http://bigocheatsheet.com/ og ser på grafene i Big-O Complexity Chart så kan vi se at QuickSort grafen vår ligger på Best/Average med Ω(n log(n)). Bubblesort og BubbleSortModified grafen ligger på Average/Worst med Θ(n^2).
<br>
<br>
# Oppgave 3
Se filen "LoopSig.go" for koden.
<br>
**"Hvor mye minne og CPU bruker programmet når det kjører."**
<br>
![evigLoop](https://i.imgur.com/9TV5n9e.png) 
<br>
Prosessen "eee" på bildet er programmet med den evige løkken. Vi ser at funksjonen bruker 6,6% av prosessoren, og 4,5MB minne.
<br>



</body>
</html>
