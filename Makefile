all: 
	go build timestamp-cli.go helpers.go
	mkdir -p build
	mv timestamp-cli build
	mv build/timestamp-cli build/timestamp

install:
	sudo cp build/timestamp /usr/bin
	echo "Successfully installed to /usr/bin!"