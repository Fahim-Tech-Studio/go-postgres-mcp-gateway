package main

import (
    "fmt"
    "log"
    "os"

    "github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/internal/guard"
)

func main() {
    fmt.Println("Initializing Go PostgreSQL MCP Gateway...")

    // Test 1: Valid Read-Only Query
    testQuery := "SELECT id, username FROM users WHERE active = true;"
    if err := guard.ValidateReadOnly(testQuery); err != nil {
        log.Fatalf("Query rejected: %v", err)
    }
    fmt.Println("[SECURITY PASSED] Query is safe and read-only:", testQuery)

    // Test 2: Blocked Malicious/Mutation Query
    badQuery := "DROP TABLE users;"
    if err := guard.ValidateReadOnly(badQuery); err != nil {
        fmt.Println("[SECURITY BLOCKED SUCCESSFULLY]:", err)
    }

    fmt.Println("MCP Gateway security guardrails initialized successfully.")
    os.Exit(0)
}
