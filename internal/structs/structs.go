package structs

type PieceType string
type DisplayInfo string
type Color int

const (
	White Color = 0
	Black       = 1
	None        = 2
)

const (
	Reset    DisplayInfo = "\033[0m"
	Red                  = "\033[31m"
	Green                = "\033[32m"
	Yellow               = "\033[33m"
	Blue                 = "\033[34m"
	Magenta              = "\033[35m"
	Cyan                 = "\033[36m"
	Gray                 = "\033[37m"
	WhiteCol             = "\033[97m"
)

const (
	King   PieceType = "K"
	Queen            = "Q"
	Tower            = "T"
	Bishop           = "B"
	Knigth           = "H"
	Peon             = "P"
	Empty            = "."
)

type Position struct {
	X int
	Y int
}

type Piece struct {
	Color       Color
	PieceType   PieceType
	DisplayInfo DisplayInfo
}

func (p *Piece) SetDisplayInfo(dinfo DisplayInfo) {
	p.DisplayInfo = dinfo
}

func EmptyPiece() Piece {
	return Piece{
		Color:       None,
		PieceType:   Empty,
		DisplayInfo: WhiteCol,
	}
}

func StartBoard() [8][8]Piece {
	return [8][8]Piece{
		{
			{Color: Black, PieceType: Tower, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Knigth, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Bishop, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Queen, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: King, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Bishop, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Knigth, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Tower, DisplayInfo: WhiteCol},
		},
		{
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: Black, PieceType: Peon, DisplayInfo: WhiteCol},
		},
		{
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
		},
		{
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
		},
		{
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
		},
		{
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
			{Color: None, PieceType: Empty, DisplayInfo: WhiteCol},
		},
		{
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Peon, DisplayInfo: WhiteCol},
		},
		{
			{Color: White, PieceType: Tower, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Knigth, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Bishop, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Queen, DisplayInfo: WhiteCol},
			{Color: White, PieceType: King, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Bishop, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Knigth, DisplayInfo: WhiteCol},
			{Color: White, PieceType: Tower, DisplayInfo: WhiteCol},
		},
	}
}
