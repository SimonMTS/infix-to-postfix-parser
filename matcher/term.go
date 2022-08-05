package matcher

import (
	"infix_postfix/model"
	r "reflect"
	"strconv"
)

func Term(ctx *model.Context) (n *model.Node) {
	isNum := func(l model.Lexeme) bool {
		_, err := strconv.ParseInt(string(l), 10, 64)
		return err == nil
	}

	if isNum(ctx.Lookahead()) {
		n = Number(ctx)
	} else if r.DeepEqual(ctx.Lookahead(), model.Lexeme("(")) {
		Parenthesis(ctx)
		n = Additive(ctx)
		Parenthesis(ctx)
	}

	return
}
