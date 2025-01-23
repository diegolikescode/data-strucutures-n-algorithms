#include "binary-search.h"

#include <stdio.h>

int binary_search(int* arr, size_t len, int v) {
    int lo = 0;
    int hi = len;

    while (lo != hi) {
        int m = lo + (hi - lo) / 2;
        if (arr[m] == v) {
            return 1;
        }

        if (v > arr[m]) {
            lo = m + 1;
        }

        if (v < arr[m]) {
            hi = m;
        }
    }
    return -1;
}
