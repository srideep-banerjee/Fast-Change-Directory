package main

import (
	"fmt"
	_ "strings"

	"me/fast-cd/ui"
	_ "me/fast-cd/validation"

	"github.com/nsf/termbox-go"
)

var input = ""

func main() {
    fmt.Println("\033[?1049h")

    err := termbox.Init()
    if err != nil {
        fmt.Println("Terminal does not support required library: termbox-go")
        panic(err)
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
    for true {

        switch (event.Type) {
        case termbox.EventKey:
            if event.Key == termbox.KeyCtrlC {
                break EventLoop
            } else if event.Key == termbox.KeyBackspace {
                input = input[:len(input) - 1]
                ui.RenderInputBox(input)
            } else {
                input += string(event.Ch)
                ui.RenderInputBox(input)
            }
        }

        event = termbox.PollEvent()
        if event.Err != nil {
            panic(event.Err)
        }
    }
}
