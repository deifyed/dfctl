.PHONY=fmt
INSTALL_DIR=~/.local/bin
BUILD_DIR=./build
BINARY_NAME=infect

fmt:
	@gofmt -w .

$(BINARY_NAME):
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

build: $(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_DIR)
	mkdir -p $(BUILD_DIR)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)

uninstall: $(INSTALL_DIR)/$(BINARY_NAME)
	rm $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm $(BINARY_NAME)
