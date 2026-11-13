package guard

import (
    "errors"
    "strings"
)

// ValidateReadOnly ensures the query is strictly a SELECT statement and blocks mutations.
func ValidateReadOnly(query string) error {
    cleaned := strings.TrimSpace(strings.ToUpper(query))

    // Must start with SELECT or WITH
    if !strings.HasPrefix(cleaned, "SELECT") && !strings.HasPrefix(cleaned, "WITH") {
        return errors.New("security violation: only SELECT and WITH (read-only) queries are allowed")
    }

    // Block forbidden mutation keywords
    forbidden := []string{"INSERT", "UPDATE", "DELETE", "DROP", "ALTER", "TRUNCATE", "CREATE", "GRANT", "REVOKE"}
    for _, word := range forbidden {
        if containsKeyword(cleaned, word) {
            return errors.New("security violation: mutation keyword '" + word + "' is blocked by MCP gateway guardrails")
        }
    }

    return nil
}

func containsKeyword(query, keyword string) bool {
    words := strings.Fields(query)
    for _, w := range words {
        cleanWord := strings.Trim(w, "();,")
        if cleanWord == keyword {
            return true
        }
    }
    return false
}
