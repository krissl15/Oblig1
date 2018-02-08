<!DOCTYPE html>
# OPPGAVE 1. Fyll ut tabellen
<html>

<body>

<table>
  <tr>
    <th>Binære tall</th>
    <th>Hexadesimaltall</th>
    <th>Desimaltall</th>
  </tr>
  <tr>
    <td><b>1101</b></td>
    <td><b>0xD</b></td>
    <td><b>13</b></td>
  </tr>
  <tr>
    <td><b>1101 1110 1010</b></td>
    <td>0xDEA</td>
    <td>3562</td>
  </tr>
  <tr>
    <td>1010 1110 0011 0100</td>
    <td><b>0xAF34</b></td>
    <td>44852</td>
  </tr>
  <tr>
    <td>1111 1111 1111 1111</td>
    <td>0xFFFF</td>
    <td><b>65535</b></td>
  </tr>
  <tr>
    <td>1 0001 0111 1000 1010</td>
    <td>0x1178A</td>
    <td><b>71562</b></td>
  </tr>
</table>


<br>

## OPPGAVE 1 A: Metode for å gjøre om fra hexadesimaltall til binære tall og motsatt, gjøre om fra desimaltall til binære tall og motsatt.
<br>
<br>
	<b>Binære tall -> hexadesimale tall</b>
  <br>
	1. For å gå fra binære tall til hexadesimale tall, deler vi de binære tallene opp i grupper på fire. Deretter legger vi sammen tallene ut ifra de fire første bitsene (8 4 2 1). Husk at A = 10, B = 11 osv.
  <br>
  <br>
	<b>Hexadesimale tall -> binære tall</b> <br>
	1. For å gå fra hexadesimale tall til binære tall, leser vi ett hexadesimaltall om gangen, og deretter skriver vi ned de binære verdiene for hvert hexadesimal.
  <br>
  F.eks A7
    <br> A = 1010
    <br> 7 = 0111
    <br> A7 = 1010 0111
  <br>
  <br>
<b>Binære tall -> desimaltall</b> <br>
	1. For å gå fra binære tall til desimaltall: les fra høyre til venstre, og legg sammen verdiene som har 1 tall. Alternativt kan vi regne ut ifra totallsystemet: <br>
	Verdien sin x posisjon ganges med 2^x, leses av fra høyre side: <br>
		Eks. 1010:
			<br> 0	&times; 2^0=0
			<br> 1	&times; 2^1  =  2
			<br> 0	&times; 2^2=0
			<br> 1	&times; 2^3  =  8
			<br> Sum = 8 + 2 = 10 
      <br>
      <br>
	<b>Desimaltall -> binære tall</b><br>
	Start med høyest mulig tall i tallrekken (venstre side) som går opp i angitt desimaltall, og legg sammen alle mulige tall til riktig desimaltall er funnet. Alle tall som går med, skal skrives ned 1 på.
  <br>

Alternativt tar vi heltallsdivisjon, og tar angitt desimaltall delt på 2. Dersom det er noen rester, skriv ned restverdien (alltid 1). Dersom det ikke er rester, skriv ned 0. Gjør slikt med alle tall helt til heltallsdivisjonen ender opp på 0.
<table>
  <tr>
    <th>Heltallsdivisjon</th>
    <th>Utregning</th>
    <th>Rest</th>
  </tr>
  <tr>
    <td>255</td>
    <td>255/2=127</td>
    <td>1</td>
  </tr>
  <tr>
    <td>127</td>
    <td>127/2=63</td>
    <td>1</td>
  </tr>
  <tr>
    <td>63</td>
    <td>63/2=31</td>
    <td>1</td>
  </tr>
  <tr>
    <td>31</td>
    <td>31/2=15</td>
    <td>1</td>
  </tr>
  <tr>
    <td>15</td>
    <td>15/2=7</td>
    <td>1</td>
  </tr>
  <tr>
    <td>7</td>
    <td>7/2=3</td>
    <td>1</td>
  </tr>
  <tr>
    <td>3</td>
    <td>3/2=1</td>
    <td>1</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1/2=0</td>
    <td>1</td>
  </tr>
  <tr>
    <td>0</td>
    <td>0/2=0</td>
    <td>0</td>
  </tr>
</table>
	Deretter skriver vi opp rest-tallet fra nederste linje og opp.<br>
  255 = 1111 1111  (ettersom 0 er på enden, vil den ikke være gjeldende)
<br>

 <b>OPPGAVE 1 B. Metode for å gjøre om fra desimaltall til hexadesimaltall og motsatt</b>
 <br>
 <br>
	<b>1. Hexadesimaltall -> desimaltall - F.eks F7:</b>
  <br>
		Samme metode som fra binære tall til desimaltall.
		Her regner vi med 16^x (hexadesimal = seksentallssystem), og regner ut i fra posisjonen til verdiene. 7 vil være i posisjon 0, mens F i posisjon 1, osv.
			<br>7 &times;16^0=7
			<br>7 &times; 16^1=240
			<br>Deretter adderer vi summen. F7 = 247.
      <br>
      <br>
			
 <b>2. Desimaltall -> Hexadesimaltall </b><br>
	Samme utførelse som fra desimaltall til binære tall, bare at vi deler på 16 nå, istedenfor 2. 

</body>
</html>
