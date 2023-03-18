package utils

import (
	"camuschino/laberth-go/models"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	nextPosition models.Coords
)

func CheckTargetPosition(win *pixelgl.Window, imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	imd.Clear()

	// Checking up key is pressed
	checkUpKey(win, imd, laberth, target)
	// Checking down key is pressed
	checkDownKey(win, imd, laberth, target)
	// Checking right key is pressed
	checkRightKey(win, imd, laberth, target)
	// Checking left key is pressed
	checkLeftKey(win, imd, laberth, target)

	DrawObjectToRender(imd, *target, colornames.Red, laberth)
	imd.Draw(win)
}

func checkUpKey(win *pixelgl.Window, imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	if win.JustPressed(pixelgl.KeyUp) {
		moveObjectUp(imd, laberth, target)
	}
}

func moveObjectUp(imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	nextPosition = *target
	nextPosition.YPoint++
	checkAndMoveObject(nextPosition, target, nextPosition.YPoint, len(laberth.ArrayToMap[0]), imd, laberth)
}

func checkDownKey(win *pixelgl.Window, imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	if win.JustPressed(pixelgl.KeyDown) {
		moveObjectDown(imd, laberth, target)
	}
}

func moveObjectDown(imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	nextPosition = *target
	nextPosition.YPoint--
	checkAndMoveObject(nextPosition, target, nextPosition.YPoint, len(laberth.ArrayToMap[0]), imd, laberth)
}

func checkRightKey(win *pixelgl.Window, imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	if win.JustPressed(pixelgl.KeyRight) {
		moveObjectRight(imd, laberth, target)
	}
}

func moveObjectRight(imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	nextPosition = *target
	nextPosition.XPoint++
	checkAndMoveObject(nextPosition, target, nextPosition.XPoint, len(laberth.ArrayToMap[0]), imd, laberth)
}

func checkLeftKey(win *pixelgl.Window, imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	if win.JustPressed(pixelgl.KeyLeft) {
		moveObjectLeft(imd, laberth, target)
	}
}

func moveObjectLeft(imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	nextPosition = *target
	nextPosition.XPoint--
	checkAndMoveObject(nextPosition, target, nextPosition.XPoint, len(laberth.ArrayToMap[0]), imd, laberth)
}

func checkAndMoveObject(nextPosition models.Coords, target *models.Coords, positionToCheck, limit int, imd *imdraw.IMDraw, laberth *models.Labyrinth) {
	// Check if the target going to move out the map limit or a wall-point
	if CheckLimit(positionToCheck, limit) && CheckMapPoint(nextPosition, laberth) {
		DrawObjectToRender(imd, *target, colornames.Black, laberth)
		*target = nextPosition
	}
}
