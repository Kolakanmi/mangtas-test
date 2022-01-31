package main

import (
	"encoding/json"
	"sort"
	"strings"
)

type Response struct {
	Word       string `json:"word"`
	Occurrence int    `json:"occurrence"`
}

func Service(input string) ([]byte, error) {
	input = strings.ToLower(input)
	wordMap := make(map[string]int)
	for _, word := range strings.Split(input, " ") {
		saved, ok := wordMap[word]
		if ok {
			wordMap[word] = saved + 1
		} else {
			wordMap[word] = 1
		}
	}

	result := make([]Response, 0, len(wordMap))

	for word, occurrence := range wordMap {
		result = append(result, Response{word, occurrence})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Occurrence > result[j].Occurrence
	})

	if len(result) > 10 {
		result = result[:10]
	}
	return json.Marshal(result)
}
