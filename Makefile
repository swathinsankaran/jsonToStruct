BIN="jsonToStruct"
BUILD_DIR="build"

exe:
	@echo "Building the binary..."
	@rm -rf $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BIN) .
	@echo "Done."

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Done."
