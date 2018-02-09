<!DOCTYPE html>
<html>

<body>
	
# Oppgave 4

## Oppgave 4A:
<br>
Se filen "ascii.go" for koden. 
<br>

## Oppgave 4B:
<br>
Se filen "sorting.go" for beste løsning. Jeg har og lagt til 2 andre eksempler i egne filer (asciiEKS2.go, asciiEKS3.go). Disse ligger i egne filer for mer ryddig kode.
<br>
###### Eksempel 1 (Koden i "sorting.go"): 
<br>
Dette eksempelet tar utgangspunkt i en []byte input og gjør den om til string. € er en del av extended ASCII i ISO-8859-15, og siden våre pc'er anvender ISO-8859-1 har vi ikke tilgang til tegnet € i default tegnsett. Derfor har jeg anvendt []rune (int32), og gjort dette videre til string. Om en prøver å bruke UTF-16 koden til € (8 364) i en byte, vil en få "overflow" error. 

###### Eksempel 2 (asciiEKS2.go): 
<br>
Dette eksempelet anvender to funksjoner: En til å konvertere symboler til hex-verdier, og en annen til å konvertere fra hex-verdier til string. Dette gir en "gratis" løsning på oppgaven, da en han bruke codepoints til å printe ut symboler, uten å vite hva det er i det hele tatt.

Eksempel 3 (asciiEKS3.go): 
<br>

## Oppgave 4C:
Se filen "iso_test.go" for koden.



</body>
</html>
