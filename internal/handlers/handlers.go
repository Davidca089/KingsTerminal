package handlers

import (
	// "fmt"
	. "github.com/Davidca089/chess_cli/internal/structs"
)

func PeonMoves(p Piece, x, y int) []Position {
	ret := make([]Position, 0)
	if p.Color == White {
		ret = append(ret, Position{X: x, Y: y - 1})
		ret = append(ret, Position{X: x, Y: y - 2})
	} else {
		ret = append(ret, Position{X: x, Y: y + 1})
		ret = append(ret, Position{X: x, Y: y + 2})
	}
	// fmt.Println(ret)
	return ret
}
func KingMoves(p Piece, x, y int) []Position {
	ret := make([]Position, 8)
	return ret
}
func QueenMoves(p Piece, x, y int) []Position {
	ret := make([]Position, 8)
	return ret
}
func TowerMoves(p Piece, x, y int) []Position {
	ret := make([]Position, 8)
	return ret
}
func BishopMoves(p Piece, x, y int) []Position {
	ret := make([]Position, 8)
	return ret
}
func KnightMoves(p Piece, x, y int) []Position {
	ret := make([]Position, 8)
	return ret
}
