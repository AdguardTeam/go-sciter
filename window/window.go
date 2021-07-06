package window

import (
	"github.com/AdguardTeam/go-sciter"
	"runtime"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}

// https://github.com/golang/go/wiki/LockOSThread
// https://github.com/AdguardTeam/go-sciter/issues/201
func init() {
	runtime.LockOSThread()
}
