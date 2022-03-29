package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// Top10 ...
func Top10(text string) []string {

	if text == "" {
		return make([]string, 0)
	}

	var sortedStruct []keyValue
	var result = make([]string, 10)
	var dict = make(map[string]int)

	words := strings.Fields(text)
	for _, ch := range words {
		dict[ch] = dict[ch] + 1
	}

	for key, value := range dict {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}

	sortStruct(sortedStruct)
	for i := range sortedStruct {
		if i == 10 {
			break
		}
		result[i] = sortedStruct[i].Key
	}

	return result
}

type keyValue struct {
	Key   string
	Value int
}

func sortStruct(sortedStruct []keyValue) {
	sort.Slice(sortedStruct, func(i, j int) bool {
		if sortedStruct[i].Value == sortedStruct[j].Value {
			return sortedStruct[i].Key < sortedStruct[j].Key
		}
		return sortedStruct[i].Value > sortedStruct[j].Value
	})
}
