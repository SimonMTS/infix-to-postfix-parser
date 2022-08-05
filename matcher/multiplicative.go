package matcher

import (
    "infix_postfix/model"
    r "reflect"
)

func Multiplicative(ctx *model.Context) *model.Node {
    root := Term(ctx)

    for r.DeepEqual(ctx.Lookahead(), model.Lexeme("*")) ||
        r.DeepEqual(ctx.Lookahead(), model.Lexeme("/")) {
        op := Operator(ctx)
        op.Children = append(op.Children, root)

        n := Term(ctx)
        op.Children = append(op.Children, n)
        root = op
    }

    return root
}
