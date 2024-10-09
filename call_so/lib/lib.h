#ifndef LIB_H
#define LIB_H

#ifdef _WIN32
    #define EXPORT __declspec(dllexport)
#else
    #define EXPORT
#endif

typedef struct Point {
    int x;
    int y;
} Point;

EXPORT int add_numbers(int a, int b);

EXPORT Point create_point();

EXPORT Point move_point(Point p, int dx, int dy);

#endif // LIB_H
