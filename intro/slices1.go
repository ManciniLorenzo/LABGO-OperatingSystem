package main

import "fmt"

func main(){
	sl := []string{"a", "b", "c", "d"}
	fmt.Println("Una parte dello slice: ", sl[1:3])
	sl[2]="e"
	fmt.Println("Lo slice modificato: ", sl)
	fmt.Println("Lunghezza dello slice: ", len(sl))
	fmt.Println("Capienza dello slice: ", cap(sl))

}