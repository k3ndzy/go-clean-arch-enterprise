# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0-phase1] - 2026-11-21

### Added
- **Plan 001:** Initialized Go module and executable entry point at `cmd/api/main.go`.
- **Plan 002:** Defined architecture boundaries document at `docs/architecture/boundaries.md`.
- **Plan 003:** Created configuration loader with environment variable validation and unit tests (`internal/config`).
- **Plan 004:** Built application bootstrap in `internal/app/app.go` with explicit dependency wiring.
- **Plan 005:** Implemented HTTP server with read/write/idle timeouts and graceful shutdown on OS signals (`internal/delivery/http/server.go`).
- **Plan 006:** Implemented `/livez` and `/readyz` health check endpoints with unit tests (`internal/delivery/http/handler/health.go`).
- **Plan 007:** Built structured JSON logger (`pkg/logger`) with automated sensitive data redaction (passwords, tokens, keys) and request ID correlation middleware.
- **Plan 008:** Standardized error responses adhering to RFC 7807 Problem Details (`pkg/problem`).
- **Plan 009:** Added developer `Makefile` with targets for build, run, test, lint, clean, and coverage.
- **Plan 010:** Authored comprehensive initial `README.md` and 100-plans progress tracking board.
