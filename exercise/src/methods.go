// Go supports _methods_ defined on struct types.

package main

import ( 
	"fmt"
	"utils"

)


type rect struct {
	width, height int
}

// This 'area' method has a _receiver type_ of 	'*rect'.
func ( r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types
// Here is an example of value receiver
func ( r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println(r.area())
	fmt.Println(r.perim())

	// Go automatically handles conversion between values and pointers for method calls. 
	// You may want to use a pointer receiver type to avoid copying on method  calls 
	// or to allow the method to mutate the receiving struct.
	rp := &r

	// Here we call the 2 methods defined for our struct.
	fmt.Println(rp.area())
	fmt.Println(rp.perim())


	n1, l1 := fullName("asdasda"  , "zzzzz"  ) 
	fmt.Printf("FullName : %v, number of char %v",n1 , l1 )

	n1, l1 = utils.fullNamesNakedReturn("asdasda"  , "zzzzz"  ) 
	fmt.Printf("FullName : %v, number of char %v",n1 , l1 )

}

func fullName(f , l string ) (string , int ){
	full := f + "  " + l 
	length := len(full)
	return full ,length
}


// func fullNamesNakedReturn(f , l string ) (full string , length int ){
// 	full = f + "  " + l 
// 	length = len(full)
// 	return 
// }



//NOTES : 
// in Go there's no function or method overloading. you can't have two functions of the same name within the same package.
// Also all of the functions that I've created here start with a lower case initial character. That makes them private to the current package,
// that is they aren't exported for use outside this package. 
// If you change their initial character to upper case though, it then becomes public, accessible to the rest of the application.
