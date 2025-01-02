package piece

import "testing"

func TestPiece_IsEmpty(t *testing.T) {
	p := InitPiece("blank")
	if !p.IsEmpty() {
		t.Errorf("Expected blank piece to be empty, got %s", p.GetChip())
	}
}

func TestPiece_IsRed(t *testing.T) {
	p := InitPiece("red")
	if !p.IsRed() {
		t.Errorf("Expected red piece to be red, got %s", p.GetChip())
	}
}

func TestPiece_IsYellow(t *testing.T) {
	p := InitPiece("yellow")
	if !p.IsYellow() {
		t.Errorf("Expected yellow piece to be yellow, got %s", p.GetChip())
	}
}

func TestPiece_Display(t *testing.T) {
	p := InitPiece("red")
	if p.Display() != "🔴" {
		t.Errorf("Expected red piece to display as 🔴, got %s", p.Display())
	}
}

func TestPiece_SetRed(t *testing.T) {
	p := InitPiece("blank")
	p.SetRed()
	if p.Display() != "🔴" {
		t.Errorf("Expected blank piece to be red, got %s", p.Display())
	}
}

func TestPiece_SetYellow(t *testing.T) {
	p := InitPiece("blank")
	p.SetYellow()
	if p.Display() != "🟡" {
		t.Errorf("Expected blank piece to be yellow, got %s", p.Display())
	}
}

func TestPiece_SetTurn(t *testing.T) {
	p := InitPiece("blank")
	red := InitPiece("red")
	p.SetTurn(*red)
	if p.Display() != "🔴" {
		t.Errorf("Expected blank piece to be red, got %s", p.Display())
	}
}

func TestPiece_InitPiece(t *testing.T) {
	p := InitPiece("red")
	if p.Display() != "🔴" {
		t.Errorf("Expected red piece to display as 🔴, got %s", p.Display())
	}
}

func TestPiece_GetChip(t *testing.T) {
	p := InitPiece("red")
	if p.GetChip() != "🔴" {
		t.Errorf("Expected red piece to get chip as 🔴, got %s", p.GetChip())
	}
}
