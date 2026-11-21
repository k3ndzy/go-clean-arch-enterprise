.PHONY: all build run test lint clean cover help

APP_NAME = enterprise-api
MAIN_PACKAGE = ./cmd/api
BIN_DIR = bin

all: test build

build: ## Biên dịch binary ứng dụng
	@echo "==> Đang biên dịch $(APP_NAME)..."
	go build -o $(BIN_DIR)/$(APP_NAME).exe $(MAIN_PACKAGE)

run: ## Chạy ứng dụng trực tiếp từ mã nguồn
	@echo "==> Khởi chạy $(APP_NAME)..."
	go run $(MAIN_PACKAGE)

test: ## Chạy toàn bộ automated tests
	@echo "==> Đang chạy automated tests..."
	go test -v -race ./...

cover: ## Chạy tests và xuất báo cáo độ phủ mã nguồn (coverage)
	@echo "==> Đang tính toán độ phủ test..."
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint: ## Kiểm tra định dạng và phân tích tĩnh (Static analysis)
	@echo "==> Chạy go fmt và go vet..."
	gofmt -w .
	go vet ./...

clean: ## Dọn dẹp thư mục bin và các file tạm
	@echo "==> Dọn dẹp..."
	@rmdir /s /q $(BIN_DIR) 2>nul || exit 0
	@del /f coverage.out 2>nul || exit 0

help: ## Hiển thị các lệnh hỗ trợ
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
