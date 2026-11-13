package postgres

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5/pgxpool"
)

// GatewayPool wraps the pgxpool connection for secure read-only execution.
type GatewayPool struct {
    Pool *pgxpool.Pool
}

// NewGatewayPool initializes a connection pool to PostgreSQL.
func NewGatewayPool(ctx context.Context, connString string) (*GatewayPool, error) {
    pool, err := pgxpool.New(ctx, connString)
    if err != nil {
        return nil, fmt.Errorf("failed to create connection pool: %w", err)
    }

    // Verify connection
    if err := pool.Ping(ctx); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return &GatewayPool{Pool: pool}, nil
}

// GetTableSchemas queries information_schema so AI agents understand available tables and columns.
func (g *GatewayPool) GetTableSchemas(ctx context.Context) (string, error) {
    query := `
        SELECT table_name, column_name, data_type 
        FROM information_schema.columns 
        WHERE table_schema = 'public' 
        ORDER BY table_name, ordinal_position;
    `

    rows, err := g.Pool.Query(ctx, query)
    if err != nil {
        return "", fmt.Errorf("failed to fetch schema: %w", err)
    }
    defer rows.Close()

    var schemaSummary string = "Public Tables & Columns:\n"
    for rows.Next() {
        var tableName, colName, dataType string
        if err := rows.Scan(&tableName, &colName, &dataType); err != nil {
            return "", err
        }
        schemaSummary += fmt.Sprintf("- Table: %s | Column: %s (%s)\n", tableName, colName, dataType)
    }

    return schemaSummary, nil
}
