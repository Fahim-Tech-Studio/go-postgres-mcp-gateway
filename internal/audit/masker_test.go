package audit

import (
    "testing"
)

func TestMaskRow(t *testing.T) {
    input := map[string]interface{}{
        "id":            1,
        "username":      "fahim_architect",
        "password_hash": "$2a$12$secretHashValue",
        "api_key":       "sk_live_9988776655",
    }

    masked := MaskRow(input)

    if masked["username"] != "fahim_architect" {
        t.Errorf("Expected username to remain unmasked, got %v", masked["username"])
    }

    if masked["password_hash"] != "[REDACTED-PII]" {
        t.Errorf("Expected password_hash to be redacted, got %v", masked["password_hash"])
    }

    if masked["api_key"] != "[REDACTED-PII]" {
        t.Errorf("Expected api_key to be redacted, got %v", masked["api_key"])
    }
}
