package parser

import "unicode"

type stateFn func(*lexer) stateFn

type itemType int

type item struct {
	typ itemType //The type of this item.
	val string   // The value of this item.
}

const (
	itemError      itemType = iota // error occured. Value is error text
	plainText                      //
	opener                         // {
	openQuote                      // "
	closeQuote                     // "
	closer                         // }
	itemName                       // name of the playlist
	itemId                         // id of the playlist
	separator                      // ',' with or without ' '
	aliasIndicator                 // word alias
	filePathOpener                 // "
	filePath                       // the path of the file
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

func lexOpener(l *lexer) stateFn {
	l.pos++
	l.emit(opener)
	return lexObjectArray
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

// func lexObjectOpener(l *lexer) stateFn {
// 	l.pos++
// 	l.emit(opener)
// 	return lexInsideObject
// }

func lexCloser(l *lexer) stateFn {
	l.pos++
	l.emit(closer)
	return lexInsideObject
}

func lexInsideObject(l *lexer) stateFn {
	for {
		if l.pos < len(l.input) {
			switch peekNext(l) {
			case "{":
				return lexFilesListOpener
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
			// return lexInsideObject
			// 	return lexSeparator
			case "-", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				return lexID
			}
		}
		l.pos++
		l.start++
	}
	return nil
}

func lexSeparator(l *lexer) stateFn {
	for {
		if peekNext(l) != " " && peekNext(l) != "," {
			l.emit(separator)
			return lexInsideObject
		}
		l.pos++
	}
}

func lexOpenQuote(l *lexer) stateFn {
	l.pos++
	l.emit(openQuote)
	return lexInsideString
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

func lexFilesListOpener(l *lexer) stateFn {
	l.pos++
	l.emit(opener)
	return lexInsideFilesList
}

func lexInsideFilesList(l *lexer) stateFn {
	if l.pos < len(l.input) {
		switch peekNext(l) {
		case "a":
			return lexAliasNotation
		case "\"":
			return lexFilePathCloser
		case ",":
			return lexSeparator
		case "}":
			return lexCloser
		}
	}
	return nil
}

func lexAliasNotation(l *lexer) stateFn {
	for {
		// Expecting the word 'alias' followed by a space
		if peekNext(l) == " " {
			l.emit(aliasIndicator)
			return lexFilePathOpener
		}
		l.pos++
	}
}

func lexFilePathOpener(l *lexer) stateFn {
	for {
		// Expecting the word 'alias' followed by a space
		if peekNext(l) != " " && peekNext(l) != "\"" {
			l.emit(filePathOpener)
			return lexFilePath
		}
		l.pos++
	}
	l.pos++
	l.emit(filePathOpener)
	return lexInsideString
}

func lexFilePath(l *lexer) stateFn {
	for {
		// Scanning path until closing quote.
		// TODO: Should be able to handle quotes inside filenames
		if peekNext(l) == "\"" {
			l.emit(filePath)
			return lexInsideFilesList
		}
		l.pos++
	}
}

func lexFilePathCloser(l *lexer) stateFn {
	l.pos++
	l.emit(closeQuote)
	return lexFilesListCloser
}

func lexFilesListCloser(l *lexer) stateFn {
	l.pos++
	l.emit(closer)
	return lexInsideObject
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

			l.pos++
			l.emit(closeQuote)
			return lexInsideObject
		}
		l.pos++
	}
}
