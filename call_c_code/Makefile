.PHONY: all c go run clean

all: c go run

c:
	make -C csrc

go:
	# 静态连接
	CGO_LDFLAGS='./csrc/libtest.a' go build -o main.1 main.go
	# 动态连接
	CGO_LDFLAGS='-L./csrc -ltest' go build -o main.2 main.go

run:
	./main.1
	# 动态连接运行
	LD_LIBRARY_PATH=./csrc DYLD_LIBRARY_PATH=./csrc ./main.2

clean:
	make -C csrc clean
	rm -f main.1 main.2
