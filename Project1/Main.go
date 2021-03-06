package main

import (
	"fmt"
	"strconv"
)

func main() {

	var esi string = "Ecole national Superieur informatique"
	polytech := "Ecole National Superieur Polytechnique"

	fmt.Printf("%v ,%T \n", esi, esi)
	fmt.Printf("%v ,%T \n", polytech, polytech)

	//type conversion
	var i int
	var k float32 = float32(i)
	var s float64 = float64(k)

	fmt.Printf("value: %v; type: %T \n", i, i)
	fmt.Printf("value: %v; type: %T \n", k, k)
	fmt.Printf("value: %v; type: %T \n", s, s)

	//converting integer to string
	var l int = 42
	var m string = strconv.Itoa(l)

	fmt.Printf("value: %v ; type:  %T\n", l, l)
	fmt.Printf("value: %v ; type:  %T\n", m, m)

	//complex numbers
	var cmp1 complex64 = 12 + 3i
	var cmp2 complex128 = 2 + 5i

	fmt.Printf("value: %v ; type:  %T\n", cmp1, cmp1)
	fmt.Printf("value: %v ; type:  %T\n", real(cmp1), real(cmp1))
	fmt.Printf("value: %v ; type:  %T\n", imag(cmp1), imag(cmp1))
	fmt.Printf("value: %v ; type:  %T\n", cmp2, cmp2)
	fmt.Printf("value: %v ; type:  %T\n", real(cmp2), real(cmp2))
	fmt.Printf("value: %v ; type:  %T\n", imag(cmp2), imag(cmp2))

	//converting integer to complex
	r := 64
	im := 5
	var o complex128 = complex(float64(r), float64(im)) //expects float arguments
	fmt.Printf("value: %v ; type:  %T\n", o, o)

	
	//Arrays and Slices
	//
	var table [2]string
	table2 := [...]int{1, 2, 3}

	fmt.Printf("value: %v ", table)
	fmt.Printf("value: %v ", table2)

	var length int = len(table) // gets length of he array
	fmt.Println("the table lengh is :", length)


	// LOOPING
	for i,j := 0,0 ;i < 10 ; i,j = i+1,j+1{
		fmt.Println(i, j)
	}
}
