package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/jamiecrisman/grid-puzzle/state"
)

func main() {
	startPos := 0
	is := state.CreateInitialState(10, startPos)
	fmt.Println("Start")
	start := time.Now()
	result := play(*is)
	elapsed := time.Since(start)
	if result.IsComplete() {
		fmt.Printf("Solved in %s\n", elapsed)
		fmt.Println(result.WinPath())
		fmt.Printf("validation: %v\n", validate(*result, startPos))
	}

}

func play(s state.State) *state.State {
	if s.IsComplete() {
		return &s
	}
	m := s.GetMoves()
	sort.Sort(state.ByScore(m))
	for i := 0; i < len(m); i++ {
		j := play(m[i])
		if j != nil {
			return j
		}
	}
	return nil
}

func validate(s state.State, startingPosition int) bool {
	nextOffset := s.Values[startingPosition]  // first value
	position := startingPosition + nextOffset // second index
	steps := make([]int, len(s.Values))
	count := 1
	steps[startingPosition] = count
	for i := 0; i < len(s.Values)-1; i++ {
		if !state.Valid(position, len(s.Values)) || steps[position] != 0 {
			return false
		}
		nextOffset = s.Values[position]
		count++
		steps[position] = count
		position = position + nextOffset
	}
	fmt.Println(steps)
	fmt.Println("position, value: ", position, s.Values[position])
	return true
}
