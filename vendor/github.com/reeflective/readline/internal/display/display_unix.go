//go:build unix
// +build unix

package display

import (
	"os"
	"os/signal"
	"syscall"
)

// WatchResize redisplays the interface on terminal resize events.
func WatchResize(eng *Engine) chan<- bool {
	done := make(chan bool, 1)

	resizeChannel := make(chan os.Signal, 1)
	signal.Notify(resizeChannel, syscall.SIGWINCH)

	go func() {
		for {
			select {
			case <-resizeChannel:
				eng.completer.GenerateCached()
				eng.Refresh()
			case <-done:
				return
			}
		}
	}()

	return done
}
