package main

import (
	"fmt"
	"strings"
)

const (
	Di string = "."
	Da string = "-"
)

var alphaNumToMorse = map[string]string{
	"A": string(Di + Da),
	"B": string(Da + Di + Di + Di),
	"C": string(Da + Di + Da + Di),
	"D": string(Da + Di + Di),
	"E": string(Di),
	"F": string(Di + Di + Da + Di),
	"G": string(Da + Da + Di),
	"H": string(Di + Di + Di + Di),
	"I": string(Di + Di),
	"J": string(Di + Da + Da + Da),
	"K": string(Da + Di + Da),
	"L": string(Di + Da + Di + Di),
	"M": string(Da + Da),
	"N": string(Da + Di),
	"O": string(Da + Da + Da),
	"P": string(Di + Da + Da + Di),
	"Q": string(Da + Da + Di + Da),
	"R": string(Di + Da + Di),
	"S": string(Di + Di + Di),
	"T": string(Da),
	"U": string(Di + Di + Da),
	"V": string(Di + Di + Di + Da),
	"W": string(Di + Da + Da),
	"X": string(Da + Di + Di + Da),
	"Z": string(Da + Da + Di + Di),
	"Æ": string(Di + Da + Di + Da),
	"Ø": string(Da + Da + Di),
	"Å": string(Di + Da + Da + Di + Da),

	//numbers
	"1": string(Di + Da + Da + Da + Da),
	"2": string(Di + Di + Da + Da + Da),
	"3": string(Di + Di + Di + Da + Da),
	"4": string(Di + Di + Di + Di + Da),
	"5": string(Di + Di + Di + Di + Di),
	"6": string(Da + Di + Di + Di + Di),
	"7": string(Da + Da + Di + Di + Di),
	"8": string(Da + Da + Da + Di + Di),
	"9": string(Da + Da + Da + Da + Di),
	"0": string(Da + Da + Da + Da + Da),

	".": string(Di + Da + Di + Da + Di + Da),
	",": string(Da + Da + Di + Di + Da + Da),
	"?": string(Di + Di + Da + Da + Di + Di),
	"=": string(Da + Di + Di + Di + Da),
	"-": string(Da + Di + Di + Di + Da + Da),
	"/": string(Da + Di + Di + Da + Di),
	":": string(Da + Da + Da + Di + Di + Di),
	"+": string(Di + Da + Di + Da + Di),

	" ": string("/"),
}

var morseToAlphaNum map[string]string

func init() {
	morseToAlphaNum = make(map[string]string)
	for k, v := range alphaNumToMorse {
		morseToAlphaNum[v] = k
	}
}

// morse code to text
func decode(code string) string {

	fmt.Printf("'%s'\n", code)
	messageText := ""
	words := strings.Split(code, " / ")

	for i, word := range words {

		letters := strings.Split(word, " ")
		if i > 0 {
			messageText += " "
		}
		for _, letter := range letters {
			var char string
			var ok bool
			if char, ok = morseToAlphaNum[letter]; !ok {
				fmt.Printf("Can't match code for letter '%s'\n", letter)
			}
			messageText += char
		}
	}

	return messageText
}

// encode text to morse
func encode(text string) string {

	r := []rune(strings.ToUpper(text))

	var messageCode = ""

	for _, x := range r {

		var code string
		var ok bool
		if code, ok = alphaNumToMorse[fmt.Sprintf("%c", x)]; !ok {
			fmt.Printf("Can't match letter for code '%s'\n", fmt.Sprintf("%c", x))
		}
		messageCode += code
		messageCode += " "
	}

	return strings.TrimSpace(messageCode)
}
