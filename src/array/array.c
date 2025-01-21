#include "array.h"

#include <errno.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

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

void push_array(Array* arr, int count, ...) {
    va_list args;
    va_start(args, count);

    if ((count + arr->length) >= arr->capacity) {
        errno = EINVAL;
        perror("Index out of bounds");
        va_end(args);
        return;
    }

    for (int i = 0; i < count; i++) {
        int a = va_arg(args, int);
        arr->data[arr->length + i] = a;
    }

    arr->length += count;
    va_end(args);
}

int pop_array(Array* arr) {
    if (arr->length <= 0) {
        errno = EINVAL;
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

Array* concat_array(Array* arr1, Array* arr2) {
    Array* a = init_array(arr1->capacity + arr2->capacity);
    int len1 = arr1->length;
    int len2 = arr2->length;

    for (int i = 0; i < len1; i++) {
        a->data[a->length++] = arr1->data[i];
    }

    printf("PRINT PARTIAL ARRAY with length: %d\n", len1);
    print_array(a);

    for (int i = 0; i < len2; i++) {
        a->data[a->length++] = arr2->data[i];
    }

    return a;
}
