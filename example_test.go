package nlp_test

import (
	"fmt"

	"github.com/duexcoast/nlp"
)

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := nlp.Tokenize(text)

	fmt.Println(tokens)

	// Output:
	// [who on first]
}
