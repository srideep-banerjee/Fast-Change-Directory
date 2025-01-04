package main

import (
	"fmt"
	_ "strings"

	"github.com/nsf/termbox-go"
    _ "me/fast-cd/validation"
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

        fmt.Println(event.Type)
        switch (event.Type) {
        case termbox.EventKey:
            if event.Key == termbox.KeyCtrlC {
                break EventLoop
            }
        }

        event = termbox.PollEvent()
        if event.Err != nil {
            panic(event.Err)
        }
    }
}

func renderInputBox() {
}
