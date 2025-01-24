#include <stdio.h>
#include <stdlib.h>

#include "./algorithms/binary-search/binary-search.h"
#include "./algorithms/crystal-ball/crystal-ball.h"
#include "./algorithms/linear-search/linear-search.h"
#include "./array/array.h"

void test_array() {
    Array* arr = init_array(11);
    push_array(arr, 6, 1, 2, 3, 4, 5, 6);
    push_array(arr, 5, 21, 22, 23, 24, 26);

    Array* arr2 = init_array(8);
    push_array(arr2, 2, 200, 300);

    Array* a = concat_array(arr, arr2);

    print_array(a);
    free(arr);
    free(arr2);
}

void test_linear_search() {
    int a1[] = {1, 2, 3, 54};
    int a2[] = {600, 210, 11, 2, 69, 420, 1};

    size_t a1len = sizeof(a1) / sizeof(a1[0]);
    size_t a2len = sizeof(a2) / sizeof(a1[0]);
    linear_search(a1, a1len, 3);
    linear_search(a2, a2len, 420);
}

void test_binary_search() {
    int a1[] = {1, 2, 4, 8, 16, 32, 64, 128, 256};
    int a2[] = {1, 2, 6, 21, 31, 49, 50, 51, 55, 66, 77};

    size_t a1len = sizeof(a1) / sizeof(a1[0]);
    size_t a2len = sizeof(a2) / sizeof(a1[0]);
    int r = binary_search(a1, a1len, 8);
    int r2 = binary_search(a2, a2len, 50);
    int r3 = binary_search(a2, a2len, 55);
    int r4 = binary_search(a2, a2len, 69);
    int r5 = binary_search(a2, a2len, 3);
    int r6 = binary_search(a2, a2len, 1);
    int r7 = binary_search(a2, a2len, 77);
    printf("r1 (find): %d\n", r);
    printf("r2 (find): %d\n", r2);
    printf("r3 (find): %d\n", r3);
    printf("r4 (not_find): %d\n", r4);
    printf("r5 (not_find): %d\n", r5);
    printf("r6 (find): %d\n", r6);
    printf("r7 (find): %d\n", r7);
}

void test_crystal_ball() {
    int* ar = (int*)malloc(100 * sizeof(int));
    for (int i = 0; i < 100; i++) {
        if (i < 67) {
            ar[i] = 0;
        } else {
            ar[i] = 1;
        }
    }

    size_t l = 100;
    int r = find_index_crystal_ball_breaks(ar, l);
}

int main() {
    // test_array();
    // test_linear_search();
    // test_binary_search();
    test_crystal_ball();

    return 0;
}
