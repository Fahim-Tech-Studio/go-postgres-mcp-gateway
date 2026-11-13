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
