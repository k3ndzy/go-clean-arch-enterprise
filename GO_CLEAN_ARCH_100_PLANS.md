# Go Clean Architecture Enterprise — 100 Implementation Plans Tracking Board

Reference inspiration: https://github.com/bxcodec/go-clean-arch

## Phase 1 — Foundation (001–010) [STATUS: DONE ✅ 10/10]

- [x] **Plan 001: Initialize Go module** — Module `github.com/k3ndzy/go-clean-arch-enterprise`, folder structure and `cmd/api/main.go`.
- [x] **Plan 002: Define architecture boundaries** — Documented in `docs/architecture/boundaries.md`.
- [x] **Plan 003: Configuration system** — Environment loading, validation, and defaults in `internal/config`.
- [x] **Plan 004: Application bootstrap** — Explicit dependency injection in `internal/app/app.go`.
- [x] **Plan 005: HTTP server** — Configurable timeouts and graceful shutdown in `internal/delivery/http/server.go`.
- [x] **Plan 006: Health checks** — Liveness `/livez` and Readiness `/readyz` in `internal/delivery/http/handler/health.go`.
- [x] **Plan 007: Structured logging** — Slog JSON logger with automatic sensitive field redaction in `pkg/logger`.
- [x] **Plan 008: Standard errors** — RFC 7807 Problem Details in `pkg/problem`.
- [x] **Plan 009: Developer Makefile** — Automated `build`, `run`, `test`, `lint`, and `clean` commands.
- [x] **Plan 010: Initial README & Changelog** — Comprehensive setup guide, status tracking and changelog.

---

## Phase 2 — Domain and Application (011–020) [STATUS: READY]

- [ ] **Plan 011: Core entities** — Define User, Role, Session and AuditEvent.
- [ ] **Plan 012: Value objects** — Validate email and identifiers.
- [ ] **Plan 013: Domain errors** — Create transport-independent domain errors.
- [ ] **Plan 014: Repository interfaces** — Define ports at the use-case boundary.
- [ ] **Plan 015: Application use cases** — Implement user create/read/update/delete workflows.
- [ ] **Plan 016: Dependency inversion** — Inject interfaces and prevent domain infrastructure imports.
- [ ] **Plan 017: DTO separation** — Keep HTTP schemas separate from domain models.
- [ ] **Plan 018: Business validation** — Enforce domain invariants.
- [ ] **Plan 019: Transactions abstraction** — Define atomic use-case transaction handling.
- [ ] **Plan 020: Architecture decisions** — Write ADRs explaining major tradeoffs.

---

## Phase 3 — PostgreSQL (021–030) [STATUS: READY]

- [ ] **Plan 021: Database service** — Configure local PostgreSQL in Docker Compose.
- [ ] **Plan 022: Connection pooling** — Set safe pool limits and timeouts.
- [ ] **Plan 023: Schema migrations** — Create versioned up/down migrations.
- [ ] **Plan 024: User repository** — Implement parameterized PostgreSQL CRUD.
- [ ] **Plan 025: Role repository** — Persist role assignments and constraints.
- [ ] **Plan 026: Pagination** — Implement bounded deterministic pagination.
- [ ] **Plan 027: Filtering and sorting** — Allowlist filter and sort fields.
- [ ] **Plan 028: Transactions** — Implement commit and rollback correctly.
- [ ] **Plan 029: Database error mapping** — Translate conflicts and missing records.
- [ ] **Plan 030: Repository integration tests** — Use isolated PostgreSQL test instances.

---

## Phase 4 — Authentication and Security (031–040) [STATUS: READY]

- [ ] **Plan 031: Password hashing** — Use a vetted password hash implementation.
- [ ] **Plan 032: Registration** — Validate and create user accounts.
- [ ] **Plan 033: Login** — Authenticate without exposing account existence.
- [ ] **Plan 034: Access tokens** — Issue and validate short-lived signed tokens.
- [ ] **Plan 035: Refresh rotation** — Persist revocable refresh tokens and detect reuse.
- [ ] **Plan 036: Auth middleware** — Authenticate protected HTTP endpoints.
- [ ] **Plan 037: RBAC** — Enforce role permissions on operations.
- [ ] **Plan 038: Login rate limiting** — Throttle repeated failed login attempts.
- [ ] **Plan 039: Logout** — Revoke refresh sessions.
- [ ] **Plan 040: Security review** — Test authorization, injection and data exposure.

---

## Phase 5 — REST API (041–050) [STATUS: READY]
- [ ] Plans 041-050: User CRUD, Profile, Role Admin, Validation, Envelopes, API v1, Search, Bulk, Audit, E2E tests.

## Phase 6 — Testing and Quality (051–060) [STATUS: READY]
- [ ] Plans 051-060: Domain tests, Use-case tests, Handler tests, Table-driven tests, Fixtures, Race detector, Static analysis, Coverage, Fuzz tests, Quality gates.

## Phase 7 — Docker and CI/CD (061–070) [STATUS: READY]
- [ ] Plans 061-070: Multi-stage Dockerfile, Compose stack, Build caching, GitHub Actions, Dependency scanning, Container scanning, Build artifacts, Release versioning, Staging workflow, Rollback guide.

## Phase 8 — Observability and Performance (071–080) [STATUS: READY]
- [ ] Plans 071-080: HTTP metrics, Tracing, Prometheus, Grafana, Redis caching, Query profiling, Timeout propagation, Bounded concurrency, Benchmarks, Load tests.

## Phase 9 — Enterprise Features (081–090) [STATUS: READY]
- [ ] Plans 081-090: Idempotency, Background jobs, Transactional outbox, Email notifications, API rate limiting, Feature flags, Multi-tenancy, Soft deletion, Data export, Backup recovery.

## Phase 10 — Documentation and Release (091–100) [STATUS: READY]
- [ ] Plans 091-100: OpenAPI, API examples, Contributing guide, Architecture guide, Security policy, License & attribution, Roadmap, Release checklist, Final verification, Portfolio release.
