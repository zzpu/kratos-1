package x

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanonicalize(t *testing.T) {
	t.Run("Canonicalize emails", func(t *testing.T) {
		errorTestAddresses := []string{"", "this is not an email address", "example.com"}

		for _, adr := range errorTestAddresses {
			res, err := CanonicalizeEmail(adr)

			assert.Equal(t, "", res)
			assert.Error(t, err)
		}

		type successCase struct {
			address  string
			expected string
		}
		successTestStrings := []successCase{
			{"hi@example.com", "hi@example.com"},
			{"😵@example.com", "😵@example.com"},
			{"tütülütü@tütülütü.com", "tütülütü@xn--ttlt-0rabbb.com"},
		}

		for _, testCase := range successTestStrings {
			res, err := CanonicalizeEmail(testCase.address)

			assert.Equal(t, testCase.expected, res)
			assert.NoError(t, err)
		}
	})

	t.Run("Canonicalize URIs", func(t *testing.T) {
		errorTestURIs := []string{"not a valid uri", ""}

		for _, uri := range errorTestURIs {
			res, err := CanonicalizeURI(uri)

			assert.Equal(t, "", res)
			assert.Error(t, err)
		}

		type successCase struct {
			uri      string
			expected string
		}
		successTestStrings := []successCase{
			{"https://tütülütü.com", "https://xn--ttlt-0rabbb.com"},
			{"http://tütülütü.com/tütülütü", "http://xn--ttlt-0rabbb.com/t%C3%BCt%C3%BCl%C3%BCt%C3%BC"},
			{"ssh://tütülütü.com/a/longer/path/even/with/thai/ยจฆฟคฏข/writing", "ssh://xn--ttlt-0rabbb.com/a/longer/path/even/with/thai/%E0%B8%A2%E0%B8%88%E0%B8%86%E0%B8%9F%E0%B8%84%E0%B8%8F%E0%B8%82/writing"},
		}

		for _, testCase := range successTestStrings {
			res, err := CanonicalizeURI(testCase.uri)

			assert.Equal(t, testCase.expected, res)
			assert.NoError(t, err)
		}
	})
}
