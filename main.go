package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
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


// Lexer
func lex(source string) (tokens []Lexeme) {
    src := []rune(source)

    var lexeme Lexeme
    for i := 0; i < len(src); i++ {
        if src[i] == ' ' {
            continue
        }

        lexeme = append(lexeme, src[i])

        if i+1 == len(src) ||
           !unicode.IsDigit(src[i]) ||
           !unicode.IsDigit(src[i+1]) {
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

