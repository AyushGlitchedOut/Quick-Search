package services

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/pkg/browser"
	"log"
	"net/url"
	"os"
)

// TextTruncate : A Function to Truncate the text provided to a certain size and then insert "..."
func TextTruncate(Maxsize int, text string) string {
	//Convert the string into array or runes
	var textRunes = []rune(text)
	//Get the Rune[] length
	var textLength = len(textRunes)
	// If te provided text is smaller than limit, return as it is
	if textLength < Maxsize-3 {
		return text
	}
	//Slice the string till Maxsize-3 and add "..."
	truncatedString := string(textRunes[:Maxsize-3]) + "..."
	//Return it
	return truncatedString
}

// ExecuteQuery TODO: Make the functionality later
// ExecuteQuery: Execute the Query from the input given as parameter
func ExecuteQuery(input *gtk.SearchEntry) {
	text, err := input.GetText()
	if err != nil {
		log.Fatal("Error getting Text from Search Query", err)

	}

	//check if the input is empty
	if text == "" {
		return

	}

	//url encode the input from user
	website := "https://google.com/search?q="
	URL := url.QueryEscape(text)

	//open the encoded url in browser
	err = browser.OpenURL(website + URL)
	if err != nil {
		log.Fatal("Error Opening your Browser!!", err)
	}

	//quit the app
	os.Exit(0)
}
