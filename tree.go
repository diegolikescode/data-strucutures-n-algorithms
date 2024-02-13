package main

import "fmt"

type Node struct {
    Key int
    Left *Node
    Right *Node
}

type Stack struct {
    Head *Node
}

type Queue struct {
    Head *Node
}

func main() {

}

func (n *Node) InsertNode(key int) {
    if key < n.Key {
        if n.Left == nil {
            n.Left = &Node{Key: key}
        } else {
            n.Left.InsertNode(key)
        }
    } else if key > n.Key {
        if n.Right == nil {
            n.Right = &Node{Key: key}
        } else {
            n.Right.InsertNode(key)
        }
    }
}

func (n *Node) Search(key int) bool {
    if n == nil {
        return false
    }
    if n.Key == key {
        return true
    } else if key < n.Key {
        return n.Left.Search(key)
    } else {
        return n.Right.Search(key)
    }
}

func (n *Node) DepthFirstSearch() {
    if n != nil {
        n.Left.DepthFirstSearch()
        fmt.Printf("%v", n.Key)
        n.Right.DepthFirstSearch()
    }
}

func (n *Node) BreadthFirstSearch() {
    if n == nil {
        return
    }

    queue := []*Node{n}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        fmt.Printf("%v", current.Key)

        if current.Left != nil {
            queue = append(queue, current.Left)
        }

        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }
}

