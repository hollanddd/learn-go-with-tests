package clockface

import "time"

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{X: 0, Y: 0}
}
