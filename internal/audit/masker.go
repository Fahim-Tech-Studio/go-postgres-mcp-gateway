package audit

import (
    "strings"
)

// SensitiveKeys defines column name substrings that trigger redaction.
var SensitiveKeys = []string{
    "password",
    "secret",
    "token",
    "api_key",
    "ssn",
    "credit_card",
}

// MaskRow inspects map keys and masks values if the column name is sensitive.
func MaskRow(row map[string]interface{}) map[string]interface{} {
    maskedRow := make(map[string]interface{})

    for k, v := range row {
        if isSensitiveColumn(k) {
            maskedRow[k] = "[REDACTED-PII]"
        } else {
            maskedRow[k] = v
        }
    }

    return maskedRow
}

func isSensitiveColumn(colName string) bool {
    lower := strings.ToLower(colName)
    for _, sensitive := range SensitiveKeys {
        if strings.Contains(lower, sensitive) {
            return true
        }
    }
    return false
}
