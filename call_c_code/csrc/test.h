#ifndef _TEST_H
#define _TEST_H

int add1(int a, int b);

/* Macro */
#define add2(a,b) (a)+(b)

int hello1(const char *name, int year, char *out);

struct Greetee {
    const char *name;
    int year;
};
int hello2(struct Greetee *g, char *out);

/* Enum */
enum levels {
  low = 10,
  medium = 20,
  high = 30
};

typedef enum {
	LOW = 0,
    MEDIUM = 1,
    HIGH = 2
} security;

double average(int num, ...);

double average_array(int num, int a[]);

#endif
