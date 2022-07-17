package main

import (
    "unicode"
)

type Lexeme []rune

func lex(source string) (tokens []Lexeme) {
    src := []rune(source)

    var lexeme Lexeme
    for i := 0; i < len(src); i++ {
        if src[i] == ' ' { continue }

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
