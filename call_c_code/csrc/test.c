#include "test.h"
#include <stdio.h>
#include <stdarg.h>

int add1(int a, int b) {
    return a+b;
}

int hello1(const char *name, int year, char *out) {
    int n;

    n = sprintf(out, "Hello1, %s from %d!", name, year);

    return n;
}

int hello2(struct Greetee *g, char *out) {
    int n;

    n = sprintf(out, "Hello2, %s from %d!", g->name, g->year);

    return n;
}

double average(int num, ...) {
    va_list valist;
    double sum = 0.0;
    int i;

    /* initialize valist for num number of arguments */
    va_start(valist, num);

    /* access all the arguments assigned to valist */
    for (i = 0; i < num; i++) {
        sum += va_arg(valist, int);
    }

    /* clean memory reserved for valist */
    va_end(valist);

    return sum/num;
}

double average_array(int num, int a[]) {
    if (a == NULL || num == 0) {
        return 0;
    }

    double sum = 0;
    for (int i = 0; i < num; i++) {
        sum+=a[i];
    }

    return sum / num;
}

