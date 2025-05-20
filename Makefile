.PHONY: run build clean settings

run:
	@-mkdir -p build &> /dev/null
	@go build -v -tags gtk_3_20 -o build/QuickSearch
	@./build/QuickSearch
build:
	@-mkdir -p build &> /dev/null
	@GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w" -tags gtk_3_20 -o build/App/QuickSearch
	@GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w" -tags gtk_3_20 -o build/App/Settings settings/settings.go
	@echo "Build Process Complete"

settings:
	@-mkdir -p build &> /dev/null
	@-rm  settings/SettingsStyles.css
	@-touch settings/SettingsStyles.css && echo "/*Auto Generated File, Don't edit! Edit the assets/SettingsStyles.css file */" >> settings/SettingsStyles.css
	@-cat assets/SettingsStyles.css >> settings/SettingsStyles.css
	@go build -v -tags gtk_3_20 -o build/Settings settings/settings.go settings/Components.go
	@./build/Settings

clean:
	@rm -rf build

