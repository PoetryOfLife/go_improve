package server

import "os"

func WaitForShutdown() {
	signals1 := make(chan os.Signal, 1)
	signals2 := make(chan os.Signal, 1)
	select {
	case <-signals1:

	case <-signals2:

	}
}
