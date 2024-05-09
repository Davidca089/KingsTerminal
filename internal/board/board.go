package board

import (
	"fmt"
	"strconv"
	tea "github.com/charmbracelet/bubbletea"
)

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
	Empty            = " "
)

type Piece struct {
	Color     Color
	PieceType PieceType
}

type ChessBoardModel struct {
	board         [64]Piece
	curPieceIndex int
	whiteView     bool
	width         int
	height        int
}

func emptyPiece() Piece {
    return Piece{
        Color: None,
        PieceType: Empty,
    }
}

func startBoard() [64]Piece {
    return [64]Piece{
        {Black, Tower},
        {Black, Knigth},
        {Black, Bishop},
        {Black, Queen},
        {Black, King},
        {Black, Bishop},
        {Black, Knigth},
        {Black, Tower},
        {Black, Peon},
        {Black, Peon},
        {Black, Peon},
        {Black, Peon},
        {Black, Peon},
        {Black, Peon},
        {Black, Peon},
        {Black, Peon},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {None, Empty},
        {White, Peon},
        {White, Peon},
        {White, Peon},
        {White, Peon},
        {White, Peon},
        {White, Peon},
        {White, Peon},
        {White, Peon},
        {White, Tower},
        {White, Knigth},
        {White, Bishop},
        {White, Queen},
        {White, King},
        {White, Bishop},
        {White, Knigth},
        {White, Tower},
    }
}

func InitChessModel() ChessBoardModel {
	return ChessBoardModel{
		board:         startBoard(),
		curPieceIndex: 0,
		whiteView:     true,
		width:         0,
		height:        0,
	}
}


func (m ChessBoardModel) Init() tea.Cmd {
	return nil
}

func padStrNewLine(pad int, str string) string {
	middle := "%" + strconv.Itoa(pad) + "s"
	return fmt.Sprintf(middle, str) + "\n"
}

func padStr(pad int, str string) string {
	middle := "%" + strconv.Itoa(pad) + "s"
	return fmt.Sprintf(middle, str)
}

func (m ChessBoardModel) View() string {
	accum := ""
	// skip up
	for range m.height/2 - 4 {
		accum += "\n"
	}
	accum += padStrNewLine(m.width/2+5, "----------")
    // print board
    for i := 0; i < 64; i += 8 {
        inner := "|"
        for j := 0; j < 8; j++ {
            inner += string(m.board[i+j].PieceType)
        }
        inner += "|"
        accum += padStrNewLine(m.width/2+5,inner)
    }
	accum += padStrNewLine(m.width/2+5, "----------")
	return accum
}

func (m ChessBoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// grab the size of the screen to render it
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c", "q", "ctrl+d":
			return m, tea.Quit

		// The "up" and "k" keys move the curPieceIndex up
		case "up", "k":
            p := m.board[0]
            m.board[16] = p
            m.board[0] = emptyPiece()
			if m.curPieceIndex > 0 {
				m.curPieceIndex -= 8
			}

		// The "down" and "j" keys move the curPieceIndex down
		case "down", "j":
			m.curPieceIndex += 8
			// if m.curPieceIndex < 8 {
			// }

		}
	}
	// Return the updated ChessBoardModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
