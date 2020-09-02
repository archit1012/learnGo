package main

import ( 
	"fmt"
	"math"
)


// Go is all about type

const s string = "constant" 			//const declares a constant value.

func main() {
    var anInteger int = 42 // is similar to anInteger := 42
    var aString string  = "This is go "   //// is similar to aString := "This is go"

    fmt.Println(anInteger)
    fmt.Println(aString)



//Constant
    const  aConstInteger int = 42 // is similar to anInteger := 42
    const aConstString  = "This is go "   //// is similar to aString := "This is go"
    fmt.Println(aConstInteger)
    fmt.Println(aConstString)

//bool 
// Fixed integer type : int8 int16 int32 int64 uint8 uint16 uint32 uint64. , float32 float64, complex64 complex128
// Aliases : byte  uintptr , [ uint int => 32 or 64 bit depends on OS ]
// Complex data type : Arrays, Slices, Maps, Structs 
// Language Organization : Functions - is a type that is why one can pass to another function as an arguement , Interfaces, Channels
// Data management : pointers

    fmt.Println("hello world")
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)


    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2				//var declares 1 or more variables. You can declare multiple variables at once.
    fmt.Println(b, c)

    var d = true           //Go will infer the type of initialized variables.
    fmt.Println(d)

    var e int 				//Variables declared without a corresponding initialization are zero-valued.
    fmt.Println(e) 		// For example, the zero value for an int is 0.

    f := "apple" 			//The := syntax is shorthand for declaring and initializing a variable, 
    fmt.Println(f) 			//e.g. for var f string = "apple" in this case.


    fmt.Println(s)

    const n = 500000000  	 // A const statement can appear anywhere a var statement can.

    const co = 3e20 / n  	// Constant expressions perform arithmetic with arbitrary precision.
    fmt.Println(co)

    fmt.Println(int64(co)) 	//A numeric constant has no type until it’s given one, such as by an explicit conversion.

    fmt.Println(math.Sin(n))	//A number can be given a type by using it in a context that requires one,
    							//Such as a variable assignment or function call. For example, here math.Sin expects a float64.


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
    fmt.Printf("value of aNumber as fload  %.2f\n",float64(aNumber))                        
    fmt.Printf("Data Type : %T %t %T \n",str1, isTrue , aNumber)                        

    myString := fmt.Sprintf("Data Types as var : %T %t %T \n",str1, isTrue , aNumber)                        
    fmt.Printf(myString)
}

// //    foo()
// //      foo1()
// //      varkeyword()
// createyourtype()
//     for i := 0; i < 10; i++ {
//     	if i%2 == 0 {
// 	    	fmt.Println(i)	
//     	}
//     } 
// }


// type hotdog int
// var var_mytype hotdog
// var var_default_type int

// func createyourtype(){
// 	var_mytype = 100
// 	fmt.Println(var_mytype)	
// 	fmt.Printf("%T\n",var_mytype)
	
	
	
// 	// b = int 
// 	// var_mytype = b is not valid
	
// 	var_default_type = 200
// 	fmt.Println(var_default_type)	
// 	fmt.Printf("%T\n",var_default_type)
	
// 	var_default_type = int (var_mytype)  // call it conversion not casting 
	 
// 	fmt.Println(var_default_type)	
// 	fmt.Printf("%T\n",var_default_type)
	
// }

// func foo(){
// 	n,err := fmt.Println("I am in foo", 42 , true )
// 	fmt.Println(n)
// 	fmt.Println(err)
	
	
// //	below one runs here but not correct , doesn't run onlin
// //	n, _ := fmt.Println("I am in foo", 42 , true )
// //	fmt.Println(n)

// 	m, _ := fmt.Println("I am in foo", 42 , true )
// 	fmt.Println(m)

// }

// func foo1(){
// 	x := 42
// 	fmt.Println(x)
// 	x = 41
// 	fmt.Println(x)
	
// 	y := 42 + 24
// 	fmt.Println(y)
// }


// var y = 42
// // or var y int = 42

// // y := 42 won't work
// var z int // assigns 0 value to z "the zero value"
// var a string = `this is a string with backquote `
// var b string = "this is a string with double quote "

// func varkeyword(){
// 	x := 41
// 	fmt.Println(x)
// 	fmt.Printf("%T\n",x)
// //	y = "This is not allowed as go is static langauge"
// 	fmt.Println(y)
// 	fmt.Printf("%T\n",y)
	
// 	fmt.Println(a)
// 	fmt.Printf("%T\n", a)
// 	fmt.Println(b)
// 	fmt.Printf("%T\n", b)
	
	
// }


// Print , Printf,  Println
// Sprint , Sprintf Sprintln
// Fprint , Fprintf Fprintln