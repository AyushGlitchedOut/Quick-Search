package services

import (
	"embed"
	"log"
)

// The Style Reader that takes the embedded FS, a filename and returns the css data in string
func StyleReader(fs embed.FS, fileName string) (css string) {
	StyleData, err := fs.ReadFile("assets/" + fileName)
	if err != nil {
		log.Fatal("Error Reading Styles!", err)
	}
	return string(StyleData)
}
