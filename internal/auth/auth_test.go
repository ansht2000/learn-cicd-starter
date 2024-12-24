package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	correctHeader := http.Header{}
	incorrectHeaderMalformed := http.Header{}
	incorrectHeaderNoAuth := http.Header{}
	correctHeader.Add("Authorization", "ApiKey secretstuff")
	incorrectHeaderMalformed.Add("Authorization", "secretstuff")

	cases := []struct{
		header http.Header
		expectedAPIKey string
		expectedError error
	}{
		{
			header: correctHeader,
			expectedAPIKey: "secretstuff",
			expectedError: nil,
		},
		{
			header: incorrectHeaderMalformed,
			expectedAPIKey: "",
			expectedError: ErrMalformedAuthHeader,
		},
		{
			header: incorrectHeaderNoAuth,
			expectedAPIKey: "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
	}

	for i, c := range cases {
		apiKey, err := GetAPIKey(c.header)
		if apiKey != c.expectedAPIKey || err != c.expectedError {
			t.Errorf("test %d: Failed to parse API key %s from header %v, instead got key %s", i + 1, c.expectedAPIKey, c.header, apiKey)
			t.Fail()
		}
	}
}