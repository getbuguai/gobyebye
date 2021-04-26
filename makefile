
build:
	@GOOS=windows GOARCH=amd64 go build -o gobyby.exe
	@GOOS=linux GOARCH=amd64 go build -o gobyby
	@GOOS=darwin GOARCH=amd64 go build -o gobyby