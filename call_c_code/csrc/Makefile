.PHONY: all c-dynamic-lib c-static-lib clean

all: c-dynamic-lib c-static-lib

c-dynamic-lib:
	gcc -shared -fPIC -o libtest.so test.c

c-static-lib:
	gcc -c -Wall -Werror -fPIC -o test.o test.c
	ar r libtest.a test.o

clean:
	rm -f test.o libtest.so libtest.a
