#include <stdio.h>
#include <stdlib.h>

struct Node {
    struct Node* next;
    int data;
};

struct List {
    struct Node* head;
    struct Node* tail;
    int length;
};

void print_list(struct List* list) {
    struct Node* n = list->head;
    while (n != NULL) {
        printf("%d ", n->data);
        n = n->next;
    }
    printf("\n");
}

void add_node(struct List* l, int data) {
    struct Node* n = (struct Node*)malloc(sizeof(struct Node));
    n->data = data;

    if (l->head == NULL) {
        l->head = n;
        l->tail = n;
        return;
    }

    l->tail->next = n;
    l->tail = n;
}

int main() {
    struct List* l = (struct List*)malloc(sizeof(struct List));

    add_node(l, 10);
    add_node(l, 20);
    add_node(l, 21);

    // printf("%d", l->head->data);

    print_list(l);

    return 0;
}
