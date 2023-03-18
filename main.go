package main

import (
	"camuschino/laberth-go/models"
	"camuschino/laberth-go/utils"
	"log"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	// initializing objects
	player := models.Coords{}
	target := models.Coords{}

	// getting configurations for windows
	log.Println("Getting configurations.. ")
	win := utils.GetWindow()
	imd := utils.GetImd()

	// generate empty maps
	emptyMap := utils.GetNewEmptyMap()

	// initializing new drawer
	targetImd := imdraw.New(nil)

	// run main go routine to check movements from user
	log.Println("Running movement thread.. ")
	go utils.MovementThread(win, targetImd, &target, &emptyMap)

	// starting infinite for loop to draw the maps and object  
	log.Println("Starting infinite loop to draw.. ")
	for {
		laberth := utils.BuildLaberthWall(&emptyMap)
		player, target = utils.SetObjectsInLaberth(laberth)
		win.Clear(colornames.Black)
		imd.Clear()
		utils.RenderMapAndObjects(laberth, imd)
		utils.ValidateMap("", player, &target, laberth, imd, win, utils.SIZE_FIELD)
	}
}

func main() {
	log.Println("Initializing laberth.. ")
	pixelgl.Run(run)
}
