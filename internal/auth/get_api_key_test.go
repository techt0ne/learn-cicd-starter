package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			wantKey: "my-secret-key",
			wantErr: false,
		},
		{
			name:    "missing header",
			headers: http.Header{},
			wantKey: "",
			wantErr: true,
		},
		{
			name: "malformed header - no ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-key"},
			},
			wantKey: "",
			wantErr: true,
		},
		{
			name: "malformed header - only one part",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tc.headers)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if gotKey != tc.wantKey {
				t.Errorf("GetAPIKey() gotKey = %q, want %q", gotKey, tc.wantKey)
			}
		})
	}
}