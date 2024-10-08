#ifndef LIB_H
#define LIB_H

#ifdef _WIN32
    #define EXPORT __declspec(dllexport)
#else
    #define EXPORT
#endif

EXPORT int add_numbers(int a, int b);

#endif // LIB_H
