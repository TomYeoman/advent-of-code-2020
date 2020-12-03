package day3

import (
	"fmt"
	"math/rand"
	"time"
)

// ------------ PRIMTIIVE API ------------- //

type Tobogganer struct {
	x              int
	y              int
	positionsRight int
	positionsDown  int
	treesCollided  int
}

func (t *Tobogganer) rideSlope(slope Slope, done chan bool) {
	fmt.Println("Riding Slope 🎿")
	time.Sleep(time.Second)
	done <- true
}

type Slope struct {
	terrain [][]string
}

// generateTerrain accepts a tettain map, and builds a useable 2D matrix
func (s *Slope) generateTerrain(source []string) {
	dataGrid := make([][]string, len(source))
	for x, row := range source {
		dataGrid[x] = make([]string, len(row))
		for y, rowEntry := range row {
			dataGrid[x][y] = string(rowEntry)
		}
	}
	s.terrain = dataGrid
}

// ------------ LOW LEVEL API ------------- //

func rideSlope(slope Slope, t []Tobogganer) {
	done := make(chan bool)

	for {
		rider := rand.Intn(len(t))
		fmt.Printf("Sending rider %d down the slope \n", rider)

		go t[rider].rideSlope(slope, done)
		<-done
	}
}

// ------------ HIGH LEVEL API ------------- //

// Skipark wraps our rider, and slopes into a single entity.
type SkiPark struct {
	slope        Slope
	tobogannists []Tobogganer
}

// This method will now be responsible, for
func OpenPark(skipark SkiPark, slopeSeed []string) error {
	skipark.slope.generateTerrain(slopeSeed)
	rideSlope(skipark.slope, skipark.tobogannists)
	return nil
}

func RideSlopeV2(entries []string) error {
	park := SkiPark{
		slope: Slope{
			terrain: [][]string{},
		},
		tobogannists: []Tobogganer{
			{
				x: 0, y: 0, positionsDown: 1, positionsRight: 1, treesCollided: 0,
			},
			{
				x: 0, y: 0, positionsDown: 5, positionsRight: 10, treesCollided: 0,
			},
		},
	}

	if err := OpenPark(park, entries); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

/*
	Requirements (V1)
		- Every 5 minutes, we're going to start a game which has a slope,
		and a bunch of tobannists wanting to ride a slope

	Step 1 - Define the API on the concrete implementation. We're optimise for correctness on first pass - defining the concrete types

		Primitive API
		--------------
		As we're dealing with state we can use a type based API
			- A tobaggan
				- `rideSlope` method uses pointer semantics, and checks a riders path down
			- A Slope
				- `generateTerrain` method uses pointer semantics, to build a useable Grid of terrain

		We could now use these alone, but we may add one more entity
			A Ski park - wraps our rider, and slopes into a single entity. It's API may form
			/ part of a higher level API

		Low Level API
		--------------
		Now we move onto the low level API. As soon as we're not dealing with state we can move back to a function based API.

		The low level API provides a wrapper over the primitive. Maybe we only want 1 rider to be able to run at a time, in our case
			- `rideSlope` is responsible for sending 1 rider at time down

		High Level API
		--------------
		These types of API deal with multiple behaviours. In our case our requirements say we
		"wish to start a game every X5minutes, that has a slope and riders". Lets define a high level API to help with this -
			- OpenPark - It's a function based API, that is going to be responsible for building terrains, and sending riders down them,

*/
