package game

import (
	"fmt"
	"strings"

	"github.com/wwelden/TermConnect4/piece"
)

func Contains(block []piece.Piece, input string) bool {
	strBlock := ""
	for _, elem := range block {
		strBlock += elem.GetChip()
	}
	return strings.Contains(strBlock, input) //cheap solution but why reinvent the wheel
}

func (g *Game) Check4Horizontal() {
	for _, row := range g.Board {
		if Contains(row, "🔴🔴🔴🔴") {
			g.HasWinner = true
			fmt.Println("Red Won")
		}
		if Contains(row, "🟡🟡🟡🟡") {
			g.HasWinner = true
			fmt.Println("Yellow Won")
		}
	}
}

func (g *Game) Check4Vertical() {
	for col := 0; col < g.Width; col++ {
		for row := 0; row <= g.Height-4; row++ {
			if g.Board[row][col].GetChip() == "🔴" &&
				g.Board[row+1][col].GetChip() == "🔴" &&
				g.Board[row+2][col].GetChip() == "🔴" &&
				g.Board[row+3][col].GetChip() == "🔴" {
				g.HasWinner = true
				fmt.Println("Red Won")
				return
			}
			if g.Board[row][col].GetChip() == "🟡" &&
				g.Board[row+1][col].GetChip() == "🟡" &&
				g.Board[row+2][col].GetChip() == "🟡" &&
				g.Board[row+3][col].GetChip() == "🟡" {
				g.HasWinner = true
				fmt.Println("Yellow Won")
				return
			}
		}
	}
}

func (g *Game) Check4Diagonal() {
	for col := 0; col <= g.Width-4; col++ {
		for row := 0; row <= g.Height-4; row++ {
			if g.Board[row][col].GetChip() == "🔴" &&
				g.Board[row+1][col+1].GetChip() == "🔴" &&
				g.Board[row+2][col+2].GetChip() == "🔴" &&
				g.Board[row+3][col+3].GetChip() == "🔴" {
				g.HasWinner = true
				fmt.Println("Red Won")
				return
			}
			if g.Board[row][col].GetChip() == "🟡" &&
				g.Board[row+1][col+1].GetChip() == "🟡" &&
				g.Board[row+2][col+2].GetChip() == "🟡" &&
				g.Board[row+3][col+3].GetChip() == "🟡" {
				g.HasWinner = true
				fmt.Println("Yellow Won")
				return
			}
		}
	}
	for col := 3; col < g.Width; col++ {
		for row := 0; row <= g.Height-4; row++ {
			if g.Board[row][col].GetChip() == "🔴" &&
				g.Board[row+1][col-1].GetChip() == "🔴" &&
				g.Board[row+2][col-2].GetChip() == "🔴" &&
				g.Board[row+3][col-3].GetChip() == "🔴" {
				g.HasWinner = true
				fmt.Println("Red Won")
				return
			}
			if g.Board[row][col].GetChip() == "🟡" &&
				g.Board[row+1][col-1].GetChip() == "🟡" &&
				g.Board[row+2][col-2].GetChip() == "🟡" &&
				g.Board[row+3][col-3].GetChip() == "🟡" {
				g.HasWinner = true
				fmt.Println("Yellow Won")
				return
			}
		}
	}
}

func (g *Game) Check4Wins() {
	g.Check4Horizontal()
	g.Check4Vertical()
	g.Check4Diagonal()
}
