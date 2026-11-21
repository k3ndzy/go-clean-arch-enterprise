# 🏛️ Go Clean Architecture Enterprise

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Architecture: Clean](https://img.shields.io/badge/Architecture-Clean%20%26%20Inversion-orange.svg)](docs/architecture/boundaries.md)
[![100 Plans Roadmap](https://img.shields.io/badge/Roadmap-100%20Plans%20Enterprise-purple.svg)](GO_CLEAN_ARCH_100_PLANS.md)

Dự án REST API cấp doanh nghiệp (Enterprise-Grade) được **xây dựng thủ công từ đầu (from scratch)** theo triết lý **Clean Architecture** và **Dependency Inversion**, tuân thủ lộ trình chuẩn mực gồm **100 kế hoạch triển khai (100 Implementation Plans)**.

---

## 📌 Tiến độ Lộ trình (Roadmap Progress)

| Giai đoạn | Nội dung | Số kế hoạch | Trạng thái |
| :--- | :--- | :---: | :---: |
| **Phase 1** | **Foundation (Nền tảng hệ thống)** | **001 – 010** | **✅ DONE (10/10)** |
| **Phase 2** | Domain and Application | 011 – 020 | ⏳ Sẵn sàng |
| **Phase 3** | PostgreSQL Persistence | 021 – 030 | ⏳ Sẵn sàng |
| **Phase 4** | Authentication & Security | 031 – 040 | ⏳ Sẵn sàng |
| **Phase 5** | REST API Endpoints | 041 – 050 | ⏳ Sẵn sàng |
| **Phase 6** | Testing & Quality Gates | 051 – 060 | ⏳ Sẵn sàng |
| **Phase 7** | Docker & CI/CD | 061 – 070 | ⏳ Sẵn sàng |
| **Phase 8** | Observability & Performance | 071 – 080 | ⏳ Sẵn sàng |
| **Phase 9** | Enterprise Features | 081 – 090 | ⏳ Sẵn sàng |
| **Phase 10** | Documentation & Release | 091 – 100 | ⏳ Sẵn sàng |

*Xem chi tiết toàn bộ lộ trình tại:* [GO_CLEAN_ARCH_100_PLANS.md](GO_CLEAN_ARCH_100_PLANS.md)

---

## 🌟 Các tính năng đã hoàn thành trong Phase 1

1. **Plan 001:** Khởi tạo module Go (`go.mod`), cấu trúc thư mục chuẩn và điểm chạy `main.go`.
2. **Plan 002:** Xác lập ranh giới kiến trúc và quy tắc phân tầng nghiêm ngặt ([docs/architecture/boundaries.md](docs/architecture/boundaries.md)).
3. **Plan 003:** Hệ thống cấu hình (`internal/config`) hỗ trợ nạp biến môi trường, xác thực giá trị và thiết lập mặc định an toàn.
4. **Plan 004:** Cơ chế khởi động ứng dụng (`internal/app`) ráp nối rõ ràng các thành phần phụ thuộc.
5. **Plan 005:** HTTP Server hỗ trợ cấu hình timeout đầy đủ và cơ chế **Graceful Shutdown** khi nhận tín hiệu `SIGINT`/`SIGTERM`.
6. **Plan 006:** Endpoints giám sát sức khỏe chuẩn Kubernetes (`/livez` cho Liveness và `/readyz` cho Readiness).
7. **Plan 007:** Structured Logger dạng JSON (`log/slog`) tích hợp middleware tự động **che giấu thông tin nhạy cảm (Secret Redaction)** như mật khẩu, token, authorization header.
8. **Plan 008:** Chuẩn hóa cấu trúc phản hồi lỗi API theo tiêu chuẩn quốc tế **RFC 7807 (Problem Details for HTTP APIs)**.
9. **Plan 009:** Developer `Makefile` hỗ trợ các lệnh build, run, test, lint, và tính toán độ phủ code coverage.
10. **Plan 010:** Khởi tạo tài liệu dự án, nhật ký thay đổi và bảng theo dõi tiến độ chi tiết.

---

## ⚡ Bắt đầu Nhanh (Quick Start)

### Yêu cầu
* Go 1.24 trở lên

### 1. Chạy Automated Tests
```bash
go test -v -race ./...
```

### 2. Kiểm tra định dạng & Linting
```bash
gofmt -w .
go vet ./...
```

### 3. Khởi chạy Ứng dụng
```bash
go run ./cmd/api
```
Server sẽ chạy tại `http://localhost:8080`.

Kiểm tra trạng thái:
* `curl http://localhost:8080/livez`
* `curl http://localhost:8080/readyz`

---

## 🏛️ Cấu trúc Thư mục

```text
go-clean-arch-enterprise/
├── cmd/
│   └── api/
│       └── main.go                  # Điểm khởi chạy chính
├── internal/
│   ├── app/                         # App bootstrapper & Dependency Injection
│   │   └── app.go
│   ├── config/                      # Quản lý cấu hình & biến môi trường
│   │   ├── config.go
│   │   └── config_test.go
│   └── delivery/
│       └── http/                    # Tầng giao tiếp HTTP
│           ├── handler/             # Route handlers (/livez, /readyz)
│           ├── middleware/          # RequestID & Structured Logger
│           └── server.go            # HTTP Server với Graceful Shutdown
├── pkg/
│   ├── logger/                      # Slog JSON Logger với Secret Redaction
│   └── problem/                     # Chuẩn hóa lỗi RFC 7807
├── docs/
│   └── architecture/                # Tài liệu ranh giới kiến trúc
├── Makefile
├── GO_CLEAN_ARCH_100_PLANS.md       # 100 kế hoạch chi tiết
├── CHANGELOG.md
└── README.md
```

---

## 📝 Giấy phép
Dự án được phát hành theo giấy phép MIT.
