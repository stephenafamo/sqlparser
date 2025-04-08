# SQL Parser

This is a Go package to parse SQL. The parsers are generated using [ANTLR](https://github.com/antlr/antlr4)
with the grammar mostly copied from [Grammars-v4](https://github.com/antlr/grammars-v4)

## About

This parser is maintained primarily to be used in Bob (<https://github.com/stephenafamo/bob>) to parse SQL queries.

As more people use **Bob**, this parser is constantly being improved and more edge cases are being handled.  
However, this means that the `SELECT`, `INSERT`, `UPDATE`, and `DELETE` statements are the most tested and used. As such, the parser is not guaranteed to be 100% accurate for all SQL queries.
