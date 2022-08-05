# infix to postfix parser

Converts expressions like:
```
  34 / 68 - 23 + 93 * ( 24   / 2  )-(4+ 38)
```
into:
```
34 68 / 23 - 93 24 2 / * + 4 38 + -
```

### Can also print intermediate states

The tokens after lexing:
```
t: -
t: 23
t: +
t: 93
t: *
t: (
t: 24
t: /
t: 2
t: )
t: -
t: (
t: 4
t: +
t: 38
t: )
```

And the AST nodes after parsing:
```
n: -
    n: +
        n: -
            n: /
                n: 34
                n: 68
            n: 23
        n: *
            n: 93
            n: /
                n: 24
                n: 2
    n: +
        n: 4
        n: 38
```

