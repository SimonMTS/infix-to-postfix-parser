package main

import (
	"fmt"
	"os"
	"strings"
)

// Parser
func parse(tokens []Lexeme) *Node {
    ctx := Context{
        tokens: tokens,
        index: 0,
    }


    // terminal matchers
    matchOperator := func (ctx *Context) *Node {
        l := ctx.lookahead()
        if l.equals(Lexeme("+")) ||
           l.equals(Lexeme("-")) ||
           l.equals(Lexeme("*")) ||
           l.equals(Lexeme("/")) {
            ctx.advance()
            return &Node{ make([]*Node, 0), l}
        } else {
            fmt.Println("error1"); os.Exit(1); return nil
        }
    }
    matchNumber := func (ctx *Context) *Node {
        l := ctx.lookahead()
        if l.isNumber() {
            ctx.advance()
            return &Node{ make([]*Node, 0), l}
        } else {
            fmt.Println("error2"); os.Exit(1); return nil
        }
    }
    matchParentheses := func (ctx *Context) *Node {
        l := ctx.lookahead()
        if l.equals(Lexeme("(")) ||
           l.equals(Lexeme(")")) {
            ctx.advance()
            return &Node{ make([]*Node, 0), l}
        } else {
            fmt.Println("error3 " + string(l)); os.Exit(1); return nil
        }
    }
    // terminal matchers


    var matchAdditive func (*Context) *Node
    var matchMultiplicative func (*Context) *Node
    var matchTerm func (*Context) *Node

    matchAdditive = func (ctx *Context) *Node {
        root := matchMultiplicative(ctx)

        for ctx.lookahead().equals(Lexeme("+")) ||
            ctx.lookahead().equals(Lexeme("-")) {
            op := matchOperator(ctx)
            op.children = append(op.children, root)

            n := matchMultiplicative(ctx)
            op.children = append(op.children, n)
            root = op
        }

        return root
    }

    matchMultiplicative = func (ctx *Context) *Node {
        root := matchTerm(ctx)

        for ctx.lookahead().equals(Lexeme("*")) ||
            ctx.lookahead().equals(Lexeme("/")) {
            op := matchOperator(ctx)
            op.children = append(op.children, root)

            n := matchTerm(ctx)
            op.children = append(op.children, n)
            root = op
        }

        return root
    }

    matchTerm = func (ctx *Context) (n *Node) {
        if ctx.lookahead().isNumber() {
            n = matchNumber(ctx)
        } else if ctx.lookahead().equals(Lexeme("(")) {
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
