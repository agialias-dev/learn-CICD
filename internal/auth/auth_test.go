package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name  string
		input *http.Request
		want  string
	}{
		{
			name: "Valid header",
			input: &http.Request{
				Header: http.Header{
					"Authorization": []string{"ApiKey test_api_key"},
				},
			},
			want: "test_api_key",
		},
		{
			name: "Invalid header",
			input: &http.Request{
				Header: http.Header{
					"Authorization": []string{"Bearer %~@>%£${}"},
				},
			},
			want: "",
		},
	}

	for _, testcase := range tests {
		result, err := GetAPIKey(testcase.input.Header)
		if result != testcase.want {
			t.Errorf("Test %q failed: expected %q, got %q (Error: %v)", testcase.name, testcase.want, result, err)
		}
	}
}
