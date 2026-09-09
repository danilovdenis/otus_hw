package main

import (
	"maps"
	"slices"
	"sort"
	"strings"
)

// Простой вариант
//func Top10(str string) []string {
//	words := make(map[string]int)
//
//	fields := strings.Fields(str)
//
//	for _, word := range fields {
//		words[word] += 1
//	}
//
//	keys := slices.Collect(maps.Keys(words))
//
//	sort.Slice(keys, func(i, j int) bool {
//		if words[keys[i]] == words[keys[j]] {
//			return keys[i] < keys[j]
//		}
//		return words[keys[i]] > words[keys[j]]
//	})
//
//	if len(keys) > 10 {
//		keys = keys[:10]
//	}
//
//	return keys
//}

// Вариант со звездочкой
func Top10(str string) []string {
	words := make(map[string]int)

	fields := strings.Fields(str)

	for _, word := range fields {
		clearWord := strings.ToLower(word)
		clearWord = strings.Trim(clearWord, ".,:;()[]")

		if clearWord == "-" {
			continue
		}

		words[strings.ToLower(clearWord)] += 1
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
