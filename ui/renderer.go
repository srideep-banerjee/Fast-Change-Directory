package ui

import (
	"log"
	"strings"

	"github.com/nsf/termbox-go"
)

func RenderInputField(input string) {
	RenderInputBox()
    RenderInputText(input)
    err := termbox.Flush()
    if err != nil {
        log.Panic(err)
    }
}

func RenderInputBox() {
    r, c := termbox.Size()
    line := "+"
    for i := 0; i < r - 2; i++ {
        line += "-"
    }
    line += "+"
    print(0, c - 3, line)
    print(0, c - 2, "|")
    print(r - 1, c - 2, "|")
    print(0, c - 1, line)
}

func RenderInputText(input string) {
    r, c := termbox.Size()
    l := len(input)
    if l  > r - 2 {
        input = input[l - (r - 2):]
        l = len(input)
    }
    print(l + 1, c - 2, strings.Repeat(" ", r - 2 - l))
    print(1, c - 2, input)
    termbox.SetCursor(len(input) + 1, c - 2)
}

func print(x int, y int, text string) {
    for i, ch := range text {
        termbox.SetCell(x + i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
    }
}

