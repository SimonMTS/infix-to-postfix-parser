package model

type Context struct {
	Tokens []Lexeme
	Index  int
}

func (ctx *Context) Lookahead() Lexeme {
	if ctx.Index < len(ctx.Tokens) {
		return ctx.Tokens[ctx.Index]
	} else {
		return Lexeme("EOF")
	}
}

func (ctx *Context) Advance() {
	ctx.Index += 1
}
