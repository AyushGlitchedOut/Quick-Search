package app

import (
	"embed"
	"fmt"
	"github.com/AyushGlitchedOut/Quick-Search/services"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

// CreateSearchBar : Create the Main App Search Bar
func CreateSearchBar(assets embed.FS) {
	//a boolean to track the state of the dropdown open/close to decide whether to  close upon focus-lost or not
	var IsMenuDialogOpen = false

	//Get the Styles from embedded AppStyles.css file and load them
	StyleData := services.StyleReader(assets, "assets/Appstyles.css")
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Error Getting Styles:", err)
	}
	err = cssProvider.LoadFromData(StyleData)
	if err != nil {
		log.Fatal("Error Getting Styles:", err)
	}

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

	//Enable Right-Click
	win.Connect("button-press-event", func(widget *gtk.Window, event *gdk.Event) {

		buttonPressed := gdk.EventButtonNewFromEvent(event).Button()

		if buttonPressed == gdk.BUTTON_SECONDARY {
			//TODO: Make it Open the settings menu
			fmt.Println("Right Click Pressed!!!")
		}

	})

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

	//Close the App Upon focus out
	win.Connect("focus-out-event", func() {
		//Check if the focus-out is due to dropdown being open
		if IsMenuDialogOpen {
			return
		}

		gtk.MainQuit()
	})

	//Add CSS Provider for the window
	gtk.AddProviderForScreen(win.GetScreen(), cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	//Get the Box
	mainBox := _Box()
	//Get the search Bar
	searchBar := _SearchBar()
	//Get the searchButton
	searchButton := _SearchButton()
	//Get the MenuButton
	menuButton := _MenuButton(&IsMenuDialogOpen)

	//add the searchbar and button and menuButton
	mainBox.Add(menuButton)
	mainBox.Add(searchBar)
	mainBox.Add(searchButton)

	//make it to when enter key is pressed, the ExecuteQuery() service is run
	win.Connect("key-press-event", func(w *gtk.Window, event *gdk.Event) {
		keyPressed := gdk.EventKeyNewFromEvent(event)

		//If key pressed == Enter, query
		if keyPressed.KeyVal() == gdk.KEY_Return {
			services.ExecuteQuery(searchBar)
		}
	})

	//when the button is pressed, the ExecuteQuery() service is run
	searchButton.Connect("clicked", func() {
		services.ExecuteQuery(searchBar)
	})

	//focus the searchbar upon load
	searchBar.GrabFocus()

	win.Add(mainBox)
	win.ShowAll()

	//Start the Main loop
	gtk.Main()
}

// THe box to contain both widgets
func _Box() *gtk.Box {
	//new Horizontal Box
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	//set size
	box.SetSizeRequest(1200, 60)

	//set a css style .box (not used for now)
	style, _ := box.GetStyleContext()
	style.AddClass("box")

	return box

}

// The search Entry field
func _SearchBar() *gtk.SearchEntry {
	SearchBar, err := gtk.SearchEntryNew()
	if err != nil {
		log.Fatal("Error Creating Search Entry:", err)
	}

	//Add css class .search
	style, _ := SearchBar.GetStyleContext()
	style.AddClass("search")

	//Set size
	SearchBar.SetSizeRequest(800, 40)

	//set placeholder text
	SearchBar.SetPlaceholderText("Enter your Query....")

	return SearchBar
}

// The search Button
func _SearchButton() *gtk.Button {
	button, err := gtk.ButtonNew()
	if err != nil {
		log.Fatal("Error Creating button:", err)
	}

	//Add css class .button
	style, _ := button.GetStyleContext()
	style.AddClass("button")

	//Set size
	button.SetSizeRequest(60, 40)

	//Take the search Icon from library and set it as the image of the button
	gIcon, err := glib.IconNewForString("system-search-symbolic")
	if err != nil {
		log.Fatal("Error Loading Icon:", err)
	}

	icon, err := gtk.ImageNewFromGIcon(gIcon, gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatal("Error Loading Icon:", err)
	}
	button.SetImage(icon)
	button.SetAlwaysShowImage(true)

	//Enable the hover to pointer
	services.EnableHoverPointer(button)
	return button
}
