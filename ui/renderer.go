package ui

import (
	"log"

	"github.com/nsf/termbox-go"
)

func RenderInputBox(input string) {
	r, c := termbox.Size()
    line := "+"
    for i := 0; i < r - 2; i++ {
        line += "-"
    }
    line += "+"
    print(0, c - 3, line)
    for i := len(input); i < r; i++ {
        print(i, c - 2, " ")
    }
    print(0, c - 2, "|" + input)
    print(r - 1, c - 2, "|")
    print(0, c - 1, line)
    err := termbox.Flush()
    if err != nil {
        log.Panic(err)
    }
}

func print(x int, y int, text string) {
    for i, ch := range text {
        termbox.SetCell(x + i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
    }
}

