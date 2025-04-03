package main

import "fmt"

var a, b, c bool // zero value per bool: false

func main(){
	var i int // dichiarazione senza inizializzazione
	fmt.Println(a, b, c, i)
}