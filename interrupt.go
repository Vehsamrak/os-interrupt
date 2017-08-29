package osinterrupt

import (
	"fmt"
	"os"
	"os/signal"
)

// Handle SIGINT (Ctrl+C) and call callback function before exit
func HandleInterruptSignal(callback func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		for osSignal := range signals {
			// if signal is ^C (SIGINT)
			if osSignal.String() == "interrupt" {
				callback()
				os.Exit(0)
			}
		}
	}()
}

// Print message on SIGINT (Ctrl+C) before exit
func PrintMessageOnInterruptSignal(message string) {
	HandleInterruptSignal(func() {
		fmt.Println("\n", message)
	})
}
