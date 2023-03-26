package sqlparser

//go:generate antlr -Dlanguage=Go ./mysql/MySqlLexer.g4 ./mysql/MySqlParser.g4
//go:generate antlr -Dlanguage=Go ./sqlite/SQLiteLexer.g4 ./sqlite/SQLiteParser.g4
