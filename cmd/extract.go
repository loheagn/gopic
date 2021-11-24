package cmd

import (
	"strings"
	"unicode"
)

type MDImage struct {
	altText string
	path    string
	title   string
}

func extractMDImages(text string) []*MDImage {
	type ImageState int
	const (
		start ImageState = iota + 1

		imageStart

		altTextStart
		altTextEscape
		altTextEnd

		pathStart
		pathEscape
		pathEnd

		titleStart
		titleEscape
		titleEnd
	)

	const (
		singleQuote = iota + 1
		doubleQuote
	)

	var (
		imageList  []*MDImage
		state      = start
		quoteState = 0
		altText    []rune
		path       []rune
		title      []rune
	)

	endAction := func() {
		imageList = append(imageList, &MDImage{
			altText: string(altText),
			path:    strings.TrimSpace(string(path)),
			title:   string(title),
		})
		state = start
	}

	escapeAndPush := func(slice []rune, ch, escapeCH rune) []rune {
		if ch == escapeCH {
			slice = slice[:len(slice)-1]
		}
		slice = append(slice, ch)
		return slice
	}

	curQuoteChar := func() rune {
		if quoteState == singleQuote {
			return '\''
		}
		return '"'
	}

	checkTitle := func(rest string) bool {
		i := -1
		target := curQuoteChar()
		for len(rest) > 0 && i < 0 {
			i = strings.Index(rest, string(target))
			// escape? which should be ignored
			if i >= 1 && rest[i-1] == '\\' {
				rest, i = rest[i+1:], -1
			}
		}
		if len(rest) <= 0 || i < 0 {
			return false
		}
		for _, ch := range rest[i+1:] {
			if unicode.IsSpace(ch) {
				continue
			}
			if ch == ')' {
				return true
			}
			return false
		}
		return false
	}

	for i, ch := range text {
		switch state {

		case start:
			if ch == '!' {
				state = imageStart
			}

		case imageStart:
			if ch == '[' {
				state = altTextStart
				altText = make([]rune, 0)
				path = make([]rune, 0)
				title = make([]rune, 0)
			}

		case altTextStart:
			switch ch {
			case ']':
				state = altTextEnd
			case '\\':
				state = altTextEscape
				fallthrough
			default:
				altText = append(altText, ch)
			}

		case altTextEscape:
			altText = escapeAndPush(altText, ch, ']')
			state = altTextStart

		case altTextEnd:
			if ch == '(' {
				state = pathStart
			} else {
				state = start
			}

		case pathStart:
			if ch == ')' {
				endAction()
				break
			} else if unicode.IsSpace(ch) {
				state = pathEnd
			} else if ch == '\\' {
				state = pathEscape
			}
			path = append(path, ch)

		case pathEscape:
			path = escapeAndPush(path, ch, ')')
			state = pathStart

		case pathEnd:
			if ch == '\'' {
				quoteState = singleQuote
				if checkTitle(text[i+1:]) {
					state = titleStart
				} else {
					path = append(path, ch)
				}
			} else if ch == '"' {
				quoteState = doubleQuote
				if checkTitle(text[i+1:]) {
					state = titleStart
				} else {
					path = append(path, ch)
				}
			} else if ch == ')' {
				endAction()
			} else {
				path = append(path, ch)
				state = pathStart
			}

		case titleStart:
			if ch == '\'' && quoteState == singleQuote ||
				ch == '"' && quoteState == doubleQuote {
				state = titleEnd
			} else if ch == '\\' {
				state = titleEscape
				title = append(title, ch)
			} else {
				title = append(title, ch)
			}

		case titleEscape:
			title = escapeAndPush(title, ch, curQuoteChar())
			state = titleStart

		case titleEnd:
			if ch == ')' {
				endAction()
			} else if !unicode.IsSpace(ch) {
				state = start
			}
		}
	}
	return imageList
}
