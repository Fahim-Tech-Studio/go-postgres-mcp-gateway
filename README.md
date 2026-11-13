# Go PostgreSQL MCP Gateway

A high-performance, security-first **Model Context Protocol (MCP)** server and proxy built in **Golang**, enabling AI agents (Claude, Cursor, and custom LLM workflows) to safely inspect schemas and run read-only queries against **PostgreSQL** databases.

## 🛡️ Security & Guardrail Architecture

Connecting AI agents directly to production databases introduces risks of destructive mutations (`DROP`, `DELETE`, `UPDATE`) and accidental PII exposure. This gateway enforces strict boundaries between the LLM and PostgreSQL:

```text
go-postgres-mcp-gateway/
├── cmd/
│   └── server/main.go        # CLI entrypoint (Stdio & SSE/HTTP MCP transports)
├── internal/
│   ├── mcp/                  # MCP protocol handlers (Tools & Resources)
│   ├── guard/                # SQL AST validation (strict SELECT-only allowlist)
│   ├── postgres/             # Read-only pgx connection pool & schema inspector
│   └── audit/                # Structured query logging & PII column masking
├── web/                      # React + Tailwind admin control plane
├── Dockerfile                # Multi-stage minimal production image
└── go.mod                    # Go module dependencies
[![Go CI](https://github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/actions/workflows/ci.yml/badge.svg)](https://github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
```
## Key Features
Driver-Level Read-Only Enforcement: Executes queries inside SET TRANSACTION READ ONLY sessions using pgx connection pooling.

SQL Guardrails: Parses incoming SQL statements to reject write operations, DDL commands, and multi-statement injections before they ever touch the database.

Dynamic Schema Resources: Exposes postgres://schemas via MCP so AI agents understand table structures and foreign keys without hallucinating column names.

PII Column Masking: Automatically redacts sensitive fields (e.g., password_hash, api_key, ssn) based on configurable rules.

Enterprise Control Plane (Roadmap): React + PostgreSQL management UI for team RBAC, token metrics, and real-time query audit trails.

## Tech Stack
Core Gateway: Go (Golang), pgx PostgreSQL Driver, MCP Go SDK

Control Plane UI: React, TypeScript, Tailwind CSS

Database: PostgreSQL

## License
Released under the MIT License. Developed by Fahim Tech Studio.
