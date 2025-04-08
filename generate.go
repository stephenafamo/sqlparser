package sqlparser

//go:generate antlr -Dlanguage=Go -visitor ./mysql/MySqlLexer.g4 ./mysql/MySqlParser.g4
//go:generate antlr-format -v ./mysql/*.g4

//go:generate antlr -Dlanguage=Go -visitor ./sqlite/SQLiteLexer.g4 ./sqlite/SQLiteParser.g4
//go:generate antlr-format -v ./sqlite/*.g4
