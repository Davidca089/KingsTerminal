package main

import (
	"fmt"
	"log"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	// "github.com/Davidca089/chess_cli/internal/board"
)

func main() {
	p := tea.NewProgram(initChessModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal("Unexpected error")
	}
}

type ChessBoardModel struct {
	board         []int
	curPieceIndex int
	whiteView     bool
	width         int
	height        int
}

func initChessModel() ChessBoardModel {
	return ChessBoardModel{
		board:         make([]int, 64),
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
    accum += padStrNewLine(m.width/2 + 5, "----------")
    accum += padStrNewLine(m.width/2 + 5, "|THBQKBHT|")
    accum += padStrNewLine(m.width/2 + 5, "|PPPPPPPP|")
    accum += padStrNewLine(m.width/2 + 5, "|        |")
    accum += padStrNewLine(m.width/2 + 5, "|        |")
    accum += padStrNewLine(m.width/2 + 5, "|        |")
    accum += padStrNewLine(m.width/2 + 5, "|        |")
    accum += padStrNewLine(m.width/2 + 5, "|PPPPPPPP|")
    accum += padStrNewLine(m.width/2 + 5, "|THBQKBHT|")
    accum += padStrNewLine(m.width/2 + 5, "----------")
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
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the curPieceIndex up
		case "up", "k":
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
