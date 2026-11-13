# Go PostgreSQL MCP Gateway

[![Go CI](https://github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/actions/workflows/ci.yml/badge.svg)](https://github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Control Plane](https://img.shields.io/badge/Control%20Plane-React%20%2B%20Tailwind-0f172a?style=flat&logo=react)](./web)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> **Stop AI agents from executing destructive mutations (`DROP`, `DELETE`, `UPDATE`), hallucinating column names, or leaking sensitive PII from production PostgreSQL databases.**

A high-performance, security-first **Model Context Protocol (MCP)** server, database proxy, and AST query firewall built in **Golang**, paired with a **React + TypeScript + Tailwind CSS** admin control plane. Designed for enterprise AI engineering teams using Claude, Cursor, and custom LLM workflows to safely inspect PostgreSQL schemas, execute read-only queries, and enforce automated PII masking.


## Why This Exists

| Production Risk | Gateway Enforcement |
| :--- | :--- |
| **Destructive Mutations:** Agents execute `DROP`, `DELETE`, `UPDATE`, or multi-statement SQL injections. | **Dual-Layer Guardrails:** AST SQL parser allowlists `SELECT`-only queries AND `pgx` executes inside `SET TRANSACTION READ ONLY` sessions. |
| **Schema Hallucinations:** Agents guess non-existent table or foreign key names. | **Dynamic Schema Resources:** Exposes `postgres://schemas` via MCP so agents inspect exact table structures and foreign keys before querying. |
| **PII Context Leaks:** Sensitive fields (`password_hash`, `api_key`, `ssn`) enter LLM context windows. | **Real-Time PII Column Masking:** Automatically redacts configured sensitive fields prior to returning MCP tool responses. |
| **Blind AI Operations:** Zero visibility into agent query patterns or blocked attempts. | **Structured Audit Logging:** Captures every query, transport call (Stdio or SSE/HTTP), and redaction event. |


## Security & Guardrail Architecture

```mermaid
flowchart LR
    A[Claude / Cursor / LLM Agent] -->|Stdio or SSE/HTTP MCP| B[internal/mcp Handlers]
    B --> C[internal/guard SQL AST Firewall]
    C -->|Write / DDL / Multi-Statement| X[Blocked & Logged]
    C -->|Verified SELECT| D[(internal/postgres pgx Read-Only Pool)]
    D -->|Raw Rows| E[internal/audit PII Masker]
    E -->|Redacted JSON| A
    F[web/ React + Tailwind Control Plane] -.->|RBAC, Metrics & Audit Trails| E
```

### Repository Structure

```text

go-postgres-mcp-gateway/
├── cmd/
│   └── server/main.go       # CLI entrypoint (Stdio & SSE/HTTP MCP transports)
├── internal/
│   ├── mcp/                 # MCP protocol handlers (Tools & Resources)
│   ├── guard/               # SQL AST validation (strict SELECT-only allowlist)
│   ├── postgres/            # Read-only pgx connection pool & schema inspector
│   └── audit/               # Structured query logging & PII column masking
├── web/                     # React + Tailwind admin control plane
├── Dockerfile               # Multi-stage minimal production image
└── go.mod                   # Go module dependencies

```


## 60-Second Quickstart

### Option A: Run with Docker

```bash
git clone https://github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway.git
cd go-postgres-mcp-gateway
docker build -t go-postgres-mcp-gateway .
docker run -p 8080:8080 \
  -e DATABASE_URL="postgres://readonly_user:password@host.docker.internal:5432/mydb?sslmode=disable" \
  go-postgres-mcp-gateway
```

### Option B: Run Locally (Go Gateway + Control Plane UI)

```bash
# 1. Start the Go MCP Gateway
export DATABASE_URL="postgres://readonly_user:password@localhost:5432/mydb?sslmode=disable"
go run ./cmd/server/main.go

# 2. Start the React Control Plane UI (in a separate terminal)
cd web && npm install && npm run dev
```

## Key Features & Tech Stack

* **Driver-Level Read-Only Enforcement:** Executes queries inside `SET TRANSACTION READ ONLY` sessions using `pgx` connection pooling.
* **SQL Guardrails:** Parses incoming SQL statements to reject write operations, DDL commands, and multi-statement injections before they ever touch the database.
* **Dynamic Schema Resources:** Exposes `postgres://schemas` via MCP so AI agents understand table structures and foreign keys without hallucinating column names.
* **PII Column Masking:** Automatically redacts sensitive fields (`password_hash`, `api_key`, `ssn`) based on configurable rules.
* **Tech Stack:** Go (Golang), `pgx` PostgreSQL Driver, MCP Go SDK, React, TypeScript, Tailwind CSS, PostgreSQL.


## Roadmap

- [x] Driver-level `SET TRANSACTION READ ONLY` via `pgx` connection pooling
- [x] SQL AST guardrail & multi-statement injection blocker
- [x] Dynamic `postgres://schemas` MCP resource exposure
- [x] Configurable PII column redaction & structured audit logging
- [ ] **Enterprise Control Plane UI:** React + PostgreSQL management UI for team RBAC, token metrics, and real-time query audit trails


## License

Released under the **MIT License**. Developed and maintained by **[Fahim Tech Studio](https://github.com/Fahim-Tech-Studio)**.
