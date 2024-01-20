package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatsFiveTimes(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatsFourTimes(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	repeat := Repeat("A", 3)
	fmt.Println(repeat)
	// Output: AAA
}
