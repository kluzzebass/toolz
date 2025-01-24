package toolz

import (
	"fmt"
	"log/slog"
	"runtime"
)

type CallStackEntry struct {
	FuncName string `json:"funcName"`
	FileName string `json:"fileName"`
	Line     int    `json:"line"`
}

type CallStack []CallStackEntry

func (c CallStackEntry) String() string {
	return fmt.Sprintf("%s:%d %s", c.FileName, c.Line, c.FuncName)
}

// CallStack is a utility function that returns the call stack an array of CallStackEntry objects.
func GetCallStack(size int) CallStack {
	stackSize := size
	if size < 1 {
		stackSize = 10
	}
	pcs := make([]uintptr, stackSize)
	n := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	cse := make(CallStack, 0, n)
	for {
		frame, more := frames.Next()
		cse = append(cse, CallStackEntry{
			FuncName: frame.Function,
			FileName: frame.File,
			Line:     frame.Line,
		})
		if !more {
			break
		}
	}

	return cse
}

func (c CallStack) String() string {
	var s string
	for _, v := range c {
		s += fmt.Sprintf("%s\n", v)
	}
	return s
}

func (c CallStack) Slogify(name string) slog.Attr {
	arr := make([]any, 0, len(c))

	// compute the number of digits in the largest index
	maxIndex := len(c) - 1
	maxIndexDigits := 1
	for maxIndex > 9 {
		maxIndex /= 10
		maxIndexDigits++
	}

	// generate the formatting string for the key
	keyFmt := fmt.Sprintf("%%0%dd", maxIndexDigits)

	for i, v := range c {
		key := fmt.Sprintf(keyFmt, i)
		grp := slog.Group(
			key,
			slog.String("funcName", v.FuncName),
			slog.String("fileName", v.FileName),
			slog.Int("line", v.Line),
		)
		arr = append(arr, grp)
	}

	return slog.Group(name, arr...)
}
