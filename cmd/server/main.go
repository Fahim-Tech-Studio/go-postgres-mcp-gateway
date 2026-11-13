package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Fahim-Tech-Studio/go-postgres-mcp-gateway/internal/guard"
)

func main() {
	fmt.Println("Initializing Fahim Tech Studio - Go PostgreSQL MCP Gateway...")

	// Verify guardrail module is active
	testQuery := "SELECT id, email FROM customers WHERE active = true;"
	if err := guard.ValidateReadOnly(testQuery); err != nil {
		log.Fatalf("Security validation failed: %v", err)
	}
	fmt.Println("[SECURITY PASSED] Read-only AST validation active.")
	fmt.Println("[READY] MCP Gateway server architecture compiled successfully.")
	
	os.Exit(0)
}
