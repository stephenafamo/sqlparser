# SQL Parser

This is a Go package to parse SQL. The parsers are generated using [ANTLR](https://github.com/antlr/antlr4)
with the grammar mostly copied from [Grammars-v4](https://github.com/antlr/grammars-v4)

## Changes to MySQL grammar

1. Make it able to take a BIND_PARAMETER as a valid expressionAtom
1. Change the 3rd parameter of `LEAD` and `LAG` to take an expression instead of a number
1. Add alias definitions in `INSERT` statements
