package utils

import (
	"camuschino/laberth-go/models"
	"math/rand"
	"time"
)

var (
	randomizer *rand.Rand
)

func GenerateValidMapPoint(laberth *models.Labyrinth) (mapPoint models.Coords) {

	for {
		mapPoint = generateRandMapPoint(laberth.ArrayToMap)
		if checkMapPointIsAvailable(mapPoint, laberth) {
			break
		}
	}

	return
}

func checkMapPointIsAvailable(mapPoint models.Coords, laberth *models.Labyrinth) bool {
	switch mapPointable := laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint].(type) {
	case models.MapBool:
		if mapPointable {
			return false
		}
	case models.MapPointable:
	default:
		return false
	}
	return true
}

func generateRandMapPoint(arrayMap [][]models.MapPointable) models.Coords {
	randX := len(arrayMap) - 2
	randY := len(arrayMap[0]) - 2
	return models.Coords{
		XPoint: generateRandInt(randX),
		YPoint: generateRandInt(randY),
	}
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

	var mapPoint models.Coords
	var newTarget models.Target

	for i := 0; i < 10; i++ {
		mapPoint = GenerateValidMapPoint(laberth)

		if generateRandBool() {
			newTarget = models.Coin{}
		} else {
			newTarget = models.Enemy{}
		}

		laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint] = getNewTarget(mapPoint, newTarget)
	}

	return
}

func createNewEnemy() {

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
	// Randomizer
	randomizer = getRandomizer()

	for i := 0; i < fieldDimentionX; i++ {
		for j := 0; j < fieldDimentionY; j++ {
			emptyMap.ArrayToMap[i][j] = getRandomMapBool(i, j)
		}
	}
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

func getRandomMapBool(XPoint, YPoint int) models.MapBool {
	return (models.MapBool)(generateRandBool() && (XPoint%2 == 0 || YPoint%2 == 0))
}

func generateRandBool() bool {
	return generateRandInt(2) == 1
}

func generateRandInt(limit int) int {
	return randomizer.Intn(limit)
}

func getRandomizer() *rand.Rand {
	timeFactor := rand.NewSource(time.Now().UnixNano())
	return rand.New(timeFactor)
}

// Get dimension of a map
//	laberth: Map to get dimensions
func GetMapDimensions(laberth *models.Labyrinth) (int, int) {
	return len(laberth.ArrayToMap), len(laberth.ArrayToMap[0])
}
