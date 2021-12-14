package main

import (
	"fmt"
	"strconv"
)

// package levels variables can be declared here
// it can only be a var type and cannot be := syntax
var globalVar float32 = 38.

// var could be made into a block to avoid redundancy
var (
	string1 = "Eliza"
	string2 = "Samantha"
)

func main() {
	fmt.Println("halo")

	/**
	VARIABLES
	3 different ways to declare a variable
	Visibility:
	lowercase variables are visible to all the files local to the package
	uppercase variables are visible across packages, it's exported from the package and globally visible
	block scope - visible within the block/method
	*/
	var i int
	i = 42
	fmt.Println(i)

	// second way
	var secondVar = 34
	fmt.Println(secondVar)

	// third method
	k := 99
	fmt.Println(k)

	fmt.Println(globalVar)

	// type conversions
	var j int32 = 42
	fmt.Printf("%v, %T \n", j, j)

	var l float32
	l = float32(j)
	fmt.Printf("%v, %T \n", l, l)

	/**
	same conversion strategy doesn't work well when you convert int -> string
	needs a string conversion package called strconv package
	*/
	var m string
	m = string(j)
	fmt.Printf("%v, %T \n", m, m) // prints *, converts to ascii equivanlent

	var n = strconv.Itoa(i)
	fmt.Printf("%v, %T \n", n, n)

	// == Boolean
	b := true
	fmt.Println(b)

	/**

	Signed Integer type size range
	int8    -127 to 128
	int16   -32k to 32k
	int32   -2B to 2B
	int64   -9Quntillion to 9Q

	unsigned integers:
	uint8   0 to 255
	uint16  0 to 65K
	uint32  0 to 4B

	Bit shift ops << >>

	*/

	/**
	Floating point numbers
	float32 or flat64
	n := 3.14
	n := 13.7e72 // exponential notaion

	We'll get a floating point result when botht he operands are fp
	*/

	/**
	complex numbers
	*/
	var complex complex64 = 1 + 2i
	fmt.Println(complex)

	/**
	Strings
	by default array based / collection of letters
	UTF8 encoding by
	immutable
	s1 + s2 => concat
	Strings in go are aliases for bytes
	*/
	str := "this is a string"
	fmt.Printf("%v, %T \n", str[2], str[2])         // prints 105, uint8
	fmt.Printf("%v, %T \n", string(str[2]), str[2]) // prints i, uint8

	// Byte slices
	byts := []byte(str)
	fmt.Printf("%v, %T \n", byts, byts) // prints the ascii values of each char in the str

	/**
	Constants
	Naming convention: myConst
	Follows the same casing rules for visibility
	shadowing allowed - local constant with the same var name can shadow the global const
	const and var can be used together for different ops

	values must be calculable at compile time. meaning,we cannot use anything from libraries
	PascalCase for exported constants
	camelCase for internal constants
	Typed constants work like immutable variables. But they can interoperate only with same type

	Untype constants work like literals

	*/
	const myConst = 42
	fmt.Printf("%v, %T \n", myConst, myConst)

	var v int16 = 27
	// since myconst doesn't have a type declared, var v's type is considered for the result
	fmt.Printf("%v, %T \n", myConst+v, myConst+v)

	/**
	Enumerated constants

	iota is a counter we can use when we are creating enumerated constants
	*/
	const (
		const1 = iota //0
		const2        //1
		const3        //2
	)

	const (
		_  = iota             // ignore first value by assigning to blank identifier
		KB = 1 << (10 * iota) // iota = 1
		MB                    // everything in the rest of the list is incrementally assigned
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	const (
		isAdmin        = 1 << iota // bitshifted 0 places  0001
		isHeadQuarters             // bitshifted two places 0010
		canSeeFinancials
		canSeeAfrica
		canSeeAsia
	)

	var roles byte = isAdmin | canSeeFinancials | canSeeAsia
	fmt.Println("Roless %b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin&roles == isAdmin) // roles act as bitmask

	/**
	Arrays and Slices
	elements are contiguous in memory

	other valid syntaxes
	grade := [...]int{97, 85, 93}
	*/

	grade := [3]int{97, 98, 99}
	fmt.Println(grade)

	var students [3]string //- declares an empty array
	students[0] = "Lisa"

}
