package main

import "fmt"

func main(){
	sl := []string{"a", "b", "c", "d"}
	sl = append(sl, "f") // aggiunta di un elemento allo slice
	fmt.Println(sl)

	sl = append(sl, "", "y", "z") // aggiunta di tre elementi allo slice
	fmt.Println(sl)
}