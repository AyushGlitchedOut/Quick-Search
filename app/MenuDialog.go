package app

import (
	"fmt"
	"github.com/AyushGlitchedOut/Quick-Search/services"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

// The Menu Button to Open Dialog
func _MenuButton(IsMenuDialogOpen *bool) *gtk.Button {
	menuButton := _MenuDialogLayout("Google.com", "dialog-information", false)

	//Add style class ".menuButton"
	style, _ := menuButton.GetStyleContext()
	style.AddClass("menuButton")

	//open MenuDialog when clicked
	menuButton.Connect("clicked", func() {
		//set the state tracking boolean to true before opening the dialog
		*IsMenuDialogOpen = true
		showMenuDialog(IsMenuDialogOpen)
	})

	menuButton.SetSizeRequest(280, 40)

	return menuButton
}

// The function to show the dialog
func showMenuDialog(IsMenuDialogOpen *bool) {
	// Create a dialog without parent to prevent focus loss
	dialog, _ := gtk.DialogNew()

	dialog.SetTitle("Choose an Option")
	//Remove window
	dialog.SetHasWindow(false)
	//Remove decoration like borders and stuff
	dialog.SetDecorated(false)
	//Set default size, width same as MenuButton, height 10x
	dialog.SetDefaultSize(280, 400)
	//styling
	style, _ := dialog.GetStyleContext()
	style.AddClass("dialog")

	// Prevent focus loss
	dialog.SetTransientFor(nil)
	dialog.SetKeepAbove(true)

	//Enable Right Click
	dialog.Connect("button-press-event", func(widget *gtk.Dialog, event *gdk.Event) {

		buttonPressed := gdk.EventButtonNewFromEvent(event).Button()

		if buttonPressed == gdk.BUTTON_SECONDARY {
			//TODO: Make it Open the settings menu
			fmt.Println("Right Click Pressed!!!")
		}

	})

	//When focused out, close the dialog
	dialog.Connect("focus-out-event", func() {
		dialog.Close()
	})

	//upon destroying the dialog, set the menuDialogOpen state to false
	dialog.Connect("destroy", func() {

		*IsMenuDialogOpen = false

	})

	//content area of the dialog
	contentArea, err := dialog.GetContentArea()
	if err != nil {
		log.Fatal("Unable to get content area:", err)
	}

	//ScrollWindow
	scrollWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatal("Error Creating Scroll Window in Menu", err)
	}
	scrollWindow.SetSizeRequest(280, 400)

	//ScrollWindowBox
	scrollWindowBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Error creating the scrollWindow Box", err)
	}
	//Test Item for now
	scrollWindowBox.Add(_MenuDialogLayout("Google.com", "dialog-information", true))

	//Add the Box to scrolling Window
	scrollWindow.Add(scrollWindowBox)

	//Add scrollingWindow to the dialog
	contentArea.Add(scrollWindow)

	dialog.ShowAll()
}

// The Dialog Button Component
func _MenuDialogLayout(text string, iconName string, isAnOption bool) *gtk.Button {
	//The parent Button
	Button, _ := gtk.ButtonNew()

	//Box to hold both ImageIcon and Label
	Box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		log.Fatal("Error Creating Box Layouts in Menu Dialog", err)
	}

	//Image Icon
	imageIcon, err := gtk.ImageNew()
	if err != nil {
		log.Fatal("Error Creating Icon in Menu Dialog Layout", err)
	}
	imageIcon.SetFromIconName(iconName, gtk.ICON_SIZE_BUTTON)
	Box.Add(imageIcon)

	//Label
	textLabel, err := gtk.LabelNew(services.TextTruncate(18, text))
	if err != nil {
		log.Fatal("Error Creating Label in Menu Dialog Layout", err)
	}

	//Adding style class ".menuLayout" to the main Button
	if isAnOption {
		style, _ := Button.GetStyleContext()
		style.AddClass("menuDialogLayout")
	}

	//Close the dialog if it's an Option Upon selection
	if isAnOption {
		Button.Connect("clicked", func() {
			//TODO: implement system to actually insert the selected option before closing dialog
			win, err := Button.GetToplevel()
			if err != nil {
				log.Fatal("Error getting the Window of MenuButtons")
			}
			win.ToWidget().Destroy()
		})
	}

	Box.Add(textLabel)
	Button.Add(Box)

	//Make the cursor Pointer on hovering
	services.EnableHoverPointer(Button)

	return Button
}
