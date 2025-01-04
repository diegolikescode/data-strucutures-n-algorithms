# DATA STRUCTURES AND ALGORITHMS

## ARRAY

``c
typedef struct {
    int *data;     // Pointer to the array data
    int size;      // Current number of elements
    int capacity;  // Maximum capacity before resizing
} DynamicArray;

// Function declarations (interface)

// Creates a new dynamic array with the specified initial capacity
DynamicArray createArray(int initialCapacity);

// Frees the memory used by the dynamic array
void freeArray(DynamicArray *array);

// Adds an element to the end of the array
void push(DynamicArray *array, int value);

// Removes and returns the last element from the array
int pop(DynamicArray *array);

// Returns the element at a specified index
int get(DynamicArray *array, int index);

// Sets the value at a specified index
void set(DynamicArray *array, int index, int value);

// Returns the current size of the array
int size(DynamicArray *array);

// Returns the current capacity of the array
int capacity(DynamicArray *array);

// Checks if the array contains a specific value
int contains(DynamicArray *array, int value);

// Finds the index of a specific value (-1 if not found)
int indexOf(DynamicArray *array, int value);

// Removes the element at a specified index
void removeAt(DynamicArray *array, int index);

// Inserts an element at a specified index
void insertAt(DynamicArray *array, int index, int value);

// Clears all elements in the array
void clear(DynamicArray *array);

// Prints the array elements (for debugging purposes)
void printArray(DynamicArray *array);
``
