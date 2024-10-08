package iteration_test

import (
	"fmt"
	"testing"

	"github.com/Jaytpa01/learn-go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := iteration.Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
