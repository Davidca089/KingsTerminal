package handlers

import (
	"strconv"

	. "github.com/Davidca089/chess_cli/internal/structs"
)

// ret true, list if there are still elements to be processed
func addPos(ret []Previous, board *[8][8]Piece, x, y, val int, col Color) (bool, []Previous) {
	if y < 0 || y > 7 || x < 0 || x > 7 || board[y][x].Color == col {
		return false, ret
	}
    prev := board[y][x]
    keys := []string{"'", ",",".", "p", "y", "f", "g", "c","r","l"}
    if val >= 10 {
        // if its more than this? WOMP WOMP
        board[y][x].PieceType = PieceType(keys[val - 10])
    } else {
        board[y][x].PieceType = PieceType(strconv.Itoa(val))
    }
    return true, append(ret, Previous{Position: Position{X: x, Y: y}, PrevPiece: prev})
}

func PeonMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	i := 1
	if p.Color == White {
		val, list := addPos(ret, board, x, y-i, i, p.Color)
		ret = list
		if !val {
			return ret
		}
		if y == 6 {
			i++
			val, list = addPos(ret, board, x, y-i, i, p.Color)
			ret = list
		}
	} else {
		val, list := addPos(ret, board, x, y+i, i, p.Color)
		ret = list
		// no more to process
		if !val {
			return ret
		}

		if y == 1 {
			i++
			val, list = addPos(ret, board, x, y+i, i, p.Color)
			ret = list
		}

	}
	// fmt.Println(ret)
	return ret
}

func KingMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	return ret
}
func QueenMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	// for i:= rang{
	//
	// }
	return ret
}
func TowerMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	i := 1
	for k := y - 1; k >= 0; k-- {
		val, list := addPos(ret, board, x, k, i, p.Color)
		ret = list
		if !val {
			break
		}
		i++
	}
	for k := y + 1; k < 8; k++ {
		val, list := addPos(ret, board, x, k, i, p.Color)
		ret = list
		if !val {
			break
		}
		i++
	}
	for k := x + 1; k < 8; k++ {
		val, list := addPos(ret, board, k, y, i, p.Color)
		ret = list
		if !val {
			break
		}
		i++
	}
	for k := x - 1; k >= 0; k-- {
		val, list := addPos(ret, board, k, y, i, p.Color)
		ret = list
		if !val {
			break
		}
		i++
	}
	return ret
}
func BishopMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	return ret
}
func KnightMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	return ret
}
