package matcher

import (
    "fmt"
    "infix_postfix/model"
    "os"
    r "reflect"
)

func Operator(ctx *model.Context) *model.Node {
    l := ctx.Lookahead()
    if r.DeepEqual(l, model.Lexeme("+")) ||
        r.DeepEqual(l, model.Lexeme("-")) ||
        r.DeepEqual(l, model.Lexeme("*")) ||
        r.DeepEqual(l, model.Lexeme("/")) {
        ctx.Advance()
        return &model.Node{
            Children: make([]*model.Node, 0),
            Content: l,
        }
    } else {
        fmt.Println("error1")
        os.Exit(1)
        return nil
    }
}
