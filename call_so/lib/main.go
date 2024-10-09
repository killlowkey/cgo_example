package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	//dll, err := syscall.LoadDLL("./call_so/lib/libshared.so") // IDE 运行需要开放如下搜索路径
	dll, err := syscall.LoadDLL("libshared.so")
	if err != nil {
		panic(err)
	}

	defer dll.Release()

	// 查找函数
	proc, err := dll.FindProc("add_numbers") // 这里替换成你的函数名
	if err != nil {
		panic(err)
	}

	// 调用函数
	result, _, err := proc.Call(uintptr(1), uintptr(2))
	if err != nil {
		// 判断是否有错误
		if !errors.Is(err, syscall.Errno(0)) {
			panic(err)
		}
	}

	// 输出返回结果
	fmt.Println("Result from add_numbers:", result)
}
