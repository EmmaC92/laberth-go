package models

import "sync"

// Labyrinth struct
type Labyrinth struct {
	SizeField        int
	MovementDistance float32
	ArrayToCheck     [][]bool
	ArrayToMap       [][]MapPointable
	Over             chan bool
	Mu               sync.Mutex
}

// SetArrays func
func (laberth *Labyrinth) CreateNewEmptyMap(fieldDimentionX, fieldDimentionY int) {
	laberth.ArrayToCheck = make([][]bool, fieldDimentionX)
	laberth.ArrayToMap = make([][]MapPointable, fieldDimentionX)

	laberth.Over = make(chan bool)

	for i := range laberth.ArrayToCheck {
		laberth.ArrayToCheck[i] = make([]bool, fieldDimentionY)
		laberth.ArrayToMap[i] = make([]MapPointable, fieldDimentionY)
	}
}
