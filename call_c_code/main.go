package main

// #include <stdlib.h>
// #include "csrc/test.h"
//
// static int add2_macro_wrapper(int a, int b) {
//  return add2(a,b); // C macros aren't supported yet, we must wrap it
// }
//
// static double average2_wrapper(int a, int b) {
//   return average(2, a, b); // Variable argument methods aren't supported yet, we must wrap it
// }
// static double average3_wrapper(int a, int b, int c) {
//   return average(3, a, b, c); // Variable argument methods aren't supported yet, we must wrap it
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func callAdd1() {
	a := C.int(10)
	b := C.int(20)

	// Call C function add1
	result := C.add1(a, b)
	fmt.Printf("add1, sum of %v %v is %v\n", a, b, result) // 30
}

// Example of call C macro
// C macros aren't supported yet, we must wrap it
func callAdd2() {
	a := C.int(10)
	b := C.int(20)

	// Call C function add2
	result := C.add2_macro_wrapper(a, b)
	fmt.Printf("add2, sum of %v %v is %v\n", a, b, result) // 30
}

func callHello1() {
	// Prepare first argument
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	// Prepare second argument
	year := C.int(2023)

	// Prepare the third argument
	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(ptr)

	// Call C function hello1
	size := C.hello1(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}

// Example of call C function, argument is struct
func callHello2() {
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2023)

	// Add prefix struct_
	g := C.struct_Greetee{
		name: name,
		year: year,
	}

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(ptr)

	// Call C function hello2
	size := C.hello2(&g, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}

func useEnum() {
	level := new(C.enum_levels)
	*level = C.low

	fmt.Println(*level) // 10

	fmt.Println("C.low =", C.low)
	fmt.Println("C.medium =", C.medium)
	fmt.Println("C.high =", C.high)

	security := new(C.security)
	*security = C.MEDIUM

	fmt.Println(*security) // 1

	fmt.Println("C.LOW =", C.LOW)
	fmt.Println("C.MEDIUM =", C.MEDIUM)
	fmt.Println("C.HIGH =", C.HIGH)
}

// Example of call C function with variable arguments
// Variable argument methods like printf aren't supported yet, we must wrap it
func callAverage() {
	a := C.int(10)
	b := C.int(20)
	c := C.int(30)

	// Call C function add2
	result := C.average2_wrapper(a, b)
	fmt.Printf("average of %v %v is %v\n", a, b, result) // 15

	result = C.average3_wrapper(a, b, c)
	fmt.Printf("average of %v %v %v is %v\n", a, b, c, result) // 20
}

// Example of call C function with array arguments
func callAverageArray() {
	a := []int {1, 2, 3, 4, 5}

	// In Golang, int is int64, But in C, int is int32.
	// So, we must convert int slice to int32 slice
	aInt32 := make([]int32, len(a))
	for i, v := range a {
		aInt32[i] = int32(v)
	}

	// Call C function average_array.
	// In C, a function argument written as a fixed size array actually requires a pointer to the first element of the array.
	// C compilers are aware of this calling convention and adjust the call accordingly, but Go cannot.
	// In Go, you must pass the pointer to the first element explicitly
	result := C.average_array(C.int(len(a)), (*C.int)(&aInt32[0]))
	fmt.Printf("average of %v is %v\n", a, result) // 3
}

func main() {
	// Simple case
	callAdd1()

	// Call macro
	callAdd2()

	// Send string (using C.CString), get string (using C.GoBytes)
	callHello1()

	// Struct argument
	callHello2()

	// Enum
	useEnum()

	// Variable argument
	callAverage()

	// Array argument
	callAverageArray()
}
