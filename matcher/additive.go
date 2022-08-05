package matcher

import (
    "infix_postfix/model"
    r "reflect"
)

func Additive(ctx *model.Context) *model.Node {
    root := Multiplicative(ctx)

    for r.DeepEqual(ctx.Lookahead(), model.Lexeme("+")) ||
        r.DeepEqual(ctx.Lookahead(), model.Lexeme("-")) {
        op := Operator(ctx)
        op.Children = append(op.Children, root)

        n := Multiplicative(ctx)
        op.Children = append(op.Children, n)
        root = op
    }

    return root
}
