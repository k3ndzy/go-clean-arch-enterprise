package main

import (
	"log"

	"github.com/k3ndzy/go-clean-arch-enterprise/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Lỗi nghiêm trọng khi khởi chạy ứng dụng: %v", err)
	}
}
