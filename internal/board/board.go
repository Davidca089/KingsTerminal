package board

import (
	"github.com/Davidca089/chess_cli/internal/handlers"
	. "github.com/Davidca089/chess_cli/internal/structs"
	"github.com/Davidca089/chess_cli/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type ChessBoardModel struct {
	board         [8][8]Piece
	prevPosition  []Position
	curX          int
	curY          int
	whiteView     bool
	pieceSelected bool
	modePad       int
	width         int
	height        int
}

func InitChessModel() ChessBoardModel {
	return ChessBoardModel{
		board:         StartBoard(),
		prevPosition:  make([]Position, 0),
		curX:          0,
		curY:          0,
		modePad:       0,
		whiteView:     true,
		pieceSelected: false,
		width:         0,
		height:        0,
	}
}

func (m ChessBoardModel) Init() tea.Cmd {
	return nil
}

func (m *ChessBoardModel) possibleMoves(x, y int) []Position {
	m.curX = x
	m.curY = y
	piece := m.board[y][x]
	switch piece.PieceType {
	case King:
		return handlers.KingMoves(&m.board, piece, x, y)
	case Queen:
		return handlers.QueenMoves(&m.board, piece, x, y)
	case Tower:
		return handlers.TowerMoves(&m.board, piece, x, y)
	case Bishop:
		return handlers.BishopMoves(&m.board, piece, x, y)
	case Knigth:
		return handlers.KnightMoves(&m.board, piece, x, y)
	case Peon:
		return handlers.PeonMoves(&m.board, piece, x, y)
	}
	return nil
}

// How i selected all these values is magical
func (m ChessBoardModel) View() string {
	accum := ""
	// skip up
	for range m.height/2 - 4 {
		accum += "\n"
	}
	accum += utils.PadStrNewLine(m.width/2+5, "----------")
	// print board
	for y := 0; y < 8; y++ {
		inner := "|"
		for x := 0; x < 8; x++ {
			piece := m.board[y][x]
			inner += string(piece.DisplayInfo) + string(piece.PieceType) + string(Reset)
		}
		inner += "|"
		accum += utils.PadStrNewLine(m.width/2+77, inner)
	}
	accum += utils.PadStrNewLine(m.width/2+5, "----------")
	return accum
}

func (m *ChessBoardModel) collectPiecesPosition() ([16]Position, [16]Position) {
	var blackPieces [16]Position
	var whitePiecesPos [16]Position

	w := 0
	b := 0
	for y := range 8 {
		for x := range 8 {
			if m.board[y][x].Color == White {
				whitePiecesPos[w] = Position{X: x, Y: y}
				w++
			} else if m.board[y][x].Color == Black {
				blackPieces[b] = Position{X: x, Y: y}
				b++
			}
		}
	}

	// Reverse so selecting pieces is intuitive from black's side
	for i, j := 0, len(blackPieces)-1; i < j; i, j = i+1, j-1 {
		blackPieces[i], blackPieces[j] = blackPieces[j], blackPieces[i]
	}

	return whitePiecesPos, blackPieces
}

func (m ChessBoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// grab the size of the screen to render it
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		whitePiecesPos, _ := m.collectPiecesPosition()
		keys := []rune{'a', 'o', 'e', 'u', 'h', 't', 'n', 's'}
		switch msg.String() {
		case "ctrl+c", "q", "ctrl+d":
			return m, tea.Quit

		case "esc":
			m.pieceSelected = false

		case ",":
			if len(whitePiecesPos) < 8 {
				m.modePad = 0
			} else {
				m.modePad = 8
			}
		case "'":
			m.modePad = 0

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if !m.pieceSelected {
				break
			}
			for _, pos := range m.prevPosition {
				x := pos.X
				y := pos.Y
				piece := m.board[m.curY][m.curX]
				if m.board[y][x].PieceType == PieceType(msg.String()) {
					m.board[y][x] = Piece{Color: White, PieceType: piece.PieceType, DisplayInfo: WhiteCol}
					m.board[m.curY][m.curX] = EmptyPiece()
                    m.curX = x
                    m.curY = y
				} else {
					m.board[y][x] = EmptyPiece()
				}
			}
            m.prevPosition = nil
			m.pieceSelected = false

		default:
			for i, r := range keys {
				if msg.String() == string(r) {
					m.pieceSelected = true
					// Update all colors before exploring new key
					m.board[m.curY][m.curX].SetDisplayInfo(WhiteCol)
					for _, pos := range m.prevPosition {
						m.board[pos.Y][pos.X].PieceType = "."
						m.board[pos.Y][pos.X].SetDisplayInfo(WhiteCol)
					}
					// Explore new possibilites
					pieceX, pieceY := whitePiecesPos[i+m.modePad].X, whitePiecesPos[i+m.modePad].Y
					m.board[pieceY][pieceX].SetDisplayInfo(Red)
					for _, pos := range m.possibleMoves(pieceX, pieceY) {
						m.board[pos.Y][pos.X].SetDisplayInfo(Yellow)
						m.prevPosition = append(m.prevPosition, pos)
					}

				}
			}

		}
	}
	// Return the updated ChessBoardModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
