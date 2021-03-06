package main

import "fmt"
import "time"

//There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.

func main() {

	// Note that you don’t need parentheses around conditions in Go, 
	// but that the braces are required.

	if 7%2 == 0 {
		fmt.Println("7 is event")
	} else {
		fmt.Println("7 is odd")
	}

	// A statement can precede conditionals; any variables declared 
	// in this statement are available in all branches.
	// this is not ok
	// if ( num := 9;  num < 0 ) 
	if num := 9; ( num < 0 ) {
        fmt.Println(num, "is negative")
    } else if ( num < 10 ) {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }



    i := 2
    fmt.Print("Write " , i , " as " )
    switch i {
    	case 1 :
    		fmt.Println("one")
    	case 2 :
    		fmt.Println("two")
    	case 3 :
    		fmt.Println("three")
    }


    // You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}


	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type `switch` compares types instead of values.  You
	// can use this to discover the type of an interface
	// value.  In this example, the variable `t` will have the
	// type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}