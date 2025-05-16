package app

import (
	"embed"
	"github.com/AyushGlitchedOut/Quick-Search/services"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

// Create the Main App Search Bar
func CreateSearchBar(assets embed.FS) {

	//Get the Styles from embedded AppStyles.css file and load them
	StyleData := services.StyleReader(assets, "Appstyles.css")
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Error Getting Styles:", err)
	}
	cssProvider.LoadFromData(StyleData)

	//Start Gtk
	gtk.Init(nil)

	//New Top-Level Window
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	//Set Size to 1200,60
	win.SetDefaultSize(1200, 60)
	//Remove TitleBar and Other stuff
	win.SetDecorated(false)
	//Remove window
	win.SetHasWindow(false)
	//Not resizable
	win.SetResizable(false)
	//Set Title of the application
	win.SetTitle("Quick Search")
	//!When its closed, quit gtk
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//If escape is pressed, Quit the App
	win.Connect("key-press-event", func(w *gtk.Window, event *gdk.Event) {
		keyPressed := gdk.EventKeyNewFromEvent(event)
		//If pressed key == Escape, quit
		if keyPressed.KeyVal() == gdk.KEY_Escape {
			gtk.MainQuit()
		}
	})

	//When App loads, set size to 1200,60 (done again for setting up purposes)
	win.Connect("realize", func() {
		win.SetSizeRequest(1200, 60)
	})
	//Add CSS Provider for the window
	gtk.AddProviderForScreen(win.GetScreen(), cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	//show everything
	win.ShowAll()

	//Start the Main loop
	gtk.Main()
}
