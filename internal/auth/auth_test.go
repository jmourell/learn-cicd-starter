package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type result struct {
	Auth string
	Err  error
}

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  result
	}{
		"simple": {input: http.Header{}, want: result{Auth: "", Err: ErrNoAuthHeaderIncluded}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			auth, err := GetAPIKey(tc.input)

			diff := cmp.Diff(tc.want, result{Auth: auth, Err: err}, cmpopts.EquateErrors())
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
