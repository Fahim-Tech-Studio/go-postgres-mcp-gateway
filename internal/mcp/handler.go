package mcp

import (
    "context"
    "fmt"

    "github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/internal/audit"
    "github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/internal/guard"
    "github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/internal/postgres"
)

// MCPToolHandler coordinates incoming AI agent requests with database safety checks.
type MCPToolHandler struct {
    DB *postgres.GatewayPool
}

// NewMCPToolHandler creates a new handler instance.
func NewMCPToolHandler(db *postgres.GatewayPool) *MCPToolHandler {
    return &MCPToolHandler{DB: db}
}

// ExecuteReadQuery validates SQL safety, runs the query, and masks PII columns.
func (h *MCPToolHandler) ExecuteReadQuery(ctx context.Context, sqlQuery string) (string, error) {
    // Step 1: Enforce SQL Guardrails (Block mutations)
    if err := guard.ValidateReadOnly(sqlQuery); err != nil {
        return "", fmt.Errorf("guardrail blocked query: %w", err)
    }

    // Step 2: Execute query against PostgreSQL connection pool
    rows, err := h.DB.Pool.Query(ctx, sqlQuery)
    if err != nil {
        return "", fmt.Errorf("database execution error: %w", err)
    }
    defer rows.Close()

    // Extract column names
    fieldDescs := rows.FieldDescriptions()
    colNames := make([]string, len(fieldDescs))
    for i, fd := range fieldDescs {
        colNames[i] = string(fd.Name)
    }

    var results []map[string]interface{}

    // Step 3: Iterate and fetch results
    for rows.Next() {
        values, err := rows.Values()
        if err != nil {
            return "", err
        }

        rowMap := make(map[string]interface{})
        for i, val := range values {
            rowMap[colNames[i]] = val
        }

        // Step 4: Apply PII Masking
        maskedRow := audit.MaskRow(rowMap)
        results = append(results, maskedRow)
    }

    return fmt.Sprintf("Query executed successfully. Rows returned: %d. Data: %v", len(results), results), nil
}
