#include "lib.h"

int add_numbers(int a, int b) {
    return a + b;
}

Point create_point() {
    Point p = {10, 10};
    return p;
}

Point move_point(Point p, int dx, int dy) {
    p.x += dx;
    p.y += dy;
    return p;
}