package game

import (
	"bufio"
	"fmt"

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

	fmt.Print(" 1  2  3  4  5  6  7  8 \n")
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
	// num := rand.Intn(g.Width)
	// g.MakeMove(num)
	g.findGoodSpot()
	g.Display()
}

func (g *Game) GameLoop() {
	for !g.HasWinner {
		g.GetMove()
		g.Check4Wins()
		// g.BasicAI()
		g.AiMove()
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

func (g *Game) Check4Horizontal(count int, player string) bool {
	checkVal := ""
	if player == "red" {
		checkVal = "🔴"
	} else if player == "yellow" {
		checkVal = "🟡"
	}
	for _, row := range g.Board {
		redString := strings.Repeat(checkVal, count)
		if Contains(row, redString) {
			return true
		}
	}

	return false
}

func (g *Game) Check4Vertical(count int, player string) bool {
	checkVal := ""
	if player == "red" {
		checkVal = "🔴"
	} else if player == "yellow" {
		checkVal = "🟡"
	}
	for col := 0; col < g.Width; col++ {
		for row := 0; row <= g.Height-count; row++ {
			win := true
			for i := 0; i < count; i++ {
				if g.Board[row+i][col].GetChip() != checkVal {
					win = false
				}
			}
			if win {
				return true
			}
		}
	}
	return false
}

func (g *Game) Check4Diagonal(count int, player string) bool {
	checkVal := ""
	if player == "red" {
		checkVal = "🔴"
	} else if player == "yellow" {
		checkVal = "🟡"
	}
	// Check diagonal down-right
	for col := 0; col <= g.Width-count; col++ {
		for row := 0; row <= g.Height-count; row++ {
			win := true
			for i := 0; i < count; i++ {
				if g.Board[row+i][col+i].GetChip() != checkVal {
					win = false
					break
				}
			}
			if win {
				return true
			}
		}
	}

	// Check diagonal down-left
	for col := count - 1; col < g.Width; col++ {
		for row := 0; row <= g.Height-count; row++ {
			win := true
			for i := 0; i < count; i++ {
				if g.Board[row+i][col-i].GetChip() != checkVal {
					win = false
					break
				}
			}
			if win {
				return true
			}
		}
	}
	return false
}

func (g *Game) RedWin() {
	g.HasWinner = true
	g.ShowWinningMove()
	fmt.Println("Red Won")
}

func (g *Game) YellowWin() {
	g.HasWinner = true
	g.ShowWinningMove()
	fmt.Println("Yellow Won")
}
func (g *Game) Check4AllWins(num int, player string) bool {
	return g.Check4Horizontal(num, player) || g.Check4Vertical(num, player) || g.Check4Diagonal(num, player)
}

func (g *Game) Check4Wins() {
	if g.Check4AllWins(4, "red") {
		g.RedWin()
	} else if g.Check4AllWins(4, "yellow") {
		g.YellowWin()
	}
}

func (g *Game) LastEmptyRow2(dex int) (int, int) {
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

func (g *Game) findGoodSpot() {
	max := 0
	row := 0
	for i := 0; i < g.Width; i++ {
		val, r := g.LastEmptyRow2(i)
		if val > max {
			max = val
			row = r
		}
	}
	g.MakeMove(row)
}

func (g *Game) IsWinningMove(col int) bool {
	// Get the row where the piece would land
	row, col := g.LastEmptyRow2(col)
	if row == -1 {
		return false
	}

	// Temporarily make the move and check for red win
	originalPiece := g.Board[row][col]
	g.Board[row][col] = *piece.InitPiece("red")
	redWins := g.Check4AllWins(4, "red")

	// Check for yellow win
	g.Board[row][col] = *piece.InitPiece("yellow")
	yellowWins := g.Check4AllWins(4, "yellow")

	// Undo the move
	g.Board[row][col] = originalPiece

	return redWins || yellowWins
}

func (g *Game) AiMove() {
	for i := 0; i < g.Width; i++ {
		if g.IsWinningMove(i) {
			g.MakeMove(i)
			g.Display()
			return
		}
	}
	g.findGoodSpot()
	g.Display()
}

func (g *Game) ShowWinningMove() {
	// Get the winning sequence coordinates
	coords := g.findWinningSequence()
	if len(coords) == 4 {
		// Store original pieces
		originals := make([]piece.Piece, 4)
		for i, coord := range coords {
			originals[i] = g.Board[coord[0]][coord[1]]
			g.Board[coord[0]][coord[1]] = *piece.InitPiece("green")
		}

		// Display the board with green pieces
		g.Display()

		// Restore original pieces
		for i, coord := range coords {
			g.Board[coord[0]][coord[1]] = originals[i]
		}
	}
}


func (g *Game) FindMatchingSequence(jVal int, hOffSet int, wOffSet int, coords [][2]int) [][2]int {
	for i := 0; i <= g.Height-hOffSet-1; i++ {
		for j := jVal; j <= g.Width-wOffSet-1; j++ {
			if !g.Board[i][j].IsEmpty() {
				chip := g.Board[i][j].GetChip()
				nextI1 := i + coords[0][0]
				nextJ1 := j + coords[0][1]
				nextI2 := i + coords[1][0]
				nextJ2 := j + coords[1][1]
				nextI3 := i + coords[2][0]
				nextJ3 := j + coords[2][1]

				if g.Board[nextI1][nextJ1].GetChip() == chip &&
					g.Board[nextI2][nextJ2].GetChip() == chip &&
					g.Board[nextI3][nextJ3].GetChip() == chip {
					return [][2]int{{i, j}, {nextI1, nextJ1}, {nextI2, nextJ2}, {nextI3, nextJ3}}
				}
			}
		}
	}
	return [][2]int{}
}

func (g *Game) findWinningSequence() [][2]int {
	horizontalVals := [][2]int{{0, 1}, {0, 2}, {0, 3}}
	verticalVals := [][2]int{{1, 0}, {2, 0}, {3, 0}}
	rightDiagonalVals := [][2]int{{1, 1}, {2, 2}, {3, 3}}
	leftDiagonalVals := [][2]int{{1, -1}, {2, -2}, {3, -3}}

	if result := g.FindMatchingSequence(0, 0, 3, horizontalVals); len(result) > 0 {return result}
	if result := g.FindMatchingSequence(0, 3, 0, verticalVals); len(result) > 0 {return result}
	if result := g.FindMatchingSequence(0, 3, 3, rightDiagonalVals); len(result) > 0 {return result}
	if result := g.FindMatchingSequence(3, 3, 3, leftDiagonalVals); len(result) > 0 {return result}

	return [][2]int{}
}
