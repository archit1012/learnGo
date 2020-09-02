// _Maps_ are Go's built-in [associative data type](http://en.wikipedia.org/wiki/Associative_array)
// (sometimes called _hashes_ or _dicts_ in other languages).

// Understanding go Runtine :  Go has runtime that's statically linked into each compiled application. Along with any external packages that are declared
// and imported in the application source code. The runtime operates in its own background threads. 
// make() and new() : These functions can be used to declare instances of maps, slices, and another type called channels
// The new() function allocates but doesn't initialize memory. When you allocate a map for example, using the new() function, 
// you'll get back a memory address indicating the location of the map, but the map has zeroed memory storage. 
// And if you try to set a value in the map, it'll cause an error. 

// In contrast, the make() function both allocates and initializes memory. 
// You'll get back the memory address, just like you do with new(), but the storage is non-zeroed and ready to accept values. 
// In most applications, you'll do best to use the make() function unless you have a specific reason to use the new() function. 
// If you don't do this correctly you can  casue a panic

package main

import "fmt"

func main() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)		// false syntax var m map[string]int  m["key"] = 42  , fmt.println(m)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.
	fmt.Println("map:", m)

	for k, v := range m{
		fmt.Printf("%v %v\n",  k , v )			
	}

	// Get a value for a key with `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs from
	// a map.
	delete(m, "k2")
	fmt.Println("map:", m)
	

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.

	_, prs := m["k2"]
	fmt.Println("prs:", prs)		// outputs : // prs: false

	// this code returns value as 0, even if key does not exist
	// so to fix that we have to use above code _ to make sure value is not zero
	prs1 := m["k2"]
	fmt.Println("prs:", prs1)

	
	// You can also declare and initialize a new map in
	// the same line with this syntax.

	// probably make is not require as we are giving value here itself
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)		// output map: map[bar:2 foo:1]
}
