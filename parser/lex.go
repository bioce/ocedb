package parser

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var (
	cIgnore   = " \n`"
	cSingle   = ",;*={}()."
	cIdentity = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cNumber   = "0123456789"
)

type Lexer struct {
	sql string
	Token Token
	Tokens []Token
	strs []string
	cur int
	str string
}

func NewLexer(sql string) (Lexer, error) {
	lex := Lexer{
		sql:   sql,
		Token: 0,
		cur:   0,
		str:   "",
		Tokens:[]Token{},
		strs:	[]string{},
	}
	strs, err := Direct(lex.sql)
	if err != nil {
		return lex, err
	}
	for i, str := range strs {
		tok := GetToken(str)
		lex.Tokens = append(lex.Tokens, tok)
		lex.strs = append(lex.strs, str)
		log.Printf("[token] %d: %10s  -> %2d  %s", i, str, tok, tokenMapping[tok])
	}
	return lex, nil
}

func (lex *Lexer) Next() (tok Token, str string, err error) {
	if lex.cur >= len(lex.Tokens) {
		return EOF, "", nil
	}
	tok = lex.Tokens[lex.cur]
	str = lex.strs[lex.cur]
	log.Printf("[lex] tok: %2d str: %10s - %s", tok, str, tokenMapping[tok])
	lex.cur++
	return
}

func (lex *Lexer) Back() () {
	lex.cur--
}

func Direct(s string) (tokens []string, err error) {
	length := len(s)
	if length == 0 {
		return tokens, nil
	}
	cur := 0
	for {
		if cur >= len(s) {
			break
		}
		ch := s[cur]
		if ' ' == ch {
			cur += 1
			continue
		} else if '\'' == ch {
			returnCur, err := readText(s, '\'', cur+1, cur+1)
			if err != nil {
				return tokens, err
			}
			tokens = append(tokens, s[cur+1:returnCur])
			cur = returnCur + 1
			continue
		} else if '"' == ch {
			returnCur, err := readText(s, '"', cur+1, cur+1)
			if err != nil {
				return tokens, err
			}
			tokens = append(tokens, s[cur+1:returnCur])
			cur = returnCur + 1
			continue
		} else if '<' == ch {
			// 可能是1个或2个字符
			if cur+1 >= length {
				tokens = append(tokens, "<")
				return tokens, nil
			} else if '>' == s[cur+1] {
				tokens = append(tokens, "<>")
				cur += 2
				continue
			} else {
				tokens = append(tokens, "<")
				cur += 1
				continue
			}
		} else if '>' == ch {
			if cur+1 >= length {
				tokens = append(tokens, ">")
				return tokens, nil
			} else if '=' == s[cur+1] {
				tokens = append(tokens, ">=")
				cur += 2
				continue
			} else {
				tokens = append(tokens, ">")
				cur += 1
				continue
			}
			time.Sleep()
		} else if '!' == ch {
			// 必须是2个字符
			if cur+1 < length && '=' == s[cur+1] {
				tokens = append(tokens, "!=")
				cur += 2
				continue
			} else {
				return tokens, fmt.Errorf("syntax error in: %s\n", s[cur:])
			}
		} else if strings.ContainsRune(cSingle, rune(ch)) {
			tokens = append(tokens, string(ch))
			cur += 1
			continue
		} else if strings.ContainsRune(cIgnore, rune(ch)) {
			cur += 1
			continue
		} else if strings.ContainsRune(cIdentity, rune(ch)) {
			returnCur, err := readIdentity(s, cur, cur+1)
			if err != nil {
				return tokens, err
			}
			tokens = append(tokens, s[cur:returnCur])
			cur = returnCur
			continue
		} else if strings.ContainsRune(cNumber, rune(ch)) {
			returnCur, err := readNumber(s, cur, cur+1)
			if err != nil {
				return tokens, err
			}
			tokens = append(tokens, s[cur:returnCur])
			cur = returnCur
			continue
		} else {
			return tokens, fmt.Errorf("syntax error in: %s\n", s[cur:])
		}
	}
	return tokens, nil
}

func readIdentity(s string, start, cur int) (returnCur int, err error) {
	if cur >= len(s) {
		return cur, nil
	}
	if strings.ContainsRune(cIdentity, rune(s[cur])) {
		return readIdentity(s, start, cur+1)
	} else {
		return cur, nil
	}
}

func readNumber(s string, start, cur int) (returnCur int, err error) {
	if cur >= len(s) {
		return cur, nil
	}
	if '.' == s[cur] {
		// 只能有一个 .
		if strings.Contains(s[start:cur], ".") {
			return cur, fmt.Errorf("syntax error in: %s", s[start:])
		} else {
			return readNumber(s, start, cur+1)
		}
	} else if strings.ContainsRune(cNumber, rune(s[cur])) {
		return readNumber(s, start, cur+1)
	} else {
		return cur, nil
	}
}

func readText(s string, label uint8, start, cur int) (returnCur int, err error) {
	if cur >= len(s) {
		return cur, fmt.Errorf("syntax error in: %s", s[start-1:])
	}
	if s[cur] == label {
		return cur, nil
	} else {
		return readText(s, label, start, cur+1)
	}
}
