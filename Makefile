SRC_FILES=timestamp-cli.go helpers.go
BUILD_DIR=build
DEFAULT_MODULE_NAME=timestamp-cli
EXECUTABLE_NAME=timestamp
INSTALL_PATH=/usr/local/bin

build: $(SRC_FILES)
	go build .
	mkdir -p $(BUILD_DIR)
	mv $(DEFAULT_MODULE_NAME) $(BUILD_DIR)
	mv $(BUILD_DIR)/$(DEFAULT_MODULE_NAME) $(BUILD_DIR)/$(EXECUTABLE_NAME)

install:
	sudo cp $(BUILD_DIR)/$(EXECUTABLE_NAME) $(INSTALL_PATH)
	@echo "Successfully installed to $(INSTALL_PATH!)"

uninstall:
	sudo mv $(INSTALL_PATH)/$(EXECUTABLE_NAME) .

all: build