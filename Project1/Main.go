package main

import ("fmt"
		"strconv")

func main() {

	
	var esi string = "Ecole national Superieur informatique"
	polytech := "Ecole National Superieur Polytechnique"
	

	fmt.Printf("%v ,%T \n", esi, esi)
	fmt.Printf("%v ,%T \n", polytech, polytech)

	//type conversion 
	var i int 
	var k float32 = float32(i)

	fmt.Printf("value: %v; type: %T \n",i,i)
	fmt.Printf("value: %v; type: %T \n",k,k)

	//converting float to integer 
	var l int = 42
	var m string = strconv.Itoa(l)

	fmt.Printf("value: %v ; type:  %T\n",l,l)
	fmt.Printf("value: %v ; type:  %T\n",m,m)
	

}
