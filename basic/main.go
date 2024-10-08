package main

/*
	#include <stdio.h>
	#include <stdlib.h>
	#include <stdarg.h>
	void print() {
		printf("hellow, cgo");
	}

	int hello(const char *name, int year, char *out) {
		int n;

		n = sprintf(out, "Hello, %s from %d!", name, year);

		return n;
	}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.print()
	callHello()
}

func callHello() {
	// Prepare first argument
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	// Prepare second argument
	year := C.int(2024)

	// Prepare the third argument
	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(ptr)

	// Call C function hello
	size := C.hello(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}
