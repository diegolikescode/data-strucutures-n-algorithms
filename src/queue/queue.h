typedef struct node {
    int value;
    struct node* next;
    struct node* prev;
    int size;
} Node;

typedef struct queue {
    int length;
    Node* head;
    Node* tail;
} Queue;

// Cria uma queue com um elemento e apontando pra NULL
Node* init(int value);

// Coloca um elemento no topo da queue
void push(Node* list, int value);

// Remove alguma coisa do topo da queue
Node* pop(Node* list);

// retorna o topo da queue
int top(Node* list);

// imprime a queue para a tela.
void print(Node* list);
