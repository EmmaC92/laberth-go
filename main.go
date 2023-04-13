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

	// building walls
	log.Println("Building walls..")
	laberth := utils.BuildLaberthWall(&emptyMap)

	// setting players and targets
	log.Println("Setting players and targets..")
	player, target = utils.CreateNewObjectsInLaberth(laberth)

	// Wash and print maps
	log.Println("Preparing maps..", player)
	win.Clear(colornames.Black)
	imd.Clear()

	// Printing walls
	log.Println("Printing walls and objects")
	utils.RenderMapAndObjects(laberth, imd)
	imd.Draw(win)
	
	// // Validating maps
	// log.Println("Validating laberyn..")
	// utils.ValidateMap("", player, &target, laberth, imd, win, utils.SIZE_FIELD)

	// starting infinite for loop to draw the maps and object  
	log.Println("Starting infinite loop to draw.. ")
	for {

	}
}

func main() {
	log.Println("Initializing laberth.. ")
	pixelgl.Run(run)
}
