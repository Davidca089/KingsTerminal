package board

import (
	"github.com/Davidca089/chess_cli/internal/handlers"
	. "github.com/Davidca089/chess_cli/internal/structs"
	"github.com/Davidca089/chess_cli/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type ChessBoardModel struct {
	Board         [8][8]Piece
	prevPosition  []Previous
	whiteView     bool
	pieceSelected bool
	shiftPressed  bool
	curX          int
	curY          int
	modePad       int
	width         int
	height        int
}

func InitChessModel() ChessBoardModel {
	return ChessBoardModel{
		Board:         StartBoard(),
		prevPosition:  make([]Previous, 0),
		curX:          4,
		curY:          6,
		modePad:       0,
		whiteView:     true,
		shiftPressed:  false,
		pieceSelected: false,
		width:         0,
		height:        0,
	}
}

func (m ChessBoardModel) Init() tea.Cmd {
	return nil
}

func (m *ChessBoardModel) possibleMoves(x, y int) []Previous {
	m.curX = x
	m.curY = y
	piece := m.Board[y][x]
	switch piece.PieceType {
	case King:
		return handlers.KingMoves(&m.Board, piece, x, y)
	case Queen:
		return handlers.QueenMoves(&m.Board, piece, x, y)
	case Tower:
		return handlers.TowerMoves(&m.Board, piece, x, y)
	case Bishop:
		return handlers.BishopMoves(&m.Board, piece, x, y)
	case Knigth:
		return handlers.KnightMoves(&m.Board, piece, x, y)
	case Pawn:
		return handlers.PawnMoves(&m.Board, piece, x, y)
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
	// print Board
	for y := 0; y < 8; y++ {
		inner := "|"
		for x := 0; x < 8; x++ {
			piece := m.Board[y][x]
			inner += string(piece.DisplayInfo) + string(piece.PieceType) + string(Reset)
		}
		inner += "|"
		accum += utils.PadStrNewLine(m.width/2+77, inner)
	}
	accum += utils.PadStrNewLine(m.width/2+5, "----------")
	return accum
}

func (m ChessBoardModel) CollectPiecesPosition() ([16]Position, [16]Position) {
	var blackPiecesPos [16]Position
	var whitePiecesPos [16]Position
	// we will inspect from the top-down to determine the piece key placement
	var whiteStack [8][]Position
	var blackStack [8][]Position
	w := 0
	b := 0
	for x := range 8 {
		for y := range 8 {
			if m.Board[y][x].Color == White {
				whiteStack[w] = append(whiteStack[w], Position{X: x, Y: y})
			} else if m.Board[y][x].Color == Black {
				blackStack[b] = append(blackStack[b], Position{X: x, Y: y})
			}
		}
		w++
        b++
	}

	// unstack
	k := 0
	p := 0
	hasElements := true
	for hasElements {
		hasElements = false
		for _, x := range whiteStack {
			if len(x) == 0 {
				k++
				continue
			}
			hasElements = true
			whitePiecesPos[p] = x[0]
			whiteStack[k] = x[1:]
			p++
			k++
		}
		k = 0
	}

	// unstack
	k = 0
	p = 0
	hasElements = true
	for hasElements {
		hasElements = false
		for _, x := range blackStack {
			if len(x) == 0 {
				k++
				continue
			}
			hasElements = true
			blackPiecesPos[p] = x[0]
			blackStack[k] = x[1:]
			p++
			k++
		}
		k = 0
	}

	// Reverse so selecting pieces is intuitive from black's side
	for i, j := 0, len(blackPiecesPos)-1; i < j; i, j = i+1, j-1 {
		blackPiecesPos[i], blackPiecesPos[j] = blackPiecesPos[j], blackPiecesPos[i]
	}

	return whitePiecesPos, blackPiecesPos
}

func (m ChessBoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// grab the size of the screen to render it
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	case tea.KeyMsg:
		whitePiecesPos, _ := m.CollectPiecesPosition()
		for i, pos := range whitePiecesPos {
			if m.modePad == 0 {
				if i >= 8 {
					break
				}
				m.Board[pos.Y][pos.X].SetDisplayInfo(Yellow)
			} else {
				if i < 8 {
                    continue
				}
				m.Board[pos.Y][pos.X].SetDisplayInfo(Yellow)
			}
		}
		keys := []rune{'a', 'o', 'e', 'u', 'h', 't', 'n', 's'}
		switch msg.String() {

		case "ctrl+c", "ctrl+d":
			return m, tea.Quit

		case "esc":
			m.pieceSelected = false
			m.Board[m.curY][m.curX].SetDisplayInfo(Yellow)
			for _, pos := range m.prevPosition {
				m.Board[pos.Y][pos.X] = pos.PrevPiece
				m.Board[pos.Y][pos.X].SetDisplayInfo(WhiteCol)
			}

			m.prevPosition = nil

		case "shift+tab":
			// if key selected u clear everything out first
			if m.pieceSelected {
				m.Board[m.curY][m.curX].SetDisplayInfo(WhiteCol)
				for _, pos := range m.prevPosition {
					m.Board[pos.Y][pos.X] = pos.PrevPiece
				}
				m.pieceSelected = false
				m.prevPosition = nil
			}

			if !m.shiftPressed {
				m.modePad = 8
				if len(whitePiecesPos) < 8 {
					m.modePad = 0
				}
				for i, pos := range whitePiecesPos {
					if i < 8 {
						m.Board[pos.Y][pos.X].SetDisplayInfo(WhiteCol)
					} else {
						m.Board[pos.Y][pos.X].SetDisplayInfo(Yellow)
					}
				}
			} else {
				m.modePad = 0
				for i, pos := range whitePiecesPos {
					if i < 8 {
						m.Board[pos.Y][pos.X].SetDisplayInfo(Yellow)
					} else {
						m.Board[pos.Y][pos.X].SetDisplayInfo(WhiteCol)
					}
				}
			}

			m.shiftPressed = !m.shiftPressed

		case "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"'", ",", ".", "p", "y", "f", "g", "c", "r", "l":
			if !m.pieceSelected {
				break
			}
			for _, pos := range m.prevPosition {
				x := pos.X
				y := pos.Y
				piece := m.Board[m.curY][m.curX]
				if m.Board[y][x].PieceType == PieceType(msg.String()) {
					m.Board[y][x] = piece
					m.Board[m.curY][m.curX] = EmptyPiece()
					m.curX = x
					m.curY = y
					m.Board[m.curY][m.curX].SetDisplayInfo(WhiteCol)
				} else {
					m.Board[y][x] = pos.PrevPiece
				}
			}
			m.prevPosition = nil
			m.pieceSelected = false

		default:
			for i, r := range keys {
				if msg.String() == string(r) {
					m.pieceSelected = true
					// Restore all colors before exploring new key
					// m.Board[m.curY][m.curX].SetDisplayInfo(Yellow)
					for _, pos := range m.prevPosition {
						m.Board[pos.Y][pos.X] = pos.PrevPiece
					}

					m.prevPosition = nil
					// Explore new possibilites
					pieceX, pieceY := whitePiecesPos[i+m.modePad].X, whitePiecesPos[i+m.modePad].Y
					m.Board[pieceY][pieceX].SetDisplayInfo(Red)
					for _, pos := range m.possibleMoves(pieceX, pieceY) {
						m.Board[pos.Y][pos.X].SetDisplayInfo(Green)
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
