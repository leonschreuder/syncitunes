package parser

import "unicode"

type stateFn func(*lexer) stateFn

type itemType int

type item struct {
	typ itemType //The type of this item.
	val string   // The value of this item.
}

const (
	itemError         itemType = iota // error occured. Value is error text
	plainText                         //
	opener                            // {
	openQuote                         // "
	closeQuote                        // "
	closer                            // }
	itemName                          // name of the playlist
	itemId                            // id of the playlist
	separator                         // ',' with or without ' '
	aliasIndicator                    // word alias
	filePathOpener                    // "
	filePathSeparator                 // :
	filePathItem                      // One node in the directory tree
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
	for state := lexInputString; state != nil; {
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

func lexInputString(l *lexer) stateFn {
	for {
		if peekNext(l) == "{" {
			l.pos++
			l.emit(opener)
			return lexObjectArray
		}
		l.pos++
	}
	return nil
}

func lexObjectArray(l *lexer) stateFn {
	for {
		if l.pos < len(l.input) {
			switch peekNext(l) {
			case "{":
				l.pos++
				l.emit(opener)
				return lexInsideObject
			case "}":
				l.pos++
				l.emit(closer)
			case ",":
				l.pos++
				l.emit(separator)
			}
			l.pos++
			l.start++
		} else {
			break
		}
	}
	return nil
}

func lexInsideObject(l *lexer) stateFn {
	for {
		if l.pos < len(l.input) {
			switch peekNext(l) {
			case "{":
				l.pos++
				l.emit(opener)
				return lexInsideFilesList
			case "}":
				l.pos++
				l.emit(closer)
				return lexObjectArray
			case "\"":
				l.pos++
				l.emit(openQuote)
				return lexInsideString
			case ",":
				l.pos++
				l.emit(separator)
			case "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				return lexID
			}
		}
		l.pos++
		l.start++
	}
	return nil
}

func lexInsideString(l *lexer) stateFn {
	for {
		switch peekNext(l) {
		case "\"":
			l.emit(itemName)

			l.pos++
			l.emit(closeQuote)
			return lexInsideObject
		}
		l.pos++
	}
}

func lexID(l *lexer) stateFn {
	for {
		nxt := peekNext(l)
		if !(unicode.IsDigit([]rune(nxt)[0]) || nxt == "-") {
			l.emit(itemId)
			return lexInsideObject
		}
		l.pos++
	}
}

func lexInsideFilesList(l *lexer) stateFn {
	for {
		if l.pos < len(l.input) {
			switch peekNext(l) {
			case "a":
				return lexAliasNotation
			case ",":
				l.pos++
				l.emit(separator)
			case "}":
				l.pos++
				l.emit(closer)
				return lexInsideObject
			}
		}
		l.pos++
		l.start++
	}
	return nil
}

func lexAliasNotation(l *lexer) stateFn {
	for {
		switch peekNext(l) {
		case " ":
			//Lex till space, just lexed 'alias'
			l.emit(aliasIndicator)
			l.start++
		case "\"":
			l.pos++
			l.emit(filePathOpener)
			return lexFilePath
		}
		l.pos++
	}
}

func lexFilePath(l *lexer) stateFn {
	for {
		// Scanning path items until closing quote.
		// TODO: Not checked for special characters in string ( : or " in names)
		switch peekNext(l) {
		case ":":
			l.emit(filePathItem)

			l.pos++
			l.emit(filePathSeparator)
		case "\"":
			l.emit(filePathItem)

			l.pos++
			l.emit(closeQuote)
			return lexInsideFilesList
		}
		l.pos++
	}
}
