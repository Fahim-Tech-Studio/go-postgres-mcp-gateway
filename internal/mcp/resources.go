package mcp

import (
    "context"
    "fmt"

    "github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/internal/postgres"
)

// MCPResourceProvider manages dynamic database schema and security metadata resources for AI agents.
type MCPResourceProvider struct {
    DB *postgres.GatewayPool
}

// NewMCPResourceProvider initializes the resource provider.
func NewMCPResourceProvider(db *postgres.GatewayPool) *MCPResourceProvider {
    return &MCPResourceProvider{DB: db}
}

// ListResources returns URIs representing available database inspection points.
func (p *MCPResourceProvider) ListResources() []map[string]string {
    return []map[string]string{
        {
            "uri":         "postgres://schema/public",
            "name":        "Public Database Schema",
            "description": "Live columnar mapping of all tables in the public PostgreSQL schema.",
            "mimeType":    "application/json",
        },
        {
            "uri":         "postgres://security/guardrails",
            "name":        "Active AST Guardrail Rules",
            "description": "Read-only enforcement policies and PII redaction definitions.",
            "mimeType":    "text/markdown",
        },
    }
}

// ReadResource fetches live payload for a given resource URI.
func (p *MCPResourceProvider) ReadResource(ctx context.Context, uri string) (string, error) {
    switch uri {
    case "postgres://schema/public":
        schema, err := p.DB.GetTableSchemas(ctx)
        if err != nil {
            return "", err
        }
        return schema, nil

    case "postgres://security/guardrails":
        return `# Active Security Policies\n- AST Whitelist: SELECT, WITH\n- AST Blocklist: DROP, INSERT, UPDATE, DELETE, ALTER\n- PII Redaction: Enabled for columns matching password, secret, token, api_key`, nil

    default:
        return "", fmt.Errorf("resource not found: %s", uri)
    }
}
