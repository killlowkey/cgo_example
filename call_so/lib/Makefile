.PHONY: all c-dynamic-lib c-static-lib clean

all: c-dynamic-lib c-static-lib

c-dynamic-lib:
	gcc -shared -fPIC -o libshared.so lib.c

c-static-lib:
	gcc -c -Wall -Werror -fPIC -o libshared.o lib.c
	ar r libshared.a libshared.o

clean:
	rm -f libshared.o libshared.so libshared.a
