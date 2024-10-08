# Examples of CGO
Examples of calling C code from Golang.

## Compile and run example
Run `make`:
```bash
$ make
make -C csrc
gcc -shared -fPIC -o libtest.so test.c
gcc -c -Wall -Werror -fPIC -o test.o test.c
ar r libtest.a test.o
ar: creating archive libtest.a
CGO_LDFLAGS='./csrc/libtest.a' go build -o main.1 main.go
CGO_LDFLAGS='-L./csrc -ltest' go build -o main.2 main.go
./main.1
add1, sum of 10 20 is 30
add2, sum of 10 20 is 30
Hello1, Gopher from 2023!
Hello2, Gopher from 2023!
10
C.low = 10
C.medium = 20
C.high = 30
1
C.LOW = 0
C.MEDIUM = 1
C.HIGH = 2
average of 10 20 is 15
average of 10 20 30 is 20
average of [1 2 3 4 5] is 3
LD_LIBRARY_PATH=./csrc DYLD_LIBRARY_PATH=./csrc ./main.2
add1, sum of 10 20 is 30
add2, sum of 10 20 is 30
Hello1, Gopher from 2023!
Hello2, Gopher from 2023!
10
C.low = 10
C.medium = 20
C.high = 30
1
C.LOW = 0
C.MEDIUM = 1
C.HIGH = 2
average of 10 20 is 15
average of 10 20 30 is 20
average of [1 2 3 4 5] is 3
```

See [Makefile](Makefile)

## 命令详解
```makefile
CGO_LDFLAGS='./csrc/libtest.a' go build -o main.1 main.go
```
指定了静态库文件 `libtest.a` 的路径。编译器会直接使用这个静态库文件


```makefile
CGO_LDFLAGS='-L./csrc -ltest' go build -o main.2 main.go
```
1. `-L./csrc` 告诉链接器在 `./csrc` 目录中查找库文件
2. `-ltest` 告诉链接器链接名为 `libtest` 的库（可以是静态库 .a 或动态库 .so）

```makefile
LD_LIBRARY_PATH=./csrc DYLD_LIBRARY_PATH=./csrc ./main.2
```
1. `LD_LIBRARY_PATH=./csrc`
   1. 这是设置 Linux 系统上的环境变量 
   2. `LD_LIBRARY_PATH` 告诉动态链接器在哪里查找共享库（.so 文件） 
   3. `./csrc` 表示在当前目录下的 csrc 子目录中查找
2. `DYLD_LIBRARY_PATH=./csrc`
   1. 这是设置 macOS 系统上的环境变量 
   2. `DYLD_LIBRARY_PATH` 是 macOS 上等同于 `LD_LIBRARY_PATH` 的环境变量 
   3. 同样指向 `./csrc` 目录
3. `./main.2`
   1. 这是要执行的程序名称 
   2. ./ 表示程序位于当前目录

## `.a`、`.o`、`.so` 区别
1. `.o` 文件 (目标文件)
   1. 这是编译器生成的中间文件 
   2. 包含机器码,但还不能直接执行 
   3. 每个源文件(.c 或 .cpp)通常会生成一个对应的 .o 文件 
   4. 用于后续链接过程 
2. `.a` 文件 (静态库)
   1. 是多个 .o 文件的集合 
   2. 使用 ar 工具创建 
   3. 在程序链接时,会被直接包含到最终的可执行文件中 
   4. 优点: 执行快,不依赖外部库 
   5. 缺点: 生成的可执行文件较大,更新库需要重新编译程序
3. `.so` 文件 (共享库/动态链接库)
   1. 在 Linux 系统中使用的动态链接库 (Windows 中对应 .dll)
   2. 包含能被多个程序共享的代码和数据 
   3. 在程序运行时动态加载 
   4. 优点: 节省内存,便于更新 
   5. 缺点: 可能存在版本兼容性问题,程序运行时需要依赖这些库

## file 指令
Windows 平台编译出的 so 文件，其本身就是个 dll 文件
```shell
root@main-desktop:/home/main# file libtest.so 
libtest.so: PE32+ executable (DLL) (console) x86-64, for MS Windows
```

Linux 平台编译出的 so 文件
```shell
[root@hcss-ecs-67d5 csrc]# file libtest.so 
libtest.so: ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, BuildID[sha1]=ab007e6441d243fdf388c2278a32ede9b6bc07ef, not stripped
```

# Reference
https://pkg.go.dev/cmd/cgo
