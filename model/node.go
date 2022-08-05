package model

import (
    "fmt"
    "strings"
)

type Node struct {
    Children []*Node
    Content  []rune
}

func (n *Node) Print(depth int) {
    fmt.Println(strings.Repeat(" ", depth*4)+"n:", string(n.Content))
    for _, c := range n.Children {
        c.Print(depth + 1)
    }
}
