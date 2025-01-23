#include "linear-search.h"

#include <stdio.h>

int linear_search(int* arr, size_t arr_len, int v) {
    if (arr_len == 0) {
        return -1;
    }
    printf("searching: %d, the array: ", v);
    for (int i = 0; i < arr_len; i++) {
        printf("%d ", arr[i]);
    }

    int found = -1;
    for (int i = 0; i < arr_len; i++) {
        if (v == arr[i]) {
            found = 1;
            break;
        }
    }

    printf("FOUND VALUE: %d\n", found);
    return found;
}
