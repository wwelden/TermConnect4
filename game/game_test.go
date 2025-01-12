package game

import (
	"testing"

	"github.com/wwelden/TermConnect4/piece"
)

func TestGame_EmptyBoard(t *testing.T) {
	g := NewGame(7, 6)
	g.EmptyBoard()

	// Check all positions are blank
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if !g.Board[i][j].IsEmpty() {
				t.Errorf("Expected blank piece at position (%d,%d), got %s", i, j, g.Board[i][j].GetChip())
			}
		}
	}
}

func TestGame_LastEmptyRow(t *testing.T) {
	g := NewGame(7, 6)
	g.Start()

	// Test partially filled column
	g.Board[0][0] = *piece.InitPiece("red")
	row := g.LastEmptyRow(0)
	if row != g.Height-1 {
		t.Errorf("Expected row 5 for partially filled column, got %d", row)
	}

	// Test full column
	for i := 0; i < g.Height; i++ {
		g.Board[i][0] = *piece.InitPiece("red")
	}
	row = g.LastEmptyRow(0)
	if row != -1 {
		t.Errorf("Expected row -1 for full column, got %d", row)
	}
}

func TestGame_NewGame(t *testing.T) {
	g := NewGame(7, 6)

	if g.Width != 7 {
		t.Errorf("Expected width 7, got %d", g.Width)
	}
	if g.Height != 6 {
		t.Errorf("Expected height 6, got %d", g.Height)
	}
	if g.HasWinner {
		t.Error("Expected HasWinner to be false")
	}
	if !g.Turn.IsRed() {
		t.Errorf("Expected first turn to be red, got %s", g.Turn.GetChip())
	}
}

func TestGame_FlipTurn(t *testing.T) {
	g := NewGame(7, 6)

	// Test red to yellow
	if !g.Turn.IsRed() {
		t.Errorf("Expected red turn, got %s", g.Turn.GetChip())
	}
	g.FlipTurn()
	if !g.Turn.IsYellow() {
		t.Errorf("Expected yellow turn, got %s", g.Turn.GetChip())
	}

	// Test yellow to red
	g.FlipTurn()
	if !g.Turn.IsRed() {
		t.Errorf("Expected red turn, got %s", g.Turn.GetChip())
	}
}

func TestCheck4Horizontal(t *testing.T) {
	g := NewGame(7, 6)
	g.Start()

	// Test red horizontal win
	g.Board[0][0] = *piece.InitPiece("red")
	g.Board[0][1] = *piece.InitPiece("red")
	g.Board[0][2] = *piece.InitPiece("red")
	g.Board[0][3] = *piece.InitPiece("red")

	g.Check4Horizontal(4, "red")
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

	g.Check4Horizontal(4, "yellow")
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

	g.Check4Vertical(4, "red")
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

	g.Check4Vertical(4, "yellow")
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

	g.Check4Diagonal(4, "red")
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

	g.Check4Diagonal(4, "yellow")
	if !g.HasWinner {
		t.Error("Expected diagonal yellow win")
	}
}

func TestCheck4AllWins(t *testing.T) {
	g := NewGame(7, 6)
	g.Start()

	g.Board[0][0] = *piece.InitPiece("red")
	g.Board[0][1] = *piece.InitPiece("red")
	g.Board[0][2] = *piece.InitPiece("red")
	g.Board[0][3] = *piece.InitPiece("red")

	g.Check4AllWins(4, "red")
	if !g.HasWinner {
		t.Error("Expected diagonal red win")
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
