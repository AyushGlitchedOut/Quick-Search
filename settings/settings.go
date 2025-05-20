package main

import (
	"embed"
	"github.com/AyushGlitchedOut/Quick-Search/services"
	"github.com/gotk3/gotk3/gtk"

	"log"
)

//go:embed SettingsStyles.css
var assets embed.FS

func main() {

	//Get the Style from SettingsStyles.css file (not the on in assets, the one created in the folder)
	StyleData := services.StyleReader(assets, "SettingsStyles.css")
	//Get the provider
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Error Creating Css Provider!", err)
	}
	//Load the Styles
	err = cssProvider.LoadFromData(StyleData)
	if err != nil {
		log.Fatal("Error Loading Styles!", err)
	}

	//Start the gtk process
	gtk.Init(nil)

	//create a new Window
	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Error Creating Settings Window:", err)
	}

	//Add the Provider to the screen
	gtk.AddProviderForScreen(window.GetScreen(), cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	//Quit the Mainloop when closed
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//variable to store the state of the current page
	var page gtk.IWidget
	//Default page is appearance Page
	page = _AppearancePage()

	//Layout Box
	layout, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	layout.PackStart(_SideBar(&page), false, false, 0)
	layout.PackStart(page, true, true, 0)

	//Add components
	window.Add(layout)

	//Set the Size and Show
	window.SetSizeRequest(1200, 800)
	window.ShowAll()

	//Start the Main loop
	gtk.Main()

}
