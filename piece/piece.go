package piece

type Piece struct {
	chip string
}

func (p *Piece) SetRed() {
	p.chip = "🔴"
}
func (p *Piece) SetYellow() {
	p.chip = "🟡"
}
func (p *Piece) Empty() {
	p.chip = "  "
}
func (p *Piece) IsEmpty() bool {
	return (p.chip == "  ")
}
func InitPiece(chp string) *Piece {
	switch chp {
	case "red":
		chp = "🔴"
	case "yellow":
		chp = "🟡"
	case "green":
		chp = "🟢"
	case "blank":
		chp = "  "
	}
	return &Piece{
		chip: chp,
	}
}

func (p *Piece) Display() string {
	return p.chip
}

func (p *Piece) SetTurn(turn Piece) {
	p.chip = turn.chip
}

func (p *Piece) IsRed() bool {
	return p.chip == "🔴"
}
func (p *Piece) IsYellow() bool {
	return p.chip == "🟡"
}
func (p *Piece) GetChip() string {
	return p.chip
}
