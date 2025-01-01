package game

import (
	"testing"

	"github.com/wwelden/TermConnect4/piece"
)

func TestCheck4Horizontal(t *testing.T) {
	g := NewGame(7, 6)
	g.Start()

	// Test red horizontal win
	g.Board[0][0] = *piece.InitPiece("red")
	g.Board[0][1] = *piece.InitPiece("red")
	g.Board[0][2] = *piece.InitPiece("red")
	g.Board[0][3] = *piece.InitPiece("red")

	g.Check4Horizontal()
	if !g.HasWinner {
		t.Error("Expected horizontal red win")
	}

	// Reset game
	g = NewGame(7, 6)
	g.Start()

	// Test yellow horizontal win
	g.Board[0][0] = *piece.InitPiece("yellow")
	g.Board[0][1] = *piece.InitPiece("yellow")
	g.Board[0][2] = *piece.InitPiece("yellow")
	g.Board[0][3] = *piece.InitPiece("yellow")

	g.Check4Horizontal()
	if !g.HasWinner {
		t.Error("Expected horizontal yellow win")
	}
}

func TestCheck4Vertical(t *testing.T) {
	g := NewGame(7, 6)
	g.Start()

	// Test red vertical win
	g.Board[0][0] = *piece.InitPiece("red")
	g.Board[1][0] = *piece.InitPiece("red")
	g.Board[2][0] = *piece.InitPiece("red")
	g.Board[3][0] = *piece.InitPiece("red")

	g.Check4Vertical()
	if !g.HasWinner {
		t.Error("Expected vertical red win")
	}

	// Reset game
	g = NewGame(7, 6)
	g.Start()

	// Test yellow vertical win
	g.Board[0][0] = *piece.InitPiece("yellow")
	g.Board[1][0] = *piece.InitPiece("yellow")
	g.Board[2][0] = *piece.InitPiece("yellow")
	g.Board[3][0] = *piece.InitPiece("yellow")

	g.Check4Vertical()
	if !g.HasWinner {
		t.Error("Expected vertical yellow win")
	}
}

func TestCheck4Diagonal(t *testing.T) {
	g := NewGame(7, 6)
	g.Start()

	// Test red diagonal win (bottom-left to top-right)
	g.Board[0][0] = *piece.InitPiece("red")
	g.Board[1][1] = *piece.InitPiece("red")
	g.Board[2][2] = *piece.InitPiece("red")
	g.Board[3][3] = *piece.InitPiece("red")

	g.Check4Diagonal()
	if !g.HasWinner {
		t.Error("Expected diagonal red win")
	}

	// Reset game
	g = NewGame(7, 6)
	g.Start()

	// Test yellow diagonal win (top-left to bottom-right)
	g.Board[3][0] = *piece.InitPiece("yellow")
	g.Board[2][1] = *piece.InitPiece("yellow")
	g.Board[1][2] = *piece.InitPiece("yellow")
	g.Board[0][3] = *piece.InitPiece("yellow")

	g.Check4Diagonal()
	if !g.HasWinner {
		t.Error("Expected diagonal yellow win")
	}
}

func TestCheck4Wins(t *testing.T) {
	g := NewGame(7, 6)
	g.Start()

	// Test that Check4Wins catches horizontal win
	g.Board[0][0] = *piece.InitPiece("red")
	g.Board[0][1] = *piece.InitPiece("red")
	g.Board[0][2] = *piece.InitPiece("red")
	g.Board[0][3] = *piece.InitPiece("red")

	g.Check4Wins()
	if !g.HasWinner {
		t.Error("Check4Wins should detect horizontal win")
	}

	// Reset and test vertical win
	g = NewGame(7, 6)
	g.Start()

	g.Board[0][0] = *piece.InitPiece("yellow")
	g.Board[1][0] = *piece.InitPiece("yellow")
	g.Board[2][0] = *piece.InitPiece("yellow")
	g.Board[3][0] = *piece.InitPiece("yellow")

	g.Check4Wins()
	if !g.HasWinner {
		t.Error("Check4Wins should detect vertical win")
	}
}
