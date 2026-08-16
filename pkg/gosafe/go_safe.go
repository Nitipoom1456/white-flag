package gosafe

import (
	"fmt"
	"runtime/debug"
)

func RecoverPanic(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				var err error
				if e, ok := r.(error); ok {
					err = e
				} else {
					err = fmt.Errorf("%v", r)
				}
				stack := debug.Stack()
				fmt.Printf("panic: %v\n%s\n", err, stack)
			}
		}()
		fn()
	}()
}
