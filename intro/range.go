package main

import "fmt"

func main(){
	sl := []string{"a", "b", "c", "d"}
	for i:=range sl{
		fmt.Print(i, ", ")
	}

}