## 运行
```shell
[root@hcss-ecs-67d5 call_so]# make
make -C lib
make[1]: Entering directory '/root/call_so/lib'
gcc -shared -fPIC -o libshared.so lib.c
gcc -c -Wall -Werror -fPIC -o libshared.o lib.c
ar r libshared.a libshared.o
make[1]: Leaving directory '/root/call_so/lib'
CGO_LDFLAGS='./lib/libshared.a' go build -o main.1 main.go
CGO_LDFLAGS='-L./lib -lshared' go build -o main.2 main.go
./main.1
Result of add_numbers(3, 4): 7
LD_LIBRARY_PATH=./lib DYLD_LIBRARY_PATH=./lib ./main.2
Result of add_numbers(3, 4): 7
```

## 参考资料
1. [https://pkg.go.dev/cmd/cgo](https://pkg.go.dev/cmd/cgo)