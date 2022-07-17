package main

import (
    "fmt"
    "os"
    r "reflect"
    "strconv"
    "strings"
)

// TODO this is a massive function
// Parser
func parse(tokens []Lexeme) *Node {
    ctx := Context{
        tokens: tokens,
        index:  0,
    }

    isNum := func(l Lexeme) bool {
        _, err := strconv.ParseInt(string(l),10,64);
        return err == nil
    }

    // terminal matchers
    matchOperator := func(ctx *Context) *Node {
        l := ctx.lookahead()
        if r.DeepEqual(l, Lexeme("+")) ||
            r.DeepEqual(l, Lexeme("-")) ||
            r.DeepEqual(l, Lexeme("*")) ||
            r.DeepEqual(l, Lexeme("/")) {
            ctx.advance()
            return &Node{make([]*Node, 0), l}
        } else {
            fmt.Println("error1")
            os.Exit(1)
            return nil
        }
    }
    matchNumber := func(ctx *Context) *Node {
        l := ctx.lookahead()
        if isNum(l) {
            ctx.advance()
            return &Node{make([]*Node, 0), l}
        } else {
            fmt.Println("error2")
            os.Exit(1)
            return nil
        }
    }
    matchParentheses := func(ctx *Context) *Node {
        l := ctx.lookahead()
        if r.DeepEqual(l, Lexeme("(")) ||
            r.DeepEqual(l, Lexeme(")")) {
            ctx.advance()
            return &Node{make([]*Node, 0), l}
        } else {
            fmt.Println("error3 " + string(l))
            os.Exit(1)
            return nil
        }
    }
    // terminal matchers

    var matchAdditive       func(*Context) *Node
    var matchMultiplicative func(*Context) *Node
    var matchTerm           func(*Context) *Node

    matchAdditive = func(ctx *Context) *Node {
        root := matchMultiplicative(ctx)

        for r.DeepEqual(ctx.lookahead(), Lexeme("+")) ||
            r.DeepEqual(ctx.lookahead(), Lexeme("-")) {
            op := matchOperator(ctx)
            op.children = append(op.children, root)

            n := matchMultiplicative(ctx)
            op.children = append(op.children, n)
            root = op
        }

        return root
    }

    matchMultiplicative = func(ctx *Context) *Node {
        root := matchTerm(ctx)

        for r.DeepEqual(ctx.lookahead(), Lexeme("*")) ||
            r.DeepEqual(ctx.lookahead(), Lexeme("/")) {
            op := matchOperator(ctx)
            op.children = append(op.children, root)

            n := matchTerm(ctx)
            op.children = append(op.children, n)
            root = op
        }

        return root
    }

    matchTerm = func(ctx *Context) (n *Node) {
        if isNum(ctx.lookahead()) {
            n = matchNumber(ctx)
        } else if r.DeepEqual(ctx.lookahead(), Lexeme("(")) {
            matchParentheses(ctx)
            n = matchAdditive(ctx)
            matchParentheses(ctx)
        }

        return
    }

    return matchAdditive(&ctx)
}

type Context struct {
    tokens []Lexeme
    index  int
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
    fmt.Println(strings.Repeat(" ", depth*4)+"n:", string(n.content))
    for _, c := range n.children {
        c.Print(depth + 1)
    }
}
