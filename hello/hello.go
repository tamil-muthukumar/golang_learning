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

	/**
	Conditionals
	Always need curly braces around if block
	**/
	statePop := map[string]int{
		"Cali":    4345435435,
		"Texas":   343243242,
		"Florida": 32432424,
	}

	if pop, ok := statePop["Cali"]; ok {
		fmt.Println(pop)
	}
	/**
	Switch:
	break keyword is implicit
	*/
	switch 4 {
	case 1, 5, 10:
		fmt.Println("Cases 1, 5, 10")
	case 9, 4, 3:
		fmt.Println("cases 9, 4, 3")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("Cases 1, 5, 10")
	case 9, 4, 3:
		fmt.Println("cases 9, 4, 3")
	}

	switchI := 10
	switch {
	case switchI <= 10:
		fmt.Println("less than 10")
		fallthrough // lets multiples cases to execute
	case switchI <= 20:
		fmt.Println("less than 20")
	}
	//Type Switch
	var switchTypeI interface{} = 1
	switch switchTypeI.(type) {
	case int:
		fmt.Println("type int")
	}

	/**
	Looping
	i++ and j++ cannot be assigned to a variable like below
	i, j = i++, j++

	because i++ is an expression on its own

	Loops:
	for initializer; test; incrementer {}
	for test {}
	for {}
	*/
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// continue and break keywords can be used in for loop
	iFor := 0
	for {
		fmt.Println(iFor)
		iFor++
		if iFor == 5 {
			break
		}
	}

	// Loops with collections
	// range is a special form of for loop

	forSlice := []int{1, 2, 3}
	for k, v := range forSlice {
		// k=> index, v=> value
		fmt.Println(k, v)
	}

	// same syntax for looping over a map
	for k, v := range statePop {
		fmt.Println(k, v)
	}

	// ignore the key with the  _
	for _, v := range statePop {
		fmt.Println(v)
	}

	/**
	Defer, Panic, Recover

	Defer: Executes the statement preceded with go after it finishes executing the last statement in the function and before it returns
	If there are multiple defer statements it is executed in LIFO order
	Defer executed BEFORE panic is called.
	Avoid using calling defer in loops as it would defer closing the resources until the end of the function is reached
	*/

	deferA := "start"

	// Prints "start" as the defer statement is evaluated at the time the defer is called and not at the time when defer is executed in the function
	defer fmt.Println(deferA)
	deferA = "end"

	/** Panic
	panicA, panicB := 1, 0
	panicAns := panicA / panicB
	fmt.Println(panicAns) // returns: panic: runtime error: integer divide by zero
	panic("something bad happened")
	Use for unrecoverable events - cannot obtain TCP port for web server
	Don't use when file can't be opened unless it is critical
	If nothing handles panic, eventually panic will go up the call stack and hits the go runtime.

	**/

	/**
	Recover:
	Anonymous function - a function which doesn't have a name. Nothing else can call this
	Only useful inside deferred functions because fo the behavior of panic functions. When an application starts to panic,
	it no longer executes the rest of the function, but will execute deferred function
	calling recover inside defer func allows us to handle the panic on our own.
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			panic(err) // throw this to the calling function if you think that function to handle it
		}
	}()

	/**
	POINTERS
	-> Creating Pointers
	-> Dereferencing pointers
	-> the new function
	-> working with nil
	-> types with internal pointers
	*/

	ptrA := 42
	ptrB := ptrA // copies value of a into b
	fmt.Println(ptrB)

	var ptrC int = 79
	var ptrD *int = &ptrC // declares this var as a pointer to an integer

	fmt.Println(ptrC, ptrD)
	fmt.Println(&ptrC) // notice that ptrD above points to the same address as of ptrC

	// An asterisk * in front of as pointer variable is used to defreference - meaning, ask go runtime to
	// go to the location the pointer var ptrD is pointing to and pull the value out
	fmt.Println(*ptrD) // prints 79

	*ptrD = 84
	fmt.Println(ptrC, *ptrD) // prints 84 for both vars

	ptrArr := [3]int{1, 2, 3}
	ptrArrB := &ptrArr[0] // gets the address of array index 0
	ptrArrC := &ptrArr[1]

	fmt.Println(ptrArr, *ptrArrB, *ptrArrC) // prints [1 2 3] 1 2

	// pointers with structs
	type myStruct struct {
		foo int
	}

	// this is one way of initializing a pointer to an object
	var ms *myStruct // a pointer to myStruct
	ms = &myStruct{foo: 420}
	fmt.Println(ms) // prints &{42}
	fmt.Println(*ms)
	fmt.Println(ms.foo) // prints 42 - structs are automatically dereferenced for us. this is expr is equivalent to (*ms).foo

	// Another way is using the builtin new function
	// note that it cannot be initialized at the same time

	// Types with internal pointers: slices and Map has a pointer to the first element in the underlying array. Hence the below example
	ptrSliceA := []int{1, 2, 3}
	ptrSliceB := ptrSliceA            // gets the address of array index 0
	fmt.Println(ptrSliceA, ptrSliceB) // prints [1 2 3] [1 2 3]
	ptrSliceA[1] = 42
	fmt.Println(ptrSliceA, ptrSliceB) // [1 42 3] [1 42 3] - this is because slice is already a pointer to an underlying array

	/**
	Functions

	Entrypoint into application is main() in package main


	*/
	message := "message"
	greeting := "greeting"
	sayMessage("Hello say message")
	sayGreeting(message, greeting)
	fmt.Println("Before calling pointer ", greeting)
	sayGreetingWithPtr(&message, &greeting)
	fmt.Println("After calling pointer ", greeting) // prints "Greetings changed with pointer"
	variadicParams(1, 2, 3, 4, 5, 6)
	retVal := pointerRetVal(1, 2, 3, 4, 5, 6)
	fmt.Println(*retVal)

	retVal1, err := multRetVals(1, 0)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(retVal1)
	}

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()

	g.greetWithPtr()
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// args of same type (string in this example) can be delimited
func sayGreeting(msg, greeting string) {
	fmt.Println(msg)
}

func sayGreetingWithPtr(msg, greeting *string) {
	fmt.Println(*msg, *greeting)
	*greeting = "Greetings changed with pointer"
}

/** Variadic Params
There can be only one arg with variadic params and it should be the last arg in the func
**/

func variadicParams(values ...int) int /**this is the return type*/ {
	fmt.Println(values) // slice
	// you can loop through those values and do whatever
	return 0
}

func pointerRetVal(values ...int) *int {
	result := 0

	/**
	When you are returning a reference to a variable from a function, the variable's memory that's local to the function doesn't get cleaned up.
	it's automatically promoted to be on the heap memory
	*/
	return &result // return the reference to the result var
}

func namedRetVal(values ...int) (result int) {
	// Named returnvalue => result is the name of the variable that will be returned from the method

	return
}

// Multiple return values from the function
func multRetVals(a, b int) (int, error) {

	if b == 0 {
		return 1, fmt.Errorf("b is set to 0") // this is set in error return variable
	}
	// success case
	return 0, nil // nil representing that there were no errors
}

type greeter struct {
	greeting string
	name     string
}

/*
This is a method. The presence (g greeter) a known context makes it a method( and not a function)
The known context could be any type

When this method is called, its called with greet object defiined in main above. The method receives a copy of the object (as opposed to a pointer to the object)
*/
func (g greeter) greet() {
	// we are operating on the copy of copy object
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) greetWithPtr() {
	// we are operating on the pointer of the object
	fmt.Println(g.greeting, g.name)
}
