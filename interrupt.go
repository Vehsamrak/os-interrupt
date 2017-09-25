package osinterrupt

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
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
        fmt.Printf("\n%v\n", message)
    })
}

// Handle signal and call callback function before exit
func HandleSignal(signalCode syscall.Signal, callback func()) {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, signalCode)

    go func() {
        for range signals {
            callback()
            os.Exit(0)
        }
    }()
}

// Handle SIGTERM and call callback function before exit
func HandleTerminateSignal(callback func()) {
    HandleSignal(syscall.SIGTERM, callback)
}
