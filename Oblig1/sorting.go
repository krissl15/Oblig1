package algorithms

// Les https://en.wikipedia.org/wiki/Bubble_sort

////////////////////opg2a///////////////
//swap funksjonen bytter om tallene i og j 
func swap(list []int, i, j int) {
	temp := list[j]
	list[j] = list[i]
	list[i] = temp
}

func Bubble_sort_modified(list []int) {
	sorted := true
	for sorted {
		sorted = false
		for i := 0; i < len(list)-1; i++ {
			if list[i+1] < list[i] {
				swap(list, i, i+1)
				sorted = true
			}
		}
	}
}
/////////////////opg2a slutt//////////////

// Implementering av Bubble_sort algoritmen
func BubbleSort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QuickSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}

////////////4B: Lag en funksjonen ExtendedASCIIText () i filen sorting.go , som skriver ut: " € ÷ ¾ dollar "
func ExtendedASCIIText(s []byte) string { // Tar []byte som argument
	var ascii = string(s)
	return ascii //<-- returnerer []byte argumentet som string
}

func main() {
	asciiR := []rune{34, 8364, 32, 190, 32, 247, 32, 100, 111, 108, 108, 97, 114, 34}// rune pga €=8364 codepoint
	asciiS := string(asciiR)  // rune->string
	asciiByteSlice := []byte(asciiS) // string->[]byte
	
	fmt.Printf("%s", ExtendedASCIIText(asciiByteSlice)) // output: "€ ¾ ÷ dollar"
       
}
///////////////////SLUTT 4B//////////////////// se flere eksempler (asciiEKS1.go, asciiEKS2.go)

