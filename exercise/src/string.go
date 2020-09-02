package main

import (
	"fmt"
	"strings"
)




func main(){
	str1 := "the quick red fox"
	str2 := "jumped over"

	stringLength, err := fmt.Println(str1,str2)
	if err == nil {
		fmt.Println("String Length:",stringLength)
	}


	stringLength1, _ := fmt.Println(str1,str2)
	fmt.Println("String Length:",stringLength1)


	aNumber := 42
	isTrue := true

// %v => verb
	fmt.Printf("value of aNumber %v\n",aNumber)
	fmt.Printf("value of isTrue %v\n",isTrue)
	fmt.Printf("value of aNumber as float  %.2f\n",float64(aNumber))

	fmt.Println(strings.ToUpper(str1))	
	fmt.Println(strings.Title(str1))	

	fmt.Println("isEqualString?", (str1 == str2))
	fmt.Println("isEqualString non case sensitive?", strings.EqualFold(str1 , str2))
	fmt.Println("string containes the ?", strings.Contains(str1 , "the"))
}