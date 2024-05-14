package board_test

import (
	"fmt"
	"testing"

	"github.com/Davidca089/chess_cli/internal/board"
	"github.com/Davidca089/chess_cli/internal/structs"
)

func TestPieceCollectingThreeStacked(t *testing.T) {
    Board := board.InitChessModel()
    Board.Board[7][1] = structs.EmptyPiece()
    Board.Board[5][2] = structs.NewPiece(structs.White, structs.Knigth, structs.WhiteCol)
    fmt.Println(Board.View())

    white, _ := Board.CollectPiecesPosition();
    fmt.Println(white)
}

func TestPieceCollecting(t *testing.T) {
    Board := board.InitChessModel()
    white, _ := Board.CollectPiecesPosition();

    fmt.Println(white)
}
