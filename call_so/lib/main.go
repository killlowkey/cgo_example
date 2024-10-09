package main

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

// Point 结构体定义，需要与 C 中的定义保持一致
type Point struct {
	X int32
	Y int32
}

func main() {
	dll, err := syscall.LoadDLL("libshared.so")
	if err != nil {
		panic(err)
	}
	defer dll.Release()

	// 调用 add_numbers 函数
	addNumbers, err := dll.FindProc("add_numbers")
	if err != nil {
		panic(err)
	}

	result, _, err := addNumbers.Call(uintptr(1), uintptr(2))
	if err != nil && !errors.Is(err, syscall.Errno(0)) {
		panic(err)
	}

	fmt.Println("Result from add_numbers:", int(result))

	// 调用 create_point 函数
	createPoint, err := dll.FindProc("create_point")
	if err != nil {
		panic(err)
	}

	pointResult, _, err := createPoint.Call()
	if err != nil && !errors.Is(err, syscall.Errno(0)) {
		panic(err)
	}

	// 将返回值解释为 Point 结构体
	createdPoint := *(*Point)(unsafe.Pointer(&pointResult))
	fmt.Printf("Created point: (%d, %d)\n", createdPoint.X, createdPoint.Y)

	// 调用 move_point 函数
	movePoint, err := dll.FindProc("move_point")
	if err != nil {
		panic(err)
	}

	// 创建一个 Point 结构体，并获取其指针
	movePointArg := Point{X: createdPoint.X, Y: createdPoint.Y}
	movePointArgPtr := uintptr(unsafe.Pointer(&movePointArg))

	movedPointResult, _, err := movePoint.Call(
		movePointArgPtr,
		uintptr(5),
		uintptr(10),
	)
	if err != nil && !errors.Is(err, syscall.Errno(0)) {
		panic(err)
	}

	// 将返回值解释为 Point 结构体
	movedPoint := *(*Point)(unsafe.Pointer(&movedPointResult))
	fmt.Printf("Moved point: (%d, %d)\n", movedPoint.X, movedPoint.Y)
}
