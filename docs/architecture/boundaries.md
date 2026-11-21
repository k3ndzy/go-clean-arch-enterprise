# Architecture Boundaries & Dependency Inversion

Tài liệu này xác lập ranh giới kiến trúc (Plan 002) cho dự án theo chuẩn **Clean Architecture**:

```
                       ┌────────────────────────────┐
                       │   Delivery (HTTP / gRPC)   │
                       └─────────────┬──────────────┘
                                     │ calls
                                     ▼
                       ┌────────────────────────────┐
                       │  Application (Use Cases)   │
                       └─────────────┬──────────────┘
                                     │ uses
                                     ▼
                       ┌────────────────────────────┐
                       │   Domain (Core Entities)   │
                       └────────────────────────────┘
                                     ▲
                                     │ implements interfaces
                       ┌─────────────┴──────────────┐
                       │ Infrastructure (DB / Repos)│
                       └────────────────────────────┘
```

## Các quy tắc ranh giới nghiêm ngặt:
1. **Domain Layer (`domain/`):**
   - Chỉ chứa các thực thể nghiệp vụ cốt lõi và mã lỗi thuần túy.
   - **Tuyệt đối không import** bất kỳ package nào thuộc `database/sql`, `net/http` hoặc thư viện bên ngoài.
2. **Application Layer (`service/` hoặc `usecase/`):**
   - Chứa logic nghiệp vụ kinh doanh.
   - Khai báo Interface (Ports) cho Repository tại nơi cần sử dụng (*Accept interfaces, return structs*).
3. **Infrastructure Layer (`internal/repository/`):**
   - Triển khai cụ thể các Adapter (PostgreSQL, Redis, S3).
   - Phụ thuộc ngược lại vào Interface của tầng Application (Dependency Inversion).
4. **Delivery Layer (`internal/delivery/http/`):**
   - Tiếp nhận HTTP requests, parse DTO, kiểm tra tính hợp lệ và ủy thác cho Application use-cases.
   - Chuẩn hóa lỗi trả về theo chuẩn RFC 7807 Problem Details.
