package guard

import (
    "testing"
)

func TestValidateReadOnly(t *testing.T) {
    tests := []struct {
        name    string
        query   string
        wantErr bool
    }{
        {"Valid Select Query", "SELECT id, name FROM users WHERE active = true;", false},
        {"Valid With Clause", "WITH recent AS (SELECT * FROM logs) SELECT * FROM recent;", false},
        {"Blocked Drop Table", "DROP TABLE users;", true},
        {"Blocked Insert Record", "INSERT INTO users (name) VALUES ('Hacker');", true},
        {"Blocked Update Record", "UPDATE users SET role = 'admin';", true},
        {"Blocked Delete Record", "DELETE FROM users WHERE id = 1;", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateReadOnly(tt.query)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateReadOnly() query = %q, error = %v, wantErr %v", tt.query, err, tt.wantErr)
            }
        })
    }
}
