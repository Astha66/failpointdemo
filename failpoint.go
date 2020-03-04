package failpointdemo

import (
	"fmt"

	"github.com/pingcap/failpoint"
)

// FailpointMarker marks a failpoint
type FailpointMarker bool

// Enable a failpoint marker
func (m *FailpointMarker) Enable(expr string) error {
	fpname := fmt.Sprintf("%p", m)
	if err := failpoint.Enable(fpname, expr); err != nil {
		return err
	}
	*m = true
	return nil
}

// Disable a failpoint
func (m *FailpointMarker) Disable() error {
	fpname := fmt.Sprintf("%p", m)
	if err := failpoint.Disable(fpname); err != nil {
		return err
	}
	*m = false
	return nil
}

// Eval a failpoint marker
func (m *FailpointMarker) Eval() (failpoint.Value, bool) {
	if !*m {
		return nil, false
	}
	fpname := fmt.Sprintf("%p", m)
	return failpoint.Eval(fpname)
}

var (
	// ReturnString returns an arbitrary string
	ReturnString FailpointMarker

	// Panic panics current routine
	Panic FailpointMarker

	// Sleep blocks the functon for a while
	Sleep FailpointMarker
)

// Foo returns "foo"
//go:noinline
func Foo() string {
	return "foo"
}

// FooWithFailpointMarker returns the string "foo", introduced failpoint markder
//go:noinline
func FooWithFailpointMarker() string {
	if ReturnString {
		if val, ok := ReturnString.Eval(); ok {
			return val.(string)
		}
	}

	return "foo"
}

// FooWithFailpointInjection returns the string "foo", introduced failpoint injection
//go:noinline
func FooWithFailpointInjection() string {
	failpoint.Inject("test.FooReturnString", func(v failpoint.Value) string {
		return v.(string)
	})

	return "foo"
}

// Bar returns string "bar"
//go:noinline
func Bar() string {
	return "bar"
}

// BarWithFailpointMarker returns the string "bar", introduced failpoint markder
//go:noinline
func BarWithFailpointMarker() string {
	if ReturnString {
		if val, ok := ReturnString.Eval(); ok {
			return val.(string)
		}
	}

	if Panic {
		if val, ok := Panic.Eval(); ok {
			panic(val)
		}
	}

	return "bar"
}

// BarWithFailpointInjection returns the string "bar", introduced failpoint injection
//go:noinline
func BarWithFailpointInjection() string {
	failpoint.Inject("test.BarReturnString", func(v failpoint.Value) string {
		return v.(string)
	})
	failpoint.Inject("test.BarPanic", func(v failpoint.Value) string {
		return v.(string)
	})

	return "bar"
}
