package main

import (
	"fmt"
	"infix_postfix/matcher"
	"infix_postfix/model"
	"unicode"
)

func main() {
	infix := "  34 / 68 - 23 + 93 * ( 24   / 2  )-(4+ 38)    "
	expectedPostfix := "34 68 / 23 - 93 24 2 / * + 4 38 + -"

	tokens := lex([]rune(infix))
	ast := parse(tokens)
	rpn := emit(ast)

	for _, t := range tokens {
		fmt.Println("t: " + string(t))
	}
	ast.Print(0)

	fmt.Printf("%s\n%s\n%s\n", infix, rpn, expectedPostfix)
}

// Lexer
func lex(src []rune) (tokens []model.Lexeme) {
	var lexeme model.Lexeme
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

// Parser
func parse(tokens []model.Lexeme) *model.Node {
	ctx := model.Context{
		Tokens: tokens,
		Index:  0,
	}

	return matcher.Additive(&ctx)
}

// RPN emitter
func emit(n *model.Node) (rpn string) {
	if len(n.Children) >= 1 {
		rpn += emit(n.Children[0])
	}
	if len(n.Children) >= 2 {
		rpn += emit(n.Children[1])
	}
	return rpn + string(n.Content) + " "
}
