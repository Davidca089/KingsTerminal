package board

import (
	// "fmt"
	"github.com/Davidca089/chess_cli/internal/handlers"
	. "github.com/Davidca089/chess_cli/internal/structs"
	"github.com/Davidca089/chess_cli/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type ChessBoardModel struct {
	board         [8][8]Piece
	curX          int
	curY          int
	whiteView     bool
	pieceSelected bool
	width         int
	height        int
}

func InitChessModel() ChessBoardModel {
	return ChessBoardModel{
		board:         StartBoard(),
		curX:          0,
		curY:          0,
		whiteView:     true,
		pieceSelected: false,
		width:         0,
		height:        0,
	}
}

func (m ChessBoardModel) Init() tea.Cmd {
	return nil
}

func (m *ChessBoardModel) movePiece(oldX, oldY, newX, newY int) {
	m.board[newY][newX] = m.board[oldY][oldX]
	m.board[oldY][oldX] = EmptyPiece()
}

func (m *ChessBoardModel) possibleMoves(x, y int) []Position {
	// King
	// Queen
	// Tower
	// Bishop
	// Knigth
	// Peon
	// Empty
	m.curX = x
	m.curY = y
	piece := m.board[y][x]
	// color := piece.Color
	switch piece.PieceType {
	case King:
		return handlers.KingMoves(piece, x, y)
	case Queen:
		return handlers.QueenMoves(piece, x, y)
	case Tower:
		return handlers.TowerMoves(piece, x, y)
	case Bishop:
		return handlers.BishopMoves(piece, x, y)
	case Knigth:
		return handlers.KnightMoves(piece, x, y)
	case Peon:
		return handlers.PeonMoves(piece, x, y)
	}
	return nil
}

func (m ChessBoardModel) View() string {
	accum := ""
	// skip up
	for range m.height/2 - 4 {
		accum += "\n"
	}
	accum += utils.PadStrNewLine(m.width/2+5, "----------")
	// print board
	for i := 0; i < 8; i++ {
		inner := "|"
		for j := 0; j < 8; j++ {
			inner += string(m.board[i][j].PieceType)
		}
		inner += "|"
		accum += utils.PadStrNewLine(m.width/2+5, inner)
	}
	accum += utils.PadStrNewLine(m.width/2+5, "----------")
	return accum
}

func (m ChessBoardModel) collectPieces() ([16]Position, [16]Position) {
	var blackPieces [16]Position
	var whitePieces [16]Position

	w := 0
	b := 0
	for y := range 8 {
		for x := range 8 {
			if m.board[y][x].Color == White {
				whitePieces[w] = Position{X: x, Y: y}
				w++
			} else if m.board[y][x].Color == Black {
				blackPieces[b] = Position{X: x, Y: y}
				b++
			}
		}
	}

	// Reverse so getting pieces is intuitive from black's side
	for i, j := 0, len(blackPieces)-1; i < j; i, j = i+1, j-1 {
		blackPieces[i], blackPieces[j] = blackPieces[j], blackPieces[i]
	}

	return whitePieces, blackPieces
}

func (m ChessBoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// grab the size of the screen to render it
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		whitePieces, _ := m.collectPieces()
		switch msg.String() {
		case "ctrl+c", "q", "ctrl+d":
			return m, tea.Quit
		case "esc":
            m.pieceSelected = false
			return m, tea.Quit

		// The "up" and "k" keys move the curPieceIndex up
		case "up", "a":
			pieceX, pieceY := whitePieces[0].X, whitePieces[0].Y
			p := m.possibleMoves(pieceX, pieceY)
			x, y := p[0].X, p[0].Y
			// fmt.Println(x)
			// fmt.Println(y)
			// fmt.Println(m.curY)
			// fmt.Println(m.curX)
			m.movePiece(m.curX, m.curY, x, y)

		// The "down" and "j" keys move the curPieceIndex down
		case "down", "j":
			// if m.curPieceIndex < 8 {
			// }

		}
	}
	// Return the updated ChessBoardModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
