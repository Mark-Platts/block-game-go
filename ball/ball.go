package ball

import (
	"github.com/hajimehoshi/ebiten/v2"
	"own.com/graphical-engine-tests/config"
)

type Ball struct {
	x_pos float64
	y_pos float64
	x_vel float64
	y_vel float64
}

var B *Ball

func Init() {
	B = &Ball{
		x_pos: 20,
		y_pos: 60,
		x_vel: 5,
		y_vel: 5,
	}
}

func UpdateState() {
	//Get mouse position
	cursorPositionX, _ := ebiten.CursorPosition()

	//calculate the next position
	newXPos := B.x_pos + B.x_vel
	newYPos := B.y_pos + B.y_vel

	//check if there are wall collisions and update accordingly
	if newXPos <= 0 {
		newXPos = 0
		FlipXVel()
	} else if newXPos >= float64(config.ScreenSizeX-(config.BallDiameter/2)) {
		newXPos = float64(config.ScreenSizeX - (config.BallDiameter / 2))
		FlipXVel()
	}
	if newYPos >= float64(config.ScreenSizeY-config.PaddleClearance-config.BallDiameter) && newXPos >= float64(cursorPositionX-config.PaddleLength) && newXPos <= float64(cursorPositionX+config.PaddleLength) {
		newYPos = float64(config.ScreenSizeY - config.PaddleClearance - config.BallDiameter)
		FlipYVel()
	} else if newYPos <= 0 {
		newYPos = 0
		FlipYVel()
	} else if newYPos >= float64(config.ScreenSizeY-(config.BallDiameter/2)) {
		newXPos = 10
		newYPos = 10
	}

	B.x_pos = newXPos
	B.y_pos = newYPos
}

func FlipXVel() {
	B.x_vel = -1 * B.x_vel
}

func FlipYVel() {
	B.y_vel = -1 * B.y_vel
}

func GetXPos() float64 {
	return B.x_pos
}

func GetYPos() float64 {
	return B.y_pos
}

func GetXVel() float64 {
	return B.x_vel
}

func GetYVel() float64 {
	return B.y_vel
}
