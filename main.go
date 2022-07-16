package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

func main() {
    infix := "345 + 567 + 789 - 2000"

    tokens := lex(infix)
    // for _, t := range tokens {
    //     fmt.Println("t: " + string(t))
    // }

    ast := parse(tokens)
    // ast.Print(0)

    rpn := emit(ast.children[0])
    fmt.Println(infix)
    fmt.Println(rpn)
}

// RPN emitter
func emit(n *Node) (rpn string) {
    if len(n.children) >= 1 { rpn += emit(n.children[0]) }
    if len(n.children) >= 2 { rpn += emit(n.children[1]) }
    return rpn + string(n.content) + " "
}

// Parser
func parse(tokens []Lexeme) (root Node) {
    root.content = []rune("root")
    active := &root
    ctx := Context{
        tokens: tokens,
        index: 0,
    }


    // terminal matchers
    matchOperator := func (ctx *Context) Lexeme {
        l := ctx.lookahead()
        if l.equals(Lexeme("+")) ||
           l.equals(Lexeme("-")) {
            ctx.advance()
            return l
        } else {
            fmt.Println("error1"); os.Exit(1); return nil
        }
    }
    matchNumber := func (ctx *Context) Lexeme {
        l := ctx.lookahead()
        if l.isNumber() {
            ctx.advance()
            return l
        } else {
            fmt.Println("error2"); os.Exit(1); return nil
        }
    }
    // terminal matchers


    matchAdditive := func (ctx *Context) {
        localRoot := &Node{ make([]*Node, 0), matchNumber(ctx)}

        for ctx.lookahead().equals(Lexeme("+")) ||
            ctx.lookahead().equals(Lexeme("-")) {
            op1 := Node{ make([]*Node, 0), matchOperator(ctx)}
            n2 := Node{ make([]*Node, 0), matchNumber(ctx)}

            op1.children = append(op1.children, localRoot)
            op1.children = append(op1.children, &n2)
            localRoot = &op1
        }

        active.children = append(active.children, localRoot)
    }

    matchAdditive(&ctx)

    return
}
type Context struct {
    tokens []Lexeme
    index int
}
func (ctx *Context) lookahead() Lexeme {
    if ctx.index < len(ctx.tokens) {
        return ctx.tokens[ctx.index]
    } else {
        return Lexeme("EOF")
    }
}
func (ctx *Context) advance() {
    ctx.index += 1
}
type Node struct {
    children []*Node
    content  []rune
}
func (n *Node) Print(depth int) {
    fmt.Println(strings.Repeat(" ", depth*4) + "n:",  string(n.content))
    for _, c := range n.children {
        c.Print(depth+1)
    }
}


// Lexer: seems overly complex
func lex(source string) (tokens []Lexeme) {
    src := []rune(source)

    var lexeme Lexeme
    for i := 0; i < len(src); i++ {
        if src[i] == ' ' {
            lexeme = nil
            continue
        }

        lexeme = append(lexeme, src[i])

        if i == len(src)-1 || !unicode.IsDigit(src[i+1]) {
            tokens = append(tokens, lexeme)
            lexeme = nil
        }
    }

    return
}
type Lexeme []rune
func (l1 Lexeme) equals(l2 Lexeme) bool {
    if len(l1) != len(l2) {
        return false
    }

    for i, r := range l1 {
        if r != l2[i] {
            return false
        }
    }
    return true
}
func (l Lexeme) isNumber() bool {
    for _, r := range l {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}

