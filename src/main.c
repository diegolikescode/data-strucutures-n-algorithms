#include "Array/array.h"

int main() {
    int myArray[] = {1, 2, 3, 4, 5, 6, 7, 9};
    int size = sizeof(myArray) / sizeof(myArray[0]);

    // Call the function from array.c
    printArray(myArray, size);

    return 0;
}
