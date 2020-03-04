package failpointdemo

import (
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo()
	}
}
func BenchmarkFooWithFailpointMarker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FooWithFailpointMarker()
	}
}
func BenchmarkFooWithFailpointInjection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FooWithFailpointInjection()
	}
}

func TestFooWithFailpointMarker(t *testing.T) {
	assert := func(got, expected string) {
		if got != expected {
			t.Fatal("expected ", expected, "got ", got)
		}
	}

	assert(FooWithFailpointMarker(), "foo")

	ReturnString.Enable("return(\"\")")
	assert(FooWithFailpointMarker(), "")
	ReturnString.Disable()
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bar()
	}
}

func BenchmarkBarWithFailpointMarker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BarWithFailpointMarker()
	}
}

func BenchmarkBarWithFailpointInjection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BarWithFailpointInjection()
	}
}
