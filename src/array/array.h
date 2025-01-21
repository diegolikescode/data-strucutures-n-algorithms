#ifndef ARRAY_H
#define ARRAY_H

#include <stdio.h>

typedef struct array {
    int capacity;
    int length;
    int* data;
} Array;

void print_array(Array* array);

// initialize the array with 8 positions
Array* init_array(int capacity);

// add an item to the end of the array
void push_array(Array* arr, int value, ...);

// removes and returns the last item of the Array
int pop_array(Array* arr);

// removes and returns the first item of the Array
int shift_array(Array* arr);

// concats 2 arrays
Array* concat_array(Array* arr1, Array* arr2);

#endif
