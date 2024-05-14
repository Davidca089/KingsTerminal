package handlers

import (
	. "github.com/Davidca089/chess_cli/internal/structs"
	"strconv"
)

// ret true, list if there are still elements to be processed
func addPos(ret []Previous, board *[8][8]Piece, x, y, val int, curPiece Piece) (bool, []Previous, bool) {
	stop := true
	col := curPiece.Color
	if y < 0 || y > 7 || x < 0 || x > 7 || board[y][x].Color == col {
		return false, ret, stop
	}

	scanning := board[y][x]
	keys := []string{"'", ",", ".", "p", "y", "f", "g", "c", "r", "l"}
	if val >= 10 {
		// if its more than this? WOMP WOMP
		board[y][x] = NewPiece(None, PieceType(keys[val-10]), WhiteCol)
	} else {
		board[y][x] = NewPiece(None, PieceType(strconv.Itoa(val)), WhiteCol)
	}

	if scanning.Color != col && scanning.Color != None && curPiece.PieceType != Peon {
		return false, append(ret, Previous{Position: Position{X: x, Y: y},
			PrevPiece: scanning}), !stop
	}

	return true, append(ret, Previous{Position: Position{X: x, Y: y}, PrevPiece: scanning}), !stop
}

func PeonMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	i := 1
	if p.Color == White {
		val, list, _ := addPos(ret, board, x, y-i, i, p)
		ret = list
		if !val {
			return ret
		}
		if y == 6 {
			i++
			val, list, _ = addPos(ret, board, x, y-i, i, p)
			ret = list
		}
	} else {
		val, list, _ := addPos(ret, board, x, y+i, i, p)
		ret = list
		// no more to process
		if !val {
			return ret
		}

		if y == 1 {
			i++
			val, list, _ = addPos(ret, board, x, y+i, i, p)
			ret = list
		}

	}
	// fmt.Println(ret)
	return ret
}

func KingMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	print_val := 1
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && y == j {
				continue
			}

			val, list, stop := addPos(ret, board, i, j, print_val, p)
			ret = list
			if !stop && !val {
				print_val++
			}
			if !val {
				continue
			}
			print_val++
		}
	}
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
		val, list, stop := addPos(ret, board, x, k, i, p)
		ret = list
		if !stop && !val {
			i++
		}
		if !val {
			break
		}
		i++
	}
	for k := y + 1; k < 8; k++ {
		val, list, stop := addPos(ret, board, x, k, i, p)
		ret = list
		if !stop && !val {
			i++
		}
		if !val {
			break
		}
		i++
	}
	for k := x + 1; k < 8; k++ {
		val, list, stop := addPos(ret, board, k, y, i, p)
		ret = list
		if !stop && !val {
			i++
		}
		if !val {
			break
		}
		i++
	}
	for k := x - 1; k >= 0; k-- {
		val, list, stop := addPos(ret, board, k, y, i, p)
		ret = list
		if !stop && !val {
			i++
		}
		if !val {
			break
		}
		i++
	}
	return ret
}
func BishopMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	print_val := 1
	for y_cross, x_cross := y-1, x-1; y_cross >= 0 || x_cross >= 0; {
		val, list, stop := addPos(ret, board, x_cross, y_cross, print_val, p)
		ret = list
		if !stop && !val {
			print_val++
		}
		if !val {
			break
		}
		print_val++
		y_cross--
		x_cross--
	}

	for y_cross, x_cross := y+1, x+1; y_cross < 8 || x_cross < 8; {
		val, list, stop := addPos(ret, board, x_cross, y_cross, print_val, p)
		ret = list
		if !stop && !val {
			print_val++
		}
		if !val {
			break
		}
		print_val++
		y_cross++
		x_cross++
	}

	for y_cross, x_cross := y+1, x-1; y_cross < 8 || x_cross < 8; {
		val, list, stop := addPos(ret, board, x_cross, y_cross, print_val, p)
		ret = list
		if !stop && !val {
			print_val++
		}
		if !val {
			break
		}
		print_val++
		y_cross++
		x_cross--
	}

	for y_cross, x_cross := y-1, x+1; y_cross < 8 || x_cross < 8; {
		val, list, stop := addPos(ret, board, x_cross, y_cross, print_val, p)
		ret = list
		if !stop && !val {
			print_val++
		}
		if !val {
			break
		}
		print_val++
		y_cross--
		x_cross++
	}

	return ret
}

func KnightMoves(board *[8][8]Piece, p Piece, x, y int) []Previous {
	ret := make([]Previous, 0)
	// make it circular so moving is intuitive
	coords := []Position{
		{Y: y - 1, X: x - 2},
		{Y: y - 2, X: x - 1},

		{Y: y - 2, X: x + 1},
		{Y: y - 1, X: x + 2},

		{Y: y + 1, X: x + 2},
		{Y: y + 2, X: x + 1},

		{Y: y + 2, X: x - 1},
		{Y: y + 1, X: x - 2},
	}

	print_val := 1
	for _, pos := range coords {
		xMove, yMove := pos.X, pos.Y
		val, list, _ := addPos(ret, board, xMove, yMove, print_val, p)
		ret = list
		if !val {
            continue
		}
		print_val++
	}
	return ret
}
