package main

import (
	"embed"
	_ "embed"
	"github.com/AyushGlitchedOut/Quick-Search/app"
)

// import and embed the assets folder
//
//go:embed assets
var assets embed.FS

func main() {

	//create the main app Search Bar
	app.CreateSearchBar(assets)
}
