#include <stdio.h>

void hello() {
    printf("hello world\n"); 
}

// gcc -c hello.c -o hello.o
// ar rcs libhello.a hello.o