package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"own.com/graphical-engine-tests/ball"
	"own.com/graphical-engine-tests/config"
)

var (
	ballImg   *ebiten.Image
	paddleImg *ebiten.Image
	err       error
)

type Game struct{}

func (g *Game) Update() error {
	ball.UpdateState()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cursorPositionX, _ := ebiten.CursorPosition()

	ballOp := &ebiten.DrawImageOptions{}
	ballOp.GeoM.Translate(ball.GetXPos(), ball.GetYPos())

	paddleOp := &ebiten.DrawImageOptions{}
	paddleOp.GeoM.Translate(float64(cursorPositionX-config.PaddleLength/2), float64(config.ScreenSizeY-config.PaddleClearance))

	// ebitenutil.DebugPrintAt(screen, "o", int(math.Round(ball.GetXPos())), int(math.Round(ball.GetYPos())))
	screen.DrawImage(ballImg, ballOp)
	screen.DrawImage(paddleImg, paddleOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenSizeX, config.ScreenSizeY
}

func main() {
	ball.Init()
	ballImg, _, err = ebitenutil.NewImageFromFile("./img/ball.png")
	if err != nil {
		log.Fatal(err)
	}
	paddleImg, _, err = ebitenutil.NewImageFromFile("./img/paddle.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(880, 660)
	ebiten.SetWindowTitle("Golang Graphics Engine Test")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
