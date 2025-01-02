package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/wwelden/TermConnect4/piece"
)

type Game struct {
	Board     [][]piece.Piece
	Turn      piece.Piece
	HasWinner bool
	Width     int
	Height    int
}

func (g *Game) Display() {
	fmt.Print("\033[H\033[2J")
	printOut := ""
	for _, row := range g.Board {
		for _, col := range row {
			printOut += "|" + col.Display()
		}
		printOut += "|\n"
	}
	bar := ""
	for i := 0; i < g.Width; i++ {
		bar += "---"
	}
	bar += "-\n"

	fmt.Print(printOut)
	fmt.Print(bar)
}

func (g *Game) EmptyBoard() {
	g.Board = make([][]piece.Piece, g.Height)
	for i := range g.Board {
		g.Board[i] = make([]piece.Piece, g.Width)
		for j := range g.Board[i] {
			g.Board[i][j] = *piece.InitPiece("blank")
		}
	}
}

func (g *Game) Start() {
	g.EmptyBoard()
	g.Display()
}

func NewGame(width, height int) *Game {
	return &Game{
		Width:     width,
		Height:    height,
		Board:     make([][]piece.Piece, height),
		HasWinner: false,
		Turn:      *piece.InitPiece("red"),
	}
}
func (g *Game) LastEmptyRow(dex int) int {
	if dex < 0 || dex >= g.Width {
		return -1
	}

	for y := g.Height - 1; y >= 0; y-- {
		if g.Board[y][dex].IsEmpty() {
			return y
		}
	}
	return -1
}
func (g *Game) FlipTurn() {
	if g.Turn.IsRed() {
		g.Turn.SetYellow()
	} else if g.Turn.IsYellow() {
		g.Turn.SetRed()
	}
}

func (g *Game) MakeMove(dex int) {
	row := g.LastEmptyRow(dex)
	if row == -1 {
		fmt.Print("Row is full, pick another \n")
		g.GetMove()
	} else {
		g.Board[row][dex].SetTurn(g.Turn)
		g.FlipTurn()
	}
}

func (g *Game) GetMove() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	num, _ := strconv.Atoi(text)
	if num < 0 || num > g.Width {
		fmt.Println("Please enter a valid number")
		return
	}
	g.MakeMove(num - 1)

	g.Display()
}

func (g *Game) BasicAI() {
	num := rand.Intn(g.Width)
	g.MakeMove(num)
	g.Display()
}

func (g *Game) GameLoop() {
	for !g.HasWinner {
		g.GetMove()
		g.Check4Wins()
		g.BasicAI()
		g.Check4Wins()
	}
}

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
