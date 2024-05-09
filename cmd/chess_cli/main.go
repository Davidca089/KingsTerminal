package main

import (
	"log"
    "github.com/Davidca089/chess_cli/internal/board"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(board.InitChessModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal("Unexpected error")
	}
}
