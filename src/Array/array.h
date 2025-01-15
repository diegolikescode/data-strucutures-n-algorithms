#ifndef ARRAY_H
#define ARRAY_H

#include <stdio.h>

typedef struct array {
    int capacity;
    int used;
    int* data;
} Array;

void print(int* array, int size);

// initialize the array with 8 positions
Array* init();

// add an item to the end of the array
void push(Array* arr, int value);

// removes and returns the last item of the Array
int pop(Array* arr);

// removes and returns the first item of the Array
int shift(Array* arr);

// concats 2 arrays
Array* concat(Array* arr1, Array* arr2);

#endif
