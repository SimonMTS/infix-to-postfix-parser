package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

//  E -> E O N | N
//  O -> '+'
//  N -> 0..9
//
//  E -> E O N | N
//  A -> A ααα | β
//  ==
//  A -> βR
//  R -> αR | ε
//
//
//  E -> N R
//  R -> O N R | ε
//  O -> '+'
//  N -> 0..9

func main() {
    // tokens := lex("  123 +  354 9  *   76")
    tokens := lex("3 + 5 + 9")

    for _, t := range tokens {
        fmt.Println("t: " + string(t))
    }

    ast := parse(tokens)

    ast.Print(0)
}

// Parser
func parse(tokens []Lexeme) (root Node) {
    root.content = []rune("root")
    active := &root
    ctx := Context{
        tokens: tokens,
        index: 0,
    }

    var matchE func (*Context)
    var matchR func (*Context)

    // terminal matchers
    matchOperator := func (ctx *Context) Lexeme {
        l := ctx.lookahead()
        if l.equals(Lexeme("+")) {
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


    matchE = func (ctx *Context) {
        n := Node{ make([]*Node, 0), matchNumber(ctx)}
        active.children = append(active.children, &n)

        oldActive := active
        active = &n
        matchR(ctx)
        active = oldActive
    }

    matchR = func (ctx *Context) {
        if ctx.lookahead().equals(Lexeme("+")) {
            op := Node{ make([]*Node, 0), matchOperator(ctx)}
            active.children = append(active.children, &op)

            n := Node{ make([]*Node, 0), matchNumber(ctx)}
            active.children = append(active.children, &n)
            oldActive := active
            active = &n

            // op.children = append(active.children, leftNum)
            // op.children = append(active.children, &rightNum)

            // oldActive := active
            // active = &op
            matchR(ctx)
            active = oldActive
        } else {
            // of leeg
            // n := Node{ make([]*Node, 0), []rune("ε")}
            // active.children = append(active.children, &n)
        }
    }




    matchE(&ctx)

    _ = matchE
    _ = matchR
    _ = active
    _ = ctx
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
    // if ctx.index < len(ctx.tokens)-1 {
        ctx.index += 1
    // } else {
    //     fmt.Println("done")
    // }
}
func (ctx *Context) done() bool {
    return ctx.index >= len(ctx.tokens)
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

