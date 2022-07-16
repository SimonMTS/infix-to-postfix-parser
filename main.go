package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)


func main() {
    // tokens := lex("  123 +  354 9  *   76")
    tokens := lex("34 + 56 + 9")

    for _, t := range tokens {
        fmt.Println("t: " + string(t))
    }

    ast := parse(tokens)

    ast.Print(0)
}

// Parser
func parse(tokens []Lexeme) (root Node) {
    root.content = "root"
    active := &root
    ctx := Context{
        tokens: tokens,
        index: 0,
    }

    matchOperator := func (ctx *Context, terminal Lexeme) *Node {
        if ctx.lookahead().equals(terminal) {
            node := Node{ make([]*Node, 0), string(ctx.lookahead())}
            // active.children = append(active.children, &node)
            ctx.advance()
            return &node
        } else {
            fmt.Println("syntax error1 | expected " +
            string(terminal) + " but found " +
            string(ctx.lookahead())); os.Exit(1)
            return nil
        }
    }


    var matchNumber func(ctx *Context) *Node
    matchNumber = func (ctx *Context) *Node {
        // is lookahead a number
        for _, r := range ctx.lookahead() {
            if !unicode.IsDigit(r) {
                fmt.Println("syntax error2"); os.Exit(1)
            }
        }

        node := Node{ make([]*Node, 0), string(ctx.lookahead())}
        // active.children = append(active.children, &node)
        ctx.advance()

        return &node
    }

    var matchExpr func (*Context)
    matchExpr = func (ctx *Context) {

        // if ctx.lookahead().equals(Lexeme("+")) {
        //     node := matchOperator(ctx, Lexeme("+"))

        //     oldActive := active
        //     active = node
        //     matchExpr(ctx)
        //     matchExpr(ctx)
        //     active = oldActive
        // } else if ctx.lookahead().isNumber() {
        //     matchNumber(ctx)
        // }

        n := matchNumber(ctx)

        if ctx.index < len(ctx.tokens)-1 {
            op := matchOperator(ctx, Lexeme("+"))
            active.children = append(active.children, op)
            oldActive := active
            active = op
            active.children = append(active.children, n)
            matchExpr(ctx)
            active = oldActive
        } else {
            active.children = append(active.children, n)
        }

    }


    _ = matchOperator
    _ = matchNumber
    _ = matchExpr

    matchExpr(&ctx)

    return
}
type Context struct {
    tokens []Lexeme
    index int
}
func (ctx *Context) lookahead() Lexeme {
    return ctx.tokens[ctx.index]
}
func (ctx *Context) advance() {
    if ctx.index < len(ctx.tokens)-1 {
        ctx.index += 1
    } else {
        fmt.Println("done")
    }
}
type Node struct {
    children []*Node
    content  string
}
func (n *Node) Print(depth int) {
    fmt.Println(strings.Repeat(" ", depth*4) + "n:",  n.content)
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

