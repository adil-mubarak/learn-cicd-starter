package auth_test

import (
	"net/http"
	"testing"

	"github.com/adil-mubarak/learn-cicd-starter/internal/auth"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	key, err := auth.GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if key != "my-secret-key" {
		t.Fatalf("expected 'my-secret-key', got %s", key)
	}
}

func TestGetAPIMissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != auth.ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderInvluded, got %v", err)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer wrongformat")

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
