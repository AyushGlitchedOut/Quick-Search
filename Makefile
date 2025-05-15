run:
	@-mkdir -p build &> /dev/null
	@go build -v -tags gtk_3_20 -o build/QuickSearch
	@./build/QuickSearch 
clean:
	@rm -rf build

