build:
	@go build -o build/main cmd/api/main.go 

clean:
	@rm -rf build

run: clean build
	@./build/main
