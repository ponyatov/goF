src/lexer/FORTH.go: src/lexer/FORTH.ragel
	ragel -Z -o $@ $<
