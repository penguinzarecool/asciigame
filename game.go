package main

import (
    "math"
)

type Game struct {
	Width, Height int
	PlayerX, PlayerY int
    EnemyX, EnemyY int
    CrystalX, CrystalY int
    CrystalMove bool
    CrystalStolen bool
    eCrystalX, eCrystalY int
    eCrystalMove bool
    Lights []Light
    eLights []eLight
    Traps []Trap
    eTraps []eTrap
    NumLights, NumTraps int
}

type Light struct {
    X, Y int
}

type eLight struct {
    X, Y int
}

type Trap struct {
    X, Y int
}

type eTrap struct {
    X, Y int
}

func NewGame(w, h int) *Game {
	return &Game{
		Width:  w,
		Height: h,
		PlayerX: 0,
		PlayerY: h / 2,
        EnemyX: w - 5,
        EnemyY: h / 2,
        CrystalX: w / 6,
        CrystalY: h / 2,
        CrystalMove: false,
        CrystalStolen: false,
        eCrystalX: w / 6 * 5,
        eCrystalY: h / 2,
        eCrystalMove: false,
        NumLights: 5,
        NumTraps: 25,
	}
}

func (g *Game) ApplyInput(input Input) {

    // test game
    g.enemyTrap()
    g.enemyLight()
    //g.CrystalStolen = true

	switch input {
    case PickUp:
        dx := math.Abs(float64(g.CrystalX - g.PlayerX))
        dy := math.Abs(float64(g.CrystalY - g.PlayerY))

        diffX := int(dx)
        diffY := int(dy)
        
        if ( ( diffX <= 1 ) && ( diffY <= 1 ) ) {
            g.CrystalMove = !g.CrystalMove
        }

        dx = math.Abs(float64(g.eCrystalX - g.PlayerX))
        dy = math.Abs(float64(g.eCrystalY - g.PlayerY))

        diffX = int(dx)
        diffY = int(dy)
        
        if ( ( diffX <= 1 ) && ( diffY <= 1 ) ) {
            g.eCrystalMove = !g.eCrystalMove
        }


    case DropLight:
        g.placeLight()
    case DropTrap:
        g.dropTrap()
	case MoveUp:
		g.move(0, -1)
        if (g.CrystalMove) {
            g.CrystalY -= 1
        }
        if (g.eCrystalMove) {
            g.eCrystalY -= 1
        }
	case MoveDown:
		g.move(0, 1)
        if (g.CrystalMove) {
            g.CrystalY += 1
        }
        if (g.eCrystalMove) {
            g.eCrystalY += 1
        }
	case MoveLeft:
		g.move(-1, 0)
        if (g.CrystalMove) {
            g.CrystalX -= 1
        }
        if (g.eCrystalMove) {
            g.eCrystalX -= 1
        }
	case MoveRight:
		g.move(1, 0)
        if (g.CrystalMove) {
            g.CrystalX += 1
        }
        if (g.eCrystalMove) {
            g.eCrystalX += 1
        }
	}
}

func (g *Game) move(dx, dy int) {
	nx := g.PlayerX + dx
	ny := g.PlayerY + dy

	if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
		g.PlayerX = nx
		g.PlayerY = ny
	}

    if g.CrystalX < 1 || g.CrystalX > g.Width - 1 || g.CrystalY < 1 || g.CrystalY > g.Height - 1 {
        g.CrystalX = g.PlayerX
        g.CrystalY = g.PlayerY
    }
    if g.eCrystalX < 1 || g.eCrystalX > g.Width - 1 || g.eCrystalY < 1 || g.eCrystalY > g.Height - 1 {
        g.eCrystalX = g.PlayerX
        g.eCrystalY = g.PlayerY
    }
}

func (g *Game) placeLight() {

    if g.NumLights == 0 {
        return
    }

    light := Light {
        X: g.PlayerX,
        Y: g.PlayerY,
    }

    g.Lights = append(g.Lights, light)
    g.NumLights -= 1
}

func (g *Game) enemyLight() {

    elight := eLight {
        X: 110,
        Y: g.Height / 4,
    }

    g.eLights = append(g.eLights, elight)
}

func (g *Game) dropTrap() {
    if g.NumTraps == 0 {
        return
    }

    trap := Trap {
        X: g.PlayerX,
        Y: g.PlayerY,
    }

    g.Traps = append(g.Traps, trap)
    g.NumTraps -= 1
}

func (g *Game) enemyTrap() {

    etrap := eTrap {
        X: 110,
        Y: g.Height / 2,
    }

    g.eTraps = append(g.eTraps, etrap)
}
