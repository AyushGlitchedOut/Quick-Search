package services

import (
	"embed"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

// StyleReader : The Style Reader that takes the embedded FS, a filename and returns the css data in string
func StyleReader(fs embed.FS, fileName string) (css string) {
	StyleData, err := fs.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error Reading Styles!", err)
	}
	return string(StyleData)
}

// To make cursor a pointer upon hovering
func EnableHoverPointer(widget gtk.IWidget) {
	button := widget.ToWidget()
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
