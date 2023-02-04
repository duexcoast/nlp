package nlp

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

type tokenizeCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	data, err := ioutil.ReadFile("tokenize_cases.toml")
	require.NoError(t, err)
	var testCases struct {
		Cases []tokenizeCase
	}

	err = toml.Unmarshal(data, &testCases)
	require.NoError(t, err, "Unmarshal TOML")
	return testCases.Cases
}

func TestTokenizeTable(t *testing.T) {
	tokenCases := loadTokenizeCases(t)
	for _, tc := range tokenCases {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)
	// if tokens != Tokenize(text) // Can't compare slices with == in Go (only to nil)

	require.Equal(t, expected, tokens)
	// if !reflect.DeepEqual(expected, tokens) {
	// 	t.Fatalf("expected %#v, got %#v", expected, tokens)
	// }

}

func FuzzTokenize(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		t.Fatal(text)
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			if !strings.Contains(lText, tok) {
				t.Fatal(tok)
			}
		}
	})

}
