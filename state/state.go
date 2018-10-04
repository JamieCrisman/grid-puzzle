package state

import (
	"fmt"
	"strings"
)

type State struct {
	Values    []int
	Width     int
	parent    *State
	Position  int
	MovesLeft int
	Step      int
}

type ByScore []State

func (b ByScore) Len() int      { return len(b) }
func (b ByScore) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByScore) Less(i, j int) bool {
	return b[i].MovesLeft+b[i].ChildrenCount() < b[j].MovesLeft+b[j].ChildrenCount()
	// return b[i].MovesLeft < b[j].MovesLeft
}

func CreateInitialState(size, startPos int) *State {
	v := make([]int, size*size)
	return &State{Values: v, Width: size, MovesLeft: len(v) - 1, Position: startPos}
}

func (g State) ChildrenCount() int {
	r := 0
	if p, _ := g.newSpot(0, -3); p != -1 && g.Values[p] == 0 {
		r++
	}
	if p, _ := g.newSpot(2, -2); p != -1 && g.Values[p] == 0 {
		r++
	}
	if p, _ := g.newSpot(3, 0); p != -1 && g.Values[p] == 0 {
		r++
	}
	if p, _ := g.newSpot(2, 2); p != -1 && g.Values[p] == 0 {
		r++
	}
	if p, _ := g.newSpot(0, 3); p != -1 && g.Values[p] == 0 {
		r++
	}
	if p, _ := g.newSpot(-2, 2); p != -1 && g.Values[p] == 0 {
		r++
	}
	if p, _ := g.newSpot(-3, 0); p != -1 && g.Values[p] == 0 {
		r++
	}
	if p, _ := g.newSpot(-2, -2); p != -1 && g.Values[p] == 0 {
		r++
	}
	return r
}

func (g State) GetMoves() []State {
	var r []State
	if p, dp := g.newSpot(0, -3); p != -1 && g.Values[p] == 0 {
		// fmt.Println("N")
		r = append(r, makeMove(g, dp))
	}
	if p, dp := g.newSpot(2, -2); p != -1 && g.Values[p] == 0 {
		// fmt.Println("NE")
		r = append(r, makeMove(g, dp))
	}
	if p, dp := g.newSpot(3, 0); p != -1 && g.Values[p] == 0 {
		// fmt.Println("E")
		r = append(r, makeMove(g, dp))
	}
	if p, dp := g.newSpot(2, 2); p != -1 && g.Values[p] == 0 {
		// fmt.Println("SE")
		r = append(r, makeMove(g, dp))
	}
	if p, dp := g.newSpot(0, 3); p != -1 && g.Values[p] == 0 {
		// fmt.Println("S")
		r = append(r, makeMove(g, dp))
	}
	if p, dp := g.newSpot(-2, 2); p != -1 && g.Values[p] == 0 {
		// fmt.Println("SW")
		r = append(r, makeMove(g, dp))
	}
	if p, dp := g.newSpot(-3, 0); p != -1 && g.Values[p] == 0 {
		// fmt.Println("W")
		r = append(r, makeMove(g, dp))
	}
	if p, dp := g.newSpot(-2, -2); p != -1 && g.Values[p] == 0 {
		// fmt.Println("NW")
		r = append(r, makeMove(g, dp))
	}
	// fmt.Println(len(r))
	return r
}

func Valid(val, limit int) bool {
	return val >= 0 && val < limit
}

func makeMove(from State, move int) State {
	val := make([]int, len(from.Values))
	copy(val, from.Values)
	val[from.Position] = move
	return State{
		Values:    val,
		MovesLeft: from.MovesLeft - 1,
		Width:     from.Width,
		Position:  from.Position + move,
		Step:      from.Step + 1,
		parent:    &from,
	}
}

func (g State) IsComplete() bool {
	return g.MovesLeft == 0
	// for i := 0; i < len(g.Values); i++ {
	// 	if g.Values[i] == 0 && i != g.Position {
	// 		return false
	// 	}
	// }
	// return true
}

func (g State) WinPath() string {
	steps := make([]int, len(g.Values))
	steps[g.Position] = g.Step
	p := g.parent
	for p != nil {
		steps[p.Position] = p.Step
		p = p.parent
	}
	var b strings.Builder
	for k, v := range steps {
		if k != 0 && k%g.Width == 0 {
			b.WriteString("\n")
		}
		if k != g.Position {
			fmt.Fprintf(&b, "[%2v] ", v+1)
		} else {
			fmt.Fprintf(&b, "[%2v] ", len(steps))
		}

	}
	return b.String()
}

func (g State) String() string {
	var b strings.Builder
	for k, v := range g.Values {
		if k != 0 && k%g.Width == 0 {
			b.WriteString("\n")
		}
		if k != g.Position {
			fmt.Fprintf(&b, "[%2v] ", v)
		} else {
			fmt.Fprintf(&b, "[%2v] ", ":)")
		}

	}
	return b.String()
}

func xy(position, width int) (int, int) {
	var x, y int
	x = position % width
	y = position / width
	return x, y
}

func pos(x, y, width int) int {
	return (y * width) + x
}

func (g State) newSpot(dX, dY int) (int, int) {
	x, y := xy(g.Position, g.Width)
	y += dY
	x += dX
	if Valid(x, g.Width) && Valid(y, g.Width) {
		pos := pos(x, y, g.Width)
		return pos, pos - g.Position
	}
	return -1, 0
}
