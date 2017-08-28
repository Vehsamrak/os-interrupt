package os_interrupt

import (
    "fmt"
    "os"
    "os/signal"
)

func HandleInterruptSignals(message string) {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    go func(){
        for osSignal := range signals {
            // if signal is ^C (SIGINT)
            if osSignal.String() == "interrupt" {
                fmt.Println(message)
                os.Exit(0)
            }
        }
    }()
}
