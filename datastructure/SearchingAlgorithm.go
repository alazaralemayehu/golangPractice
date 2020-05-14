package datastructure

import "errors"

func StringSearch(String string, needle string) (int, error) {
	index := 0

	for i := 0; i < len(String); i++ {
		if String[i] == needle[index] {
			index++
		} else {
			index = 0
		}

		if index == len(needle) {
			return i - len(needle) + 1, nil
		}
	}
	return 0, (errors.New("Nothing is found"))
}
