#include "array.h"

#include <stdio.h>
#include <stdlib.h>

// ARRAY
// int capacity;
// int length;
// int* data;

void print_array(Array* array) {
    for (int i = 0; i < array->length; i++) {
        printf("%d ", array->data[i]);
    }
    printf("\n");
}

Array* init_array(int capacity) {
    Array* arr = (Array*)malloc(sizeof(Array*));
    if (arr == NULL) {
        return NULL;
    }
    arr->capacity = capacity;
    arr->length = 0;
    arr->data = (int*)malloc(capacity * sizeof(int));
    if (arr->data == NULL) {
        return NULL;
    }

    return arr;
}

void push_array(Array* arr, int value) {
    if (arr->length >= arr->capacity) {
        perror("Index out of bounds");
        return;
    }

    arr->data[arr->length] = value;
    arr->length++;
}

int pop_array(Array* arr) {
    if (arr->length <= 0) {
        perror("Index out of bounds");
        return 0;
    }
    int n = arr->data[arr->length - 1];
    arr->data[arr->length - 1] = 0;
    arr->length--;

    return n;
}

int shift_array(Array* arr) {
    if (arr->length <= 0) {
        perror("Index out of bounds");
        return 0;
    }
    int n = arr->data[0];

    if (arr->length > 1) {
        for (int i = 1; i < arr->length; i++) {
            arr->data[i - 1] = arr->data[i];
        }
    }

    arr->length--;
    return n;
}

Array* concat_array(Array* arr1, Array* arr2) { return NULL; }
