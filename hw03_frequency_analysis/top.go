package hw03frequencyanalysis

import (
	"maps"
	"slices"
	"sort"
	"strings"
	"unicode"
)

func Top10(str string, advanced bool) []string {
	words := make(map[string]int)

	fields := strings.Fields(str)

	for _, word := range fields {
		if advanced {
			word = clearWord(word)

			if word == "-" {
				continue
			}
		}

		words[word]++
	}

	keys := slices.Collect(maps.Keys(words))

	sort.Slice(keys, func(i, j int) bool {
		if words[keys[i]] == words[keys[j]] {
			return keys[i] < keys[j]
		}
		return words[keys[i]] > words[keys[j]]
	})

	if len(keys) > 10 {
		keys = keys[:10]
	}

	return keys
}

func clearWord(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '-'
	})

	return word
}
