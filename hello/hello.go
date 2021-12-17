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

	// Two dim matrix
	var identityMatrix = [3][3]int{}
	identityMatrix[0] = [3]int{1, 0, 1}

	a := [...]int{1, 2, 3}
	bCopy := &a
	bCopy[1] = 5
	fmt.Println(a)
	fmt.Println(bCopy)

	// Slice
	//
	slc := []int{1, 2, 3}
	fmt.Println(slc)
	fmt.Printf("%v", len(slc))
	fmt.Printf("%v\n", cap(slc))

	/**
	number of elements in the slice doesn't the size of the backing array. Slice is a projection of the underlying array
	Slices are by default reference types. They refer to the same underlying data
	*/

	slc1 := slc
	slc1[1] = 5 // this would change the value in slc as well
	fmt.Println(slc)
	fmt.Println(slc1)
	/**
	Other ways of creating slices
	*/
	aSlc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bSlc := aSlc[:]   // slice of all elements
	cSlc := aSlc[3:]  // slice from 4th to end
	dSlc := aSlc[:6]  // Slice first 6 elements
	eSlc := aSlc[3:6] // slice the 4,5,6th elements
	fmt.Println(aSlc)
	fmt.Println(bSlc)
	fmt.Println(cSlc)
	fmt.Println(dSlc)
	fmt.Println(eSlc)

	// Chaning the value in the underlying array, will change the value in all the slices referencing to it
	aSlc[5] = 42
	fmt.Println(aSlc)
	fmt.Println(bSlc)
	fmt.Println(cSlc)
	fmt.Println(dSlc)
	fmt.Println(eSlc)

	// Make function arg2: length of the slice, arg3: capacity
	aMake := make([]int, 3, 100)
	aMake = append(aMake, 55, 66, 76, 88, 98)
	fmt.Println(aMake)

	bMake := make([]int, 0, 0)
	fmt.Println(bMake)

	// append
	bMake = append(bMake, 1, 2, 3, 4, 5)
	fmt.Println(bMake)
	fmt.Println("Length:", len(bMake))
	fmt.Println("Capacity:", cap(bMake))

	// concatenate slices
	// use spread operator (...)  to append
	bMake = append(bMake, aMake...)
	fmt.Println(bMake)

	// remove elements in the middle lets say remove 4
	bMake = append(bMake[:2], bMake[4:]...)
	fmt.Println(bMake)

	/**
	MAPS
	syntax: map[key]value
	An array is a valid key type to the map, a slice is not
	*/

	dMap := map[string]int{
		"Cali":  12345,
		"Texas": 34567,
	}
	fmt.Println(dMap)
	// other different types of map declaration
	eMap := make(map[string]int)
	eMap = dMap
	fmt.Println(eMap)
	eMap["Florida"] = 5345345
	fmt.Println(eMap)

	// delete func
	delete(eMap, "Florida")

	// 1st param gets the value, 2nd params tells you whether the element was in the map or not
	// val would default to the type default - 0 for int type values
	val, ok := eMap["Florida"]
	fmt.Println(val, ok)

	// Side Effect:  By default, the maps are also pass by reference. Any mutation would affect the original
	fMap := eMap
	delete(fMap, "Texas")
	fmt.Println(eMap)

	/**
	Struct
	Allows to mix any type of data together

	Casing rules apply for both the struct name and the fields within it
	uppercases - export and lowercases private to the module

	Unlike maps and slices, when you copy a struct into another, we are copying data into (and not a reference).
	*/
	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}

	// alternate syntax: positional syntax is also allowed. Dont use it though - maintenance problem
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Poewty",
		companions: []string{
			"Joe Shaw",
			"Shaw Joe",
			"Joo joo",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[2])

	// anonymouse struct
	aStruct := struct {
		name string
	}{name: "Joey Joey"}
	fmt.Println(aStruct)

	bStruct := aDoctor // copies the values and the type into bstruct
	bStruct.actorName = "Slly Silly"
	fmt.Println(aDoctor) // retains the originally assigned vaLUE
	fmt.Println(bStruct) // RETAINS THE CHANGES WE DID TO IT : Silly Silly

	// if we do want to point to the same underlying data, we can use & operator
	cStruct := &aDoctor
	cStruct.actorName = "name changed to Silly"
	fmt.Println(aDoctor)
	fmt.Println(cStruct)

	/**
	Composition
	Go doesn't support inheritance
	*/

	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal
		SpeedKMPH float32
		CanFly    bool
	}

	bBird := Bird{}
	bBird.Name = "Emu"
	bBird.Origin = "Aussie"
	bBird.SpeedKMPH = 50
	bBird.CanFly = false
	fmt.Println(bBird)

	/**
	tags
	Tags can be added to structs field to describe field
	You can add some optional validation params here. Parse it with reflection and m
	*/

	type Animal1 struct {
		Name   string `required max: "100"`
		Origin string
	}
}
