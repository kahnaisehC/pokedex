package main

import "strings"

func cleanInput(text string) []string {
	newStrings := strings.Split(text, " ")
	for i := 0; i < len(newStrings); i++ {
		newStrings[i] = strings.ToLower(newStrings[i])
	}
	return newStrings
}
