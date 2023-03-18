package utils

import (
	"camuschino/laberth-go/models"
	"math/rand"
	"time"
)

func generateRandInt(randLimit int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(randLimit)
}

func generateRandMapPoint(randX, randY int) models.Coords {
	return models.Coords{
		XPoint: generateRandInt(randX),
		YPoint: generateRandInt(randY),
	}
}

func GenerateValidMapPoint(laberth *models.Labyrinth) (mapPoint models.Coords) {

	randX := len(laberth.ArrayToMap) - 2
	randY := len(laberth.ArrayToMap[0]) - 2

	for {
		mapPoint = generateRandMapPoint(randX, randY)
		switch mapPointable := laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint].(type) {
		case models.MapBool:
			if mapPointable {
				continue
			}
		case models.MapPointable:
		default:
			continue
		}
		break
	}

	return
}

func getNewTarget(mapPoint models.Coords, newTarget models.Target) models.MapPointable {
	var score int

	switch newTarget.(type) {
	case models.Enemy:
		score = 20
	case models.Coin:
	default:
		score = 10
	}

	return models.MapPoint{
		TargetInPoint: newTarget.SetScore(score).SetMapPoint(mapPoint),
	}
}

// SetObjectPositions func
func SetObjectPositions(laberth *models.Labyrinth) (player, target models.Coords) {

	player = GenerateValidMapPoint(laberth)
	target = GenerateValidMapPoint(laberth)
	var mapPoint models.Coords

	for i := 0; i < 10; i++ {
		mapPoint = GenerateValidMapPoint(laberth)
		switch generateRandInt(2) {
		case 0:
			laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint] = getNewTarget(mapPoint, models.Enemy{})
		case 1:
			laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint] = getNewTarget(mapPoint, models.Coin{})
		}
	}

	return
}

// Create new Labyrinth by setting point and walls in empty map
func CreateNewLabyrinth(emptyMap *models.Labyrinth) {
	// Iterate in map using dimentions
	iterateIntoEmptyMap(emptyMap)
	// Creating empty map to check movements
	createEmptyMapToCheck(emptyMap)
}


func iterateIntoEmptyMap(emptyMap *models.Labyrinth) {
	// Getting map dimesions to iterate
	fieldDimentionX, fieldDimentionY := GetMapDimensions(emptyMap)
	// Getting new randomizer to generate walls
	randomizer := getRandomizer()
	for i := 0; i < fieldDimentionX; i++ {
		for j := 0; j < fieldDimentionY; j++ {
			emptyMap.ArrayToMap[i][j] = getRandomBool(randomizer, i, j)
		}
	}
}

func getRandomBool(randomizer *rand.Rand, i, j int) models.MapBool {
	return !((i%2 == 0 && j%2 == 0) || randomizer.Intn(2) == 0)
}

func createEmptyMapToCheck(emptyMap *models.Labyrinth) {
	// Getting map dimesions to iterate and create empty arrays to check
	fieldDimentionX, fieldDimentionY := GetMapDimensions(emptyMap)

	// making new arrays to check movements
	emptyMap.ArrayToCheck = make([][]bool, fieldDimentionX)
	for i := range emptyMap.ArrayToCheck {
		emptyMap.ArrayToCheck[i] = make([]bool, fieldDimentionY)
	}
}

func getRandomizer() (randomizer *rand.Rand) {
	timeFactor := rand.NewSource(time.Now().UnixNano())
	randomizer = rand.New(timeFactor)
	return
}

// Get dimension of a map
//	laberth: Map to get dimensions
func GetMapDimensions(laberth *models.Labyrinth) (int, int) {
	return len(laberth.ArrayToMap), len(laberth.ArrayToMap[0])
}
