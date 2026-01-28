package main

import "github.com/gdamore/tcell/v2"

type Input int

const (
	MoveUp Input = iota
	MoveDown
	MoveLeft
	MoveRight
    PickUp
    DropLight
    DropTrap
)

func startInputLoop(
	screen tcell.Screen,
	inputCh chan<- Input,
	quitCh chan<- struct{},
) {

    go func() {
        for {
            ev := screen.PollEvent()
            switch e := ev.(type) {

            case *tcell.EventKey:
                switch e.Key() {

                case tcell.KeyUp:
                    inputCh <- MoveUp

                case tcell.KeyDown:
                    inputCh <- MoveDown

                case tcell.KeyLeft:
                    inputCh <- MoveLeft

                case tcell.KeyRight:
                    inputCh <- MoveRight

                case tcell.KeyRune:
                    if e.Rune() == ' ' {
                        inputCh <- PickUp
                    }

                    if e.Rune() == '1' {
                        inputCh <- DropLight
                    }

                    if e.Rune() == '2' {
                        inputCh <- DropTrap
                    }

                case tcell.KeyEscape, tcell.KeyCtrlC:
                    close(quitCh)
                    return
                }

            case *tcell.EventResize:
                screen.Sync()
            }
        }
    }()

}

