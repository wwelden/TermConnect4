package ai

import (
	"github.com/wwelden/TermConnect4/game"
)

//find the best move

// find prob of winning game

//IDK still planning this, probably just going with the same thing as the functional project

// eventually just use min max

func LastEmptyRow2(g game.Game, dex int) (int, int) {
	if dex < 0 || dex >= g.Width {
		return -1, -1
	}

	for y := g.Height - 1; y >= 0; y-- {
		if g.Board[y][dex].IsEmpty() {
			return y, dex
		}
	}
	return -1, -1
}

func findGoodSpot(g game.Game) {
	max := 0
	row := 0
	for i := 0; i < g.Width-1; i++ {
		val, r := LastEmptyRow2(g, i)
		if val > max {
			max = val
			row = r
		}
	}
	g.MakeMove(row)
}

func findKnRow(k int, g game.Game) {
	match := ""
	for i := 0; i < k; i++ {
		match += "🔴"
	}
	for _, row := range g.Board {
		if game.Contains(row, match) {

		}
	}
}
