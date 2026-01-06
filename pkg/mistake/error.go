package mistake

import (
	"runtime/debug"
	"super-llm/pkg/gstr"

	"github.com/pkg/errors"
)

func Unwrap(err error) {
	if err != nil {
		panic(errors.Wrap(err, gstr.BytesToString(debug.Stack())))
	}
}

func UnwrapNotTrace(err error) {
	if err != nil {
		panic(err)
	}
}
