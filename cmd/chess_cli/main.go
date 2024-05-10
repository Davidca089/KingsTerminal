package main

import (
    "os"
    "fmt"
	"github.com/Davidca089/chess_cli/internal/board"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
	p := tea.NewProgram(board.InitChessModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal("Unexpected error")
	}
}
