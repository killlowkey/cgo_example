package main

/*
#cgo LDFLAGS: -L./lib -lshared
#include "lib/lib.h"

// 为了避免 Go 的垃圾回收器移动内存，我们需要使用这个函数
void setPointValues(struct Point* p, int x, int y) {
    p->x = x;
    p->y = y;
}
*/
import "C"
import (
	"fmt"
)

// GoPoint 是与 C 的 Point 结构体对应的 Go 结构体
type GoPoint struct {
	X int
	Y int
}

// CPoint 将 GoPoint 转换为 C.struct_Point
func (p GoPoint) CPoint() C.struct_Point {
	var cp C.struct_Point
	C.setPointValues(&cp, C.int(p.X), C.int(p.Y))
	return cp
}

// callAddNumber 调用 add_numbers 函数
func callAddNumber() {
	a, b := 3, 4
	result := C.add_numbers(C.int(a), C.int(b))
	fmt.Printf("Result of add_numbers(%d, %d): %v\n", a, b, result)
}

// callCreatePoint 调用 create_point 函数
func callCreatePoint() {
	// 调用 create_point 函数
	createdPoint := C.create_point()
	fmt.Printf("Created point: (%v, %v)\n", createdPoint.x, createdPoint.y)
}

// callMovePoint 调用 move_point 函数
func callMovePoint() {
	// 创建一个 Go 的 Point 结构体
	goPoint := GoPoint{X: 10, Y: 20}

	// 将 Go 的 Point 转换为 C 的 Point 并传递给 move_point 函数
	cPoint := goPoint.CPoint()
	movedPoint := C.move_point(cPoint, C.int(5), C.int(10))

	// 打印结果
	fmt.Printf("Original point: (%d, %d)\n", goPoint.X, goPoint.Y)
	fmt.Printf("Moved point: (%v, %v)\n", movedPoint.x, movedPoint.y)
}

func main() {
	callAddNumber()
	callCreatePoint()
	callMovePoint()
}
