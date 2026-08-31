package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 2)
	expected := "aa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCompare(t *testing.T) {
	compared := Compare("a", "a")
	expected := true

	if compared != expected {
		t.Errorf("expected %t but got %t", expected, compared)
	}
}

// Benchmarking

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

// Example to document
func ExampleRepeat() {
	repeated := Repeat("hello", 2)
	fmt.Println(repeated)
	// Output: hellohello
}
