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
	snake  := models.Coords{}

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
	go utils.MovementThread(win, targetImd, &snake, &emptyMap)

	// building walls
	log.Println("Building walls..")
	laberth := utils.BuildLaberthWall(&emptyMap)

	// setting player and snake
	log.Println("Setting players and snake..")
	player, snake = utils.CreateNewObjectsInLaberth(laberth)

	// setting targets
	log.Println("Setting targets..")	
	utils.SetTargetPositionsInLabyrinth(laberth)

	// Wash and print maps
	log.Println("Preparing maps..")
	win.Clear(colornames.Black)
	imd.Clear()

	// Printing walls
	log.Println("Printing walls and objects")
	utils.RenderMapAndObjects(laberth, imd)
	imd.Draw(win)
	
	// // Validating maps
	log.Println("Validating laberyn..")
	utils.ValidateMap("", player, &snake, laberth, imd, win, utils.SIZE_FIELD)

	// starting infinite for loop to draw the maps and object  
	log.Println("Starting infinite loop to draw.. ")
	// for {

	// }
}

func main() {
	log.Println("Initializing laberth.. ")
	pixelgl.Run(run)
}
