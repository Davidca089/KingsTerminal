package handlers

import (
	"strconv"

	. "github.com/Davidca089/chess_cli/internal/structs"
)

// ret true, list if there are still elements to be processed
func addPos(ret []Position, board *[8][8]Piece, x, y, val int, col Color) (bool, []Position) {
	if y < 0 || y > 7 || x < 0 || x > 7 || board[y][x].Color == col {
		return false, ret
	}
	board[y][x].PieceType = PieceType(strconv.Itoa(val))
	return true, append(ret, Position{X: x, Y: y})
}

func PeonMoves(board *[8][8]Piece, p Piece, x, y int) []Position {
	ret := make([]Position, 0)
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

func KingMoves(board *[8][8]Piece, p Piece, x, y int) []Position {
	ret := make([]Position, 0)
	return ret
}
func QueenMoves(board *[8][8]Piece, p Piece, x, y int) []Position {
	ret := make([]Position, 0)
	// for i:= rang{
	//
	// }
	return ret
}
func TowerMoves(board *[8][8]Piece, p Piece, x, y int) []Position {
	ret := make([]Position, 0)
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
func BishopMoves(board *[8][8]Piece, p Piece, x, y int) []Position {
	ret := make([]Position, 0)
	return ret
}
func KnightMoves(board *[8][8]Piece, p Piece, x, y int) []Position {
	ret := make([]Position, 0)
	return ret
}
