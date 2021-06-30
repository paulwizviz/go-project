package mylib

import "fmt"

func FormSentence(words ...string) string {

	var sentence string
	for index, word := range words {

		if index == 0 {
			sentence = word
		} else if index == len(words)-1 {
			sentence = fmt.Sprintf("%s %s.", sentence, word)
		} else {
			sentence = fmt.Sprintf("%s %s", sentence, word)
		}
	}

	return sentence
}
