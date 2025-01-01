package main

import (
	"github.com/wwelden/TermConnect4/game"
)

func main() {
	game := game.NewGame(8, 8)
	game.Start()
	game.GameLoop()
}
