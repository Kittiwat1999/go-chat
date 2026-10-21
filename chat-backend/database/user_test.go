package database

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUserJSONDoesNotExposePassword(t *testing.T) {
	const passwordHash = "$2a$12$not-a-real-hash"

	response, err := json.Marshal(User{
		Username: "alice",
		Password: passwordHash,
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	if strings.Contains(string(response), "password") || strings.Contains(string(response), passwordHash) {
		t.Fatalf("user JSON exposed password data: %s", response)
	}
}
