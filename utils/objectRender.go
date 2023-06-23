package utils

import (
	"camuschino/laberth-go/models"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	ANTIQUEWHITE_COLOR color.RGBA = colornames.Antiquewhite
	WHITE_COLOR        color.RGBA = colornames.White
)

// RenderMapAndObjects func
func RenderMapAndObjects(laberth *models.Labyrinth, imd *imdraw.IMDraw) {
	// Getting map dimesions to render walls and targets
	fieldDimentionX, fieldDimentionY := GetMapDimensions(laberth)

	// Iterate  map to render all walls and targets
	for current_x_point := 0; current_x_point < fieldDimentionX; current_x_point++ {
		for current_y_point := 0; current_y_point < fieldDimentionY; current_y_point++ {

			// getting map point to check with swtich type how it should be rendered
			mapPoint := laberth.ArrayToMap[current_x_point][current_y_point]
			switch mapPointable := mapPoint.(type) {
				// map point is a normal point.
				// wall or non-wall
				case models.MapBool:
					if mapPointable {
						pixelPoint := getWall(current_x_point, current_y_point, laberth.SizeField)
						renderBoolPoint(pixelPoint, imd)
					}
				// map is a target point
				// coin, enemy, etc.
				case models.MapPoint:
					target := mapPointable.TargetInPoint
					renderTargetPoint(target, imd, laberth)
			}
		}
	}
}

func renderBoolPoint(pixelPoint pixel.Rect, imd *imdraw.IMDraw) {
	imd.Color = WHITE_COLOR
	imd.Push(pixelPoint.Min, pixelPoint.Max)
	imd.Rectangle(0)
}

func renderTargetPoint(target models.Target, imd *imdraw.IMDraw, laberth *models.Labyrinth) {
	switch target.(type) {
	case models.Coin:
		DrawObjectToRender(imd, target.GetMapPoint(), colornames.Blue, laberth)
	case models.Enemy:
		DrawObjectToRender(imd, target.GetMapPoint(), colornames.Greenyellow, laberth)
	}
}

func getWall(x, y, sizeField int) (px pixel.Rect) {
	posX := float64(x * sizeField)
	posY := float64(y * sizeField)
	px = pixel.R(posX, posY, posX+float64(sizeField), posY+float64(sizeField))
	return
}

func DrawObjectToRender(imd *imdraw.IMDraw, object models.Coords, color color.Color, laberth *models.Labyrinth) {
	objectToRender := pixel.Vec{
		X: float64((object.XPoint * laberth.SizeField) + int(laberth.MovementDistance/2)),
		Y: float64((object.YPoint * laberth.SizeField) + int(laberth.MovementDistance/2)),
	}

	imd.Color = color
	imd.Push(objectToRender)
	imd.Circle(float64(laberth.MovementDistance/2), 0)
}

func RenderingStep(point models.Coords, laberth *models.Labyrinth, color color.Color, imd *imdraw.IMDraw, win *pixelgl.Window) {
	laberth.Mu.Lock()
	imd.Color = color
	px := getWall(point.XPoint, point.YPoint, laberth.SizeField)
	imd.Push(px.Min, px.Max)
	imd.Rectangle(0)
	imd.Draw(win)
	laberth.Mu.Unlock()
}
