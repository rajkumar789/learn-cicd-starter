package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// arrange test cases
	// each case should include:
	// - a name
	// - headers
	// - expected key
	// - whether you expect an error

	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		wantErr     bool
	}{
		{
			name: "valid authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey dummy-api-key-12345"},
			},
			expectedKey: "dummy-api-key-12345",
			wantErr:     false,
		},
		{
			name: "missing or malformed authorization header",
			headers: http.Header{
				// Leaving this empty simulates a missing Authorization header
			},
			expectedKey: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if tt.wantErr && err == nil {
				t.Fatal("expected an error, but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("did not expect error, but got: %v", err)
			}

			if key != tt.expectedKey {
				t.Fatalf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}

