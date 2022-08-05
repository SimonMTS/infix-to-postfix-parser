package matcher

import (
    "fmt"
    "infix_postfix/model"
    "os"
    "strconv"
)

func Number(ctx *model.Context) *model.Node {
    isNum := func(l model.Lexeme) bool {
        _, err := strconv.ParseInt(string(l),10,64);
        return err == nil
    }

    l := ctx.Lookahead()
    if isNum(l) {
        ctx.Advance()
        return &model.Node{
            Children: make([]*model.Node, 0),
            Content: l,
        }
    } else {
        fmt.Println("error2")
        os.Exit(1)
        return nil
    }
}
