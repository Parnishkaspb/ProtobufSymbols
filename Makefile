TOOL_NAME=protosym
SRC_DIR=./cmd
OUTPUT_DIR=bin

build: build-linux build-darwin build-windows

build-linux:
	@echo "Сборка под Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(TOOL_NAME)-linux $(SRC_DIR)

build-darwin:
	@echo "Сборка под macOS..."
	go build -o $(OUTPUT_DIR)/$(TOOL_NAME) $(SRC_DIR)

build-windows:
	@echo "Сборка под Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(TOOL_NAME).exe $(SRC_DIR)

# === Прогон по mydir/*.proto ===
test_files:
	@echo "Сборка под текущую ОС..."
	go build -o $(OUTPUT_DIR)/$(TOOL_NAME) $(SRC_DIR)
	@echo "Прогон по файлам из mydir/"
	@for file in mydir/*.proto; do \
		echo "→ $$file"; \
		./$(OUTPUT_DIR)/$(TOOL_NAME) $$file; \
	done