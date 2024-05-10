package board

import (
	"github.com/Davidca089/chess_cli/internal/handlers"
	. "github.com/Davidca089/chess_cli/internal/structs"
	"github.com/Davidca089/chess_cli/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type ChessBoardModel struct {
	board              [8][8]Piece
	curX               int
	curY               int
	whiteView          bool
	pieceSelected      bool
	width              int
	height             int
}

func InitChessModel() ChessBoardModel {
	return ChessBoardModel{
		board:              StartBoard(),
		curX:               0,
		curY:               0,
		whiteView:          true,
		pieceSelected:      false,
		width:              0,
		height:             0,
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
    var Reset = "\033[0m"
    var Red = "\033[31m"
	accum := ""
	// skip up
	for range m.height/2 - 4 {
		accum += "\n"
	}
	accum += utils.PadStrNewLine(m.width/2+5, "----------")
	// print board
    var selectedPad = 0
	for y := 0; y < 8; y++ {
		inner := "|"
		for x := 0; x < 8; x++ {
			if m.pieceSelected && x == m.curX && y == m.curY {
                // fmt.Println("MEA")
				inner += Red + (string(m.board[y][x].PieceType)) + Reset
                selectedPad = 9
			} else {
				inner += string(m.board[y][x].PieceType)
			}
		}
		inner += "|"
        // fmt.Println(inner)
		accum += utils.PadStrNewLine(m.width/2+5 + selectedPad, inner)
        selectedPad = 0
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
		case "a":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[0].X, whitePieces[0].Y
			m.possibleMoves(pieceX, pieceY)
		case "o":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[1].X, whitePieces[1].Y
			m.possibleMoves(pieceX, pieceY)
		case "e":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[2].X, whitePieces[2].Y
			m.possibleMoves(pieceX, pieceY)
		case "u":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[3].X, whitePieces[3].Y
			m.possibleMoves(pieceX, pieceY)
		case "h":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[4].X, whitePieces[4].Y
			m.possibleMoves(pieceX, pieceY)
		case "t":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[5].X, whitePieces[5].Y
			m.possibleMoves(pieceX, pieceY)
		case "n":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[6].X, whitePieces[6].Y
			m.possibleMoves(pieceX, pieceY)
		case "s":
			m.pieceSelected = true
			pieceX, pieceY := whitePieces[7].X, whitePieces[7].Y
			m.possibleMoves(pieceX, pieceY)

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
