package utils

import (
	"camuschino/laberth-go/models"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	WINDOW_DIMENTION_X, WINDOW_DIMENTION_Y, SIZE_BLOCK int     = 700, 700, 50 // window dimention AND large: 100. medium: 50, little: 20, nano: 10
	FIELD_DIMENTION_X, FIELD_DIMENTION_Y               int     = ((WINDOW_DIMENTION_X / SIZE_BLOCK) * 2) + 1, ((WINDOW_DIMENTION_Y / SIZE_BLOCK) * 2) + 1
	SIZE_FIELD                                         int     = SIZE_BLOCK / 2
	MOVEMENT_DISTANCE                                  float32 = float32(SIZE_FIELD)
)

func GetWindowConfigs() pixelgl.WindowConfig {
	return pixelgl.WindowConfig{
		Title:  "Laberth",
		Bounds: pixel.R(0, 0, float64(WINDOW_DIMENTION_Y), float64(WINDOW_DIMENTION_X)),
		VSync:  true,
	}
}

func GetWindow() *pixelgl.Window {
	win, err := pixelgl.NewWindow(GetWindowConfigs())
	if err != nil {
		panic(err)
	}
	return win
}

func GetImd() *imdraw.IMDraw {
	return imdraw.New(nil)
}

func GetNewEmptyMap() models.Labyrinth {
	laberth := models.Labyrinth{
		SizeField:        SIZE_FIELD,
		MovementDistance: MOVEMENT_DISTANCE,
	}
	laberth.CreateNewEmptyMap(FIELD_DIMENTION_X, FIELD_DIMENTION_Y)
	return laberth
}

func CreateNewObjectsInLaberth(laberth *models.Labyrinth) (models.Coords, models.Coords) {

	player, target := SetObjectPositions(laberth)
	return player, target
}

func BuildLaberthWall(newEmptyMap *models.Labyrinth) *models.Labyrinth {
	CreateNewLabyrinth(newEmptyMap)
	return newEmptyMap
}

func MovementThread(win *pixelgl.Window, imd *imdraw.IMDraw, target *models.Coords, laberth *models.Labyrinth) {
	for {
		time.Sleep(10 * time.Millisecond)
		win.Update()
		CheckTargetPosition(win, imd, laberth, target)
	}
}
