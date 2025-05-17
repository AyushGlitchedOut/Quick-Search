package services

import (
	"embed"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/pkg/browser"

	//"github.com/pkg/browser"
	"log"
	"net/url"
	"os"
)

// The Style Reader that takes the embedded FS, a filename and returns the css data in string
func StyleReader(fs embed.FS, fileName string) (css string) {
	StyleData, err := fs.ReadFile("assets/" + fileName)
	if err != nil {
		log.Fatal("Error Reading Styles!", err)
	}
	return string(StyleData)
}

// The hover motion controller
func EnableHoverPointer(button *gtk.Button) {

	//connect event of cursor as pointer when it enters the button
	button.Connect("enter-notify-event", func() {
		display, err := button.GetDisplay()
		if err != nil {
			log.Fatal("Error Getting Display Info!", err)
		}

		cursor, err := gdk.CursorNewFromName(display, "pointer")
		if err != nil {
			log.Fatal("Error Changing Cursor to Pointer!", err)
		}
		window, err := button.GetWindow()
		if err != nil {
			log.Fatal("Error Getting Window Info!!", err)
		}
		window.SetCursor(cursor)
	})

	//set the cursor to default when it leaves the button
	button.Connect("leave-notify-event", func() {
		window, err := button.GetWindow()
		if err != nil {
			log.Fatal("Error Getting Window Info!!", err)
		}
		window.SetCursor(nil)
	})
}

// TODO: Make the functionality later
// Execute the Query from the input given as parameter
func ExecuteQuery(input *gtk.SearchEntry) {
	text, err := input.GetText()
	if err != nil {
		log.Fatal("Error getting Text from Search Query")
	}

	//check if the input is empty
	if text == "" {
		return
	}

	//url encode the input from user
	website := "https://google.com/search?q="
	url := url.QueryEscape(text)

	//open the encoded url in browser
	browser.OpenURL(website + url)

	//quit the app
	os.Exit(0)
}
