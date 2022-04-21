package sentinel

import (
	"bytes"
	"regexp"

	"golang.org/x/text/unicode/norm"
)

var (
	wordSegmenter = regexp.MustCompile(`[\pL\p{Mc}\p{Mn}\p{N}-_']+`)
)

type StopWord struct {
	stop map[string]string
}

func NewStopWord(s []string) *StopWord {
	st := make(map[string]string)

	for _, w := range s {
		st[w] = ""
	}
	return &StopWord{
		stop: st,
	}
}

func (sw *StopWord) CleanString(content string) string {
	с := removeStopWords([]byte(content), sw.stop)

	return string(с)
}

func removeStopWords(content []byte, dict map[string]string) []byte {
	var result []byte
	content = norm.NFC.Bytes(content)
	content = bytes.ToLower(content)
	words := wordSegmenter.FindAll(content, -1)
	for _, w := range words {
		if _, ok := dict[string(w)]; ok {
			result = append(result, ' ')
		} else {
			result = append(result, []byte(w)...)
			result = append(result, ' ')
		}
	}
	return result
}
