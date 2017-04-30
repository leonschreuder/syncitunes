package applescript_cli

import "unicode"

type stateFn func(*lexer) stateFn

type itemType int

type item struct {
	typ itemType //The type of this item.
	val string   // The value of this item.
}

const (
	itemError  itemType = iota // error occured. Value is error text
	plainText                  //
	opener                     // {
	openQuote                  // "
	closeQuote                 // "
	closer                     // }
	itemName                   //name of the playlist
	itemId                     //name of the playlist
	separator                  // ',' with or without ' '
)

type lexer struct {
	input string    // the strig being scanned
	start int       // start position of this item
	pos   int       // current position in the input
	width int       // width of last rune read
	items chan item // channel of scanned items
}

func lex(input string) *lexer {
	l := &lexer{
		input: input,
		items: make(chan item),
	}
	go l.run() // Concurrently run state machine
	return l
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func peekNext(l *lexer) string {
	return string(l.input[l.pos])
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func lexText(l *lexer) stateFn {
	for {
		if peekNext(l) == "{" {
			l.emit(plainText)
			return lexOpener
		}
		l.pos++
	}
	return nil
}

func lexOpener(l *lexer) stateFn {
	l.pos++
	l.emit(opener)
	return lexInsideObject
}

func lexCloser(l *lexer) stateFn {
	l.pos++
	l.emit(closer)
	return lexInsideObject
}

func lexInsideObject(l *lexer) stateFn {
	if l.pos < len(l.input) {
		switch peekNext(l) {
		case "{":
			return lexOpener
		case "}":
			return lexCloser
		case "\"":
			return lexOpenQuote
		case ",":
			return lexSeparator
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			return lexID
		}
	}
	return nil
}

func lexOpenQuote(l *lexer) stateFn {
	l.pos++
	l.emit(openQuote)
	return lexInsideString
}

func lexCloseQuote(l *lexer) stateFn {
	l.pos++
	l.emit(closeQuote)
	return lexInsideObject
}

func lexInsideString(l *lexer) stateFn {
	for {
		switch peekNext(l) {
		case "\"":
			l.emit(itemName)
			return lexCloseQuote
		}
		l.pos++
	}
	return nil
}

func lexSeparator(l *lexer) stateFn {
	for {
		if peekNext(l) != " " && peekNext(l) != "," {
			// l.pos++ //we want to include the current rune
			l.emit(separator)
			return lexInsideObject
		}
		l.pos++
	}
	return nil
}

func lexID(l *lexer) stateFn {
	for {
		if !unicode.IsDigit([]rune(peekNext(l))[0]) {
			l.emit(itemId)
			return lexInsideObject
		}
		l.pos++
	}
}
