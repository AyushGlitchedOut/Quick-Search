.PHONY: run build clean

run:
	@-mkdir -p build &> /dev/null
	@go build -v -tags gtk_3_20 -o build/QuickSearch
	@./build/QuickSearch
build:
	@-mkdir -p build &> /dev/null
	@GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w" -tags gtk_3_20 -o build/App/QuickSearch
	@echo "Build Process Complete"

clean:
	@rm -rf build

