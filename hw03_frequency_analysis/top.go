package hw03frequencyanalysis

import (
	"maps"
	"slices"
	"sort"
	"strings"
	"unicode"
)

func Top10(str string, advanced bool) []string {
	// Для чистоты кода можно сделать одну функию т.к. это дублирование кода.
	// Вынести в отдельную функцию метод очистки слов и применять этот опараметр там.
	if advanced {
		return top10Advanced(str)
	}

	return top10Simple(str)
}

// Простой вариант.
func top10Simple(str string) []string {
	words := make(map[string]int)

	fields := strings.Fields(str)

	for _, word := range fields {
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

// Вариант со звездочкой.
func top10Advanced(str string) []string {
	words := make(map[string]int)

	fields := strings.Fields(str)

	for _, word := range fields {
		clearWord := strings.ToLower(word)
		clearWord = strings.TrimFunc(clearWord, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '-'
		})

		if clearWord == "-" {
			continue
		}

		words[clearWord]++
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
