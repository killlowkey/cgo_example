.PHONY: all c go run clean

all: c go run

c:
	make -C lib

go:
	@# 静态连接，这里是绝对路径
	CGO_LDFLAGS='./lib/libshared.a' go build -o main.1 main.go
	@# 动态连接，告诉搜索 libshared.dll 或 libshared.so 文件
	@# 在 -l 选项后面，传递的库名称不应包含前缀 lib，也就是说 -lshared 是正确的，而 -llibshared 是不正确的。
	@# 系统会自动假定 lib 前缀，所以只需要提供库名的后半部分
	CGO_LDFLAGS='-L./lib -lshared' go build -o main.2 main.go

run:
	./main.1
	@# 动态连接运行
	LD_LIBRARY_PATH=./lib DYLD_LIBRARY_PATH=./lib ./main.2
	@# 注释声明加载 so 文件; go run main.go
	go build -o main.3 main.go && ./main.3
	@# 动态加载 so 文件
	cd lib && go run main.go

clean:
	make -C lib clean
	rm -f main.1 main.2 main.3
