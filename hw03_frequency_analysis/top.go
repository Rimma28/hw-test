package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// Top10 ...
func Top10(text string) []string {
	if text == "" {
		return nil
	}
	var result []string
	dict := make(map[string]int)

	words := strings.Fields(text)
	for _, ch := range words {
		dict[ch]++
	}
	sortedStruct := make([]keyValue, 0, len(dict))
	for key, value := range dict {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}

	sortStruct(sortedStruct)
	for i := 0; i < len(sortedStruct) && i < 10; i++ {
		if i == 10 {
			break
		}
		result = append(result, sortedStruct[i].Key)
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
