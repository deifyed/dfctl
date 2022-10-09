.PHONY=fmt
INSTALL_DIR=~/.local/bin
BUILD_DIR=./build
BINARY_NAME=infect

fmt:
	@gofmt -w .

test:
	go test -v ./...

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

install:
	test -f $(BUILD_DIR)/$(BINARY_NAME)
	mkdir -p $(INSTALL_DIR)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)

uninstall:
	rm $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	@rm -rf $(BUILD_DIR)
