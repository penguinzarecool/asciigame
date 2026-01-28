package main

import (
    "math"
    "github.com/gdamore/tcell/v2"
)

func Render(screen tcell.Screen, g *Game) {
    screen.Clear()

    bgStyle := tcell.StyleDefault.
        Foreground(tcell.NewRGBColor(1, 18, 10))
    
    playerStyle := tcell.StyleDefault.
        Foreground(tcell.ColorGreen).
        Bold(true)

    //enemyStyle := tcell.StyleDefault.
        //Foreground(tcell.ColorRed).
        //Foreground(tcell.NewRGBColor(1, 18, 10)).
        //Bold(true)

    lightStyle := tcell.StyleDefault.
        Foreground(tcell.ColorWhite).
        Bold(true)

    trapStyle := tcell.StyleDefault.
        Foreground(tcell.ColorRed).
        Bold(true)

    crystalStyle := tcell.StyleDefault.
        Foreground(tcell.ColorWhite).
        Bold(true)

    for y := 0; y < g.Height; y++ {

        for x := 0; x < g.Width; x++ {

            screen.SetContent(x, y, '^', nil, bgStyle)

            // ------- player ------- 

            if ( x < ( g.Width / 2 ) ) {

                dx := math.Abs(float64(x - g.PlayerX))
                dy := math.Abs(float64(y - g.PlayerY))

                diffX := int(dx)
                diffY := int(dy)
                
                var colorDiff int = 0

                if ( diffX < 6 && diffY < 5 ) && ( (diffX + diffY < 8 ) ) {
                    colorDiff = 70;
                    screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(int32(80 - colorDiff), int32(114 - colorDiff), int32(92 - colorDiff))))
                    if g.EnemyX == x && g.EnemyY == y {
                        screen.SetContent(x, y, '^', nil, trapStyle )
                    }
                }
                if ( ( diffX < 4 && diffY < 3 ) ) {
                    colorDiff = 30;
                    screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(int32(70 - colorDiff), int32(94 - colorDiff), int32(72 - colorDiff))))
                    if g.EnemyX == x && g.EnemyY == y {
                        screen.SetContent(x, y, '^', nil, trapStyle )
                    }
                }

            }

            if ( x >= ( g.Width / 2 ) ) {

                dx := math.Abs(float64(x - g.PlayerX))
                dy := math.Abs(float64(y - g.PlayerY))

                diffX := int(dx)
                diffY := int(dy)
                
                var colorDiff int = 0


                if ( ( diffX < 6 && diffY < 5 ) && ( (diffX + diffY < 8 ) ) ) {
                    colorDiff = 70;
                    screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(int32(74 - colorDiff), int32(89 - colorDiff), int32(135 - colorDiff))))
                    if g.EnemyX == x && g.EnemyY == y {
                        screen.SetContent(x, y, '^', nil, trapStyle )
                    }
                }
                if ( ( diffX < 4 && diffY < 3 ) ) {
                    colorDiff = 30;
                    screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(int32(74 - colorDiff), int32(89 - colorDiff), int32(135 - colorDiff))))
                    if g.EnemyX == x && g.EnemyY == y {
                        screen.SetContent(x, y, '^', nil, trapStyle )
                    }
                }

                //if g.EnemyX == x && g.EnemyY == y {
                    //screen.SetContent(x, y, '^', nil, trapStyle )
                //}
            }
            // ------- crystal ------- 

            dx := math.Abs(float64(x - g.CrystalX))
            dy := math.Abs(float64(y - g.CrystalY))

            diffX := int(dx)
            diffY := int(dy)

            dxp := math.Abs(float64(g.PlayerX - x))
            dyp := math.Abs(float64(g.PlayerY - y))

            diffXp := int(dxp)
            diffYp := int(dyp)
                
            if !g.CrystalStolen || ( g.CrystalStolen && ( diffXp < 6 && diffYp < 5 ) ) {
                if ( ( diffX + diffY ) == 3 ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(07, 97, 52)))
                }
                
                if ( ( diffX + diffY ) == 2 ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(117, 217, 172)))
                }
                
                if ( diffX == 1 && diffY == diffX - 1 ) || ( diffY == 1 && diffX == diffY - 1 ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(117, 217, 172)))
                }
                    
                if ( diffX == 0 && diffY == 0 ) {
                        screen.SetContent(x, y, '@', nil, crystalStyle)
                }
            }

            // ------- enemy crystal ------- 

            dx = math.Abs(float64(x - g.eCrystalX))
            dy = math.Abs(float64(y - g.eCrystalY))

            diffX = int(dx)
            diffY = int(dy)

            dxp = math.Abs(float64(g.PlayerX - x))
            dyp = math.Abs(float64(g.PlayerY - y))

            diffXp = int(dxp)
            diffYp = int(dyp)
                
            if diffXp < 6 && diffYp < 5 {
                if ( ( diffX + diffY ) == 3 ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(97, 16, 11)))
                }
                
                if ( ( diffX + diffY ) == 2 ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(247, 83, 72)))
                }
                
                if ( diffX == 1 && diffY == diffX - 1 ) || ( diffY == 1 && diffX == diffY - 1 ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(242, 170, 165)))
                }
                
                if ( diffX == 0 && diffY == 0 ) {
                        screen.SetContent(x, y, '@', nil, crystalStyle)
                }
            }

            // ------- light ------- 

            for _, light := range g.Lights {
                
                dx = math.Abs(float64(x - light.X))
                dy = math.Abs(float64(y - light.Y))

                diffX = int(dx)
                diffY = int(dy)
                
                if ( ( diffX < 8 ) && ( diffY < 6 ) && ( diffX + diffY < 10 ) ) {
                    screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(87, 83, 16)))
                    if g.EnemyX == x && g.EnemyY == y {
                        screen.SetContent(x, y, '^', nil, trapStyle )
                    }
                }

                if ( ( diffX < 6 ) && ( diffY < 4 ) && ( diffX + diffY < 6 ) ) {
                    screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(143, 124, 33)))
                    if g.EnemyX == x && g.EnemyY == y {
                        screen.SetContent(x, y, '^', nil, trapStyle )
                    }
                }

                if ( ( diffX < 2 ) && ( diffY < 2 ) && ( diffX + diffY < 4 ) ) {
                    screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(253, 167, 42)))
                    if g.EnemyX == x && g.EnemyY == y {
                        screen.SetContent(x, y, '^', nil, trapStyle )
                    }
                }
            }
            
            // ------- enemy light -------

            for _, elight := range g.eLights {
               
                // measures from x to light 
                dx = math.Abs(float64(x - elight.X))
                dy = math.Abs(float64(y - elight.Y))

                diffX = int(dx)
                diffY = int(dy)
                
                // measures from x to player
                dxp = math.Abs(float64(g.PlayerX - x))
                dyp = math.Abs(float64(g.PlayerY - y))

                diffXp = int(dxp)
                diffYp = int(dyp)
                
                if diffXp < 6 && diffYp < 5 {
                    if ( ( diffX < 8 ) && ( diffY < 6 ) && ( diffX + diffY < 10 ) ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(87, 83, 16)))
                        if g.EnemyX == x && g.EnemyY == y {
                            screen.SetContent(x, y, '^', nil, trapStyle )
                        }
                    }

                    if ( ( diffX < 6 ) && ( diffY < 4 ) && ( diffX + diffY < 6 ) ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(143, 124, 33)))
                        if g.EnemyX == x && g.EnemyY == y {
                            screen.SetContent(x, y, '^', nil, trapStyle )
                        }
                    }

                    if ( ( diffX < 2 ) && ( diffY < 2 ) && ( diffX + diffY < 4 ) ) {
                        screen.SetContent(x, y, '^', nil, tcell.StyleDefault.Foreground(tcell.NewRGBColor(253, 167, 42)))
                        if g.EnemyX == x && g.EnemyY == y {
                            screen.SetContent(x, y, '^', nil, trapStyle )
                        }
                    }
                }
            }
        }
    }

// ------- individual sprites -------

    for _, light := range g.Lights {
        screen.SetContent(
            light.X,
            light.Y,
            'O',
            nil,
            lightStyle,
        )
    }

    for _, trap := range g.Traps {
        screen.SetContent(
            trap.X,
            trap.Y,
            'X',
            nil,
            trapStyle,
        )
    }

    for _, etrap := range g.eTraps {
        screen.SetContent(
            etrap.X,
            etrap.Y,
            'X',
            nil,
            bgStyle,
        )
    }

    screen.SetContent(
        g.PlayerX,
        g.PlayerY,
        '^',
        nil,
        playerStyle,
    )

//    screen.SetContent(
        //g.EnemyX,
        //g.EnemyY,
        //'^',
        //nil,
        //bgStyle,
        //enemyStyle,
    //)

    //screen.SetContent(
        //g.CrystalX,
        //g.CrystalY,
        //'@',
        //nil,
        //crystalStyle,
    //)

    //screen.SetContent(
        //g.eCrystalX,
        //g.eCrystalY,
        //'@',
        //nil,
        //crystalStyle,
    //)

    screen.Show()
}

