package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectedError  error
	}{
		{
			name:           "Valid API Key",
			headers:        http.Header{"Authorization": []string{"ApiKey abc123"}},
			expectedAPIKey: "abc123",
			expectedError:  nil,
		},
		{
			name:           "No Auth Header",
			headers:        http.Header{},
			expectedAPIKey: "",
			expectedError:  ErrNoAuthHeaderIncluded,
		},
		{
			name:           "Malformed Header",
			headers:        http.Header{"Authorization": []string{"Bearer abc123"}},
			expectedAPIKey: "",
			expectedError:  errors.New("malformed authorization header"),
		},
		{
			name:           "Empty Header",
			headers:        http.Header{"Authorization": []string{"ApiKey"}},
			expectedAPIKey: "",
			expectedError:  errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call your GetAPIKey function here
			apiKey, err := GetAPIKey(tt.headers)

			// Error handling check
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error %v, but got %v", tt.expectedError, err)
			}

			// API key check
			if apiKey != tt.expectedAPIKey {
				t.Errorf("Expected key %v, but got %v", tt.expectedAPIKey, apiKey)
			}
		}) // <-- close the t.Run function here
	}
}
