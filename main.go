package main

import (
	"fmt"
	"log"
	"os"
	_ "strings"

	"me/fast-cd/ui"
	_ "me/fast-cd/validation"

	"github.com/nsf/termbox-go"
)

var input = ""

func main() {
    f, err := os.Create("logs.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    log.SetOutput(f)
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    fmt.Println("\033[?1049h")

    err = termbox.Init()
    if err != nil {
        log.Panic(err)
    }
    defer close()
    eventLoop()
}

func close() {
    termbox.Close()
    fmt.Print("\033[?1049l")
}

func eventLoop() {
    var event = termbox.PollEvent()
    if event.Err != nil {
        panic(event.Err)
    }

    EventLoop:
    for {

        switch (event.Type) {
        case termbox.EventKey:
            if event.Key == termbox.KeyCtrlC {
                break EventLoop
            } else if event.Key == termbox.KeyBackspace {
                input = input[:len(input) - 1]
                ui.RenderInputField(input)
            } else {
                input += string(event.Ch)
                ui.RenderInputField(input)
            }
        }

        event = termbox.PollEvent()
        if event.Err != nil {
            log.Panic(event.Err)
        }
    }
}
