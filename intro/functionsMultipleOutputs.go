package main

import (
	"fmt"
	"math"
)

func maxmin (x float64, y float64) (float64, float64) {
	M := math.Max(x,y)
	m := math.Min(x,y)
	return M, m
}

func main(){
	M, m := maxmin(18,9)
	fmt.Print(" I numeri ordinati sono ", m, " e ", M)
}