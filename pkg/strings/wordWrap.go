package strings

import (
	"strings"
)

func WordWrap(text string, lineWidth int) (wrapped [][]string) {
	words := strings.Fields(strings.TrimSpace(text))

	if len(words) == 0 {
		return wrapped
	}

	spaceLeft := lineWidth - len(words[0])
	wrapped = append(wrapped, []string{})

	for _, word := range words {
		if len(word)+1 > spaceLeft {
			spaceLeft = lineWidth - len(word)

			wrapped = append(wrapped, []string{})
		} else {
			spaceLeft -= 1 + len(word)
		}

		wrapped[len(wrapped)-1] = append(wrapped[len(wrapped)-1], word)
	}

	return wrapped

}
