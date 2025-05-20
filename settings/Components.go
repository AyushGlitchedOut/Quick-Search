package main

import (
	"github.com/AyushGlitchedOut/Quick-Search/services"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func _SideBar(pageState *gtk.IWidget) *gtk.Box {
	Box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Error Creating Box:", err)
	}

	//Styling
	style, _ := Box.GetStyleContext()
	style.AddClass("sidebar")

	Box.SetSizeRequest(300, -1)
	Box.Add(_SideBarButton("preferences-desktop-appearance-symbolic", "Appearance", func() {
		*pageState = _AppearancePage()

	}))
	Box.Add(_SideBarButton("system-run-symbolic", "Functionality", func() {
		*pageState = _FunctionalityPage()

	}))
	Box.Add(_SideBarButton("list-add-symbolic", "Custom Scripts", func() {
		*pageState = _CustomScriptsPage()
	}))
	Box.Add(_SideBarButton("help-about-symbolic", "About", func() {
		*pageState = _AboutPage()
	}))
	return Box
}

func _SideBarButton(iconName string, text string, changePage func()) *gtk.Button {
	Button, err := gtk.ButtonNew()
	if err != nil {
		log.Fatal("Error Creating Button:", err)
	}

	//Layout
	LayoutBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		log.Fatal("Error Creating LayoutBox:", err)
	}

	//Icon
	icon, _ := gtk.ImageNew()
	icon.SetFromIconName(iconName, gtk.ICON_SIZE_BUTTON)
	LayoutBox.Add(icon)

	//Label
	label, err := gtk.LabelNew(text)
	if err != nil {
		log.Fatal("Error Creating Label:", err)
	}
	LayoutBox.Add(label)

	//enable hovering
	services.EnableHoverPointer(Button)

	//styling
	style, _ := Button.GetStyleContext()
	style.AddClass("sidebarButton")

	Button.Connect("clicked", changePage)

	//Adding Layout
	Button.Add(LayoutBox)

	return Button
}

func _AppearancePage() *gtk.Box {
	page, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Error Creating Page:", err)
	}

	//Label
	label, _ := gtk.LabelNew("Appearance LOL")

	//styling
	style, _ := page.GetStyleContext()
	style.AddClass("page")

	page.Add(label)

	return page
}
func _FunctionalityPage() *gtk.Box {
	page, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Error Creating Page:", err)
	}

	//styling
	style, _ := page.GetStyleContext()
	style.AddClass("page")

	return page
}
func _CustomScriptsPage() *gtk.Box {
	page, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Error Creating Page:", err)
	}

	//styling
	style, _ := page.GetStyleContext()
	style.AddClass("page")

	return page
}
func _AboutPage() *gtk.Box {
	page, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Error Creating Page:", err)
	}

	//styling
	style, _ := page.GetStyleContext()
	style.AddClass("page")

	return page
}
