package concat_test

import (
	"strings"
	"testing"
)
var args = []string{"1", "2", "3", "4", "5", "test", "this", "is", "awesome"}

func concat(args []string) {
	r, sep := "", ""
	for _, a := range args {
		r += sep + a
		sep = " "
	}
}

func BenchmarkConcat(b *testing.B) { // to use the testing, must use Camel naming rule. First should be capital char.
	for i := 0; i < b.N; i++ {
		concat(args)
	}
} //537ns

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
} //199ns
