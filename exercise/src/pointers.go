package main

// Go supports pointers, allowing you to pass references to values and records within your program.

import "fmt"

func zeroval(ival int){
	ival = 0
}

func zeroptr(iptr *int){
	*iptr = 0
}



func main(){
	var p *int
	if p != nil {
		fmt.Println("value of p:",*p)
	} else {
		fmt.Println("p is nil")
	}

	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("value of p:",*p)
	} else {
		fmt.Println("p is nil")
	}

	*p = *p * 2
	fmt.Println("value of p:",*p)


	i := 1
	fmt.Println("initial:", i)
    

    // zeroval has an int parameter, so arguments will be passed to it by value. 
    // zeroval will get a copy of ival distinct from the one in the calling function.
    zeroval(i)
    fmt.Println("zeroval:", i)


	// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. 
	// The *iptr code in the function body then dereferences the pointer from its memory address
	// to the current value at that address. Assigning a value to a dereferenced pointer 
	// changes the value at the referenced address
    zeroptr(&i)	//The &i syntax gives the memory address of i, i.e. a pointer to i.
    fmt.Println("zeropts:",i)
	
	fmt.Println("pointer:", &i)		

	//zeroval doesn’t change the i in main, but zeroptr does because 
	//it has a reference to the memory address for that variable.


}