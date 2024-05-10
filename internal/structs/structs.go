package structs

type PieceType string
type Color int

const (
	White Color = 0
	Black       = 1
	None        = 2
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
	Color     Color
	PieceType PieceType
}

func EmptyPiece() Piece {
	return Piece{
		Color:     None,
		PieceType: Empty,
	}
}

func StartBoard() [8][8]Piece {
	return [8][8]Piece{
		{
            {Color: Black, PieceType: Tower},
            {Color: Black, PieceType: Knigth},
			{Color: Black, PieceType: Bishop},
			{Color: Black, PieceType: Queen},
			{Color: Black, PieceType: King},
			{Color: Black, PieceType: Bishop},
			{Color: Black, PieceType: Knigth},
			{Color: Black, PieceType: Tower},
		},
		{
			{Color: Black, PieceType: Peon},
			{Color: Black, PieceType: Peon},
			{Color: Black, PieceType: Peon},
			{Color: Black, PieceType: Peon},
			{Color: Black, PieceType: Peon},
			{Color: Black, PieceType: Peon},
			{Color: Black, PieceType: Peon},
			{Color: Black, PieceType: Peon},
		},
		{
            {Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
		},
		{
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
		},
		{
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
		},
		{
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
			{Color: None, PieceType: Empty},
		},
		{
            {Color: White, PieceType: Peon},
			{Color: White, PieceType: Peon},
			{Color: White, PieceType: Peon},
			{Color: White, PieceType: Peon},
			{Color: White, PieceType: Peon},
			{Color: White, PieceType: Peon},
			{Color: White, PieceType: Peon},
			{Color: White, PieceType: Peon},
		},
		{
			{Color: White, PieceType: Tower},
			{Color: White, PieceType: Knigth},
			{Color: White, PieceType: Bishop},
			{Color: White, PieceType: Queen},
			{Color: White, PieceType: King},
			{Color: White, PieceType: Bishop},
			{Color: White, PieceType: Knigth},
			{Color: White, PieceType: Tower},
		},
	}
}
