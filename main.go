package main

import (
    "fmt"
)

func main() {
    infix := "  34 / 68 - 23 + 93 * ( 24   / 2  )-(4+ 38)    "

    tokens := lex(infix)
    // for _, t := range tokens {
    //     fmt.Println("t: " + string(t))
    // }

    ast := parse(tokens)
    // ast.Print(0)

    rpn := emit(ast)
    fmt.Println(infix)
    fmt.Println(rpn)
    fmt.Println("34 68 / 23 - 93 24 2 / * + 4 38 + -")
}

// RPN emitter
func emit(n *Node) (rpn string) {
    if len(n.children) >= 1 { rpn += emit(n.children[0]) }
    if len(n.children) >= 2 { rpn += emit(n.children[1]) }
    return rpn + string(n.content) + " "
}

