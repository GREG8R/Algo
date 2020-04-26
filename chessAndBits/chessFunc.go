package main

import (
	"strconv"
	"strings"
)

func King(strPosition string) []uint64 {
	result := make([]uint64, 2)
	position, _ := strconv.Atoi(strPosition)

	k := uint64(1) << position
	nA := uint64(0xFEFEFEFEFEFEFEFE)
	nH := uint64(0x7F7F7F7F7F7F7F7F)
	p := make([]uint64, 8)
	p[0] = (k & nA) >> 9
	p[1] = k >> 8
	p[2] = (k & nH) >> 7
	p[3] = (k & nA) >> 1
	p[4] = (k & nH) << 1
	p[5] = (k & nA) << 7
	p[6] = k << 8
	p[7] = (k & nH) << 9
	result[1] = p[0] | p[1] | p[2] |
		p[3] | p[4] |
		p[5] | p[6] | p[7]

	for _, v := range p {
		if v > 0 {
			result[0]++
		}
	}

	return result
}

func Knight(strPosition string) []uint64 {
	result := make([]uint64, 2)
	position, _ := strconv.Atoi(strPosition)

	k := uint64(1) << position
	nA := uint64(0xFEFEFEFEFEFEFEFE)
	nAB := uint64(0xFCFCFCFCFCFCFCFC)
	nGH := uint64(0x3F3F3F3F3F3F3F3F)
	nH := uint64(0x7F7F7F7F7F7F7F7F)
	p := make([]uint64, 10)
	p[0] = (k & nA) << 15
	p[1] = (k & nH) << 17
	p[2] = (k & nGH) << 10
	p[3] = (k & nAB) << 6
	p[4] = (k & nGH) >> 6
	p[5] = (k & nH) >> 15
	p[6] = (k & nA) >> 17
	p[7] = (k & nAB) >> 10

	result[1] = p[0] | p[1] | p[2] |
		p[3] | p[4] |
		p[5] | p[6] | p[7]

	for _, v := range p {
		if v > 0 {
			result[0]++
		}
	}

	return result
}

// start position in fen notation rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
const (
	whitePawns   = 'P'
	whiteKnights = 'N'
	whiteBishops = 'B'
	whiteRooks   = 'R'
	whiteQueens  = 'Q'
	whiteKing    = 'K'

	blackPawns   = 'p'
	blackKnights = 'n'
	blackBishops = 'b'
	blackRooks   = 'r'
	blackQueens  = 'q'
	blackKing    = 'k'
)

var m = map[rune]int{
	whitePawns:   0,
	whiteKnights: 1,
	whiteBishops: 2,
	whiteRooks:   3,
	whiteQueens:  4,
	whiteKing:    5,
	blackPawns:   6,
	blackKnights: 7,
	blackBishops: 8,
	blackRooks:   9,
	blackQueens:  10,
	blackKing:    11,
}

func FEN(fen string) []uint64 {

	result := make([]uint64, 12)
	strArr := strings.Split(fen, "/")

	for i, str := range strArr {
		j := 0
		for _, c := range str {
			currentPosition := 63 - (i * 8) - 7 + j
			binaryPosition := uint64(1) << currentPosition

			switch c {
			case whitePawns, whiteKnights, whiteBishops,
				whiteRooks, whiteQueens, whiteKing,
				blackPawns, blackKnights, blackBishops,
				blackRooks, blackQueens, blackKing:
				result[m[c]] = result[m[c]] | binaryPosition
				j++
			default:
				j += int(c - '0')
			}
		}
	}
	return result
}

func LineItems(fen string) []uint64 {
	result := make([]uint64, 3)

	arrayOfPositions := FEN(fen)

	whiteQueenPosition := arrayOfPositions[m[whiteQueens]]
	whiteBishopPosition := arrayOfPositions[m[whiteBishops]]
	whiteRookPosition := arrayOfPositions[m[whiteRooks]]

	unionBlackMask := arrayOfPositions[m[blackRooks]] |
		arrayOfPositions[m[blackKnights]] |
		arrayOfPositions[m[blackBishops]] |
		arrayOfPositions[m[blackQueens]] |
		arrayOfPositions[m[blackKing]] |
		arrayOfPositions[m[blackPawns]]

	unionWhiteMaskWithoutRook := arrayOfPositions[m[whiteKnights]] |
		arrayOfPositions[m[whiteBishops]] |
		arrayOfPositions[m[whiteQueens]] |
		arrayOfPositions[m[whiteKing]] |
		arrayOfPositions[m[whitePawns]]

	unionWhiteMaskWithoutBishop := arrayOfPositions[m[whiteKnights]] |
		arrayOfPositions[m[whiteRooks]] |
		arrayOfPositions[m[whiteQueens]] |
		arrayOfPositions[m[whiteKing]] |
		arrayOfPositions[m[whitePawns]]

	unionWhiteMaskWithoutQueen := arrayOfPositions[m[whiteKnights]] |
		arrayOfPositions[m[whiteBishops]] |
		arrayOfPositions[m[whiteRooks]] |
		arrayOfPositions[m[whiteKing]] |
		arrayOfPositions[m[whitePawns]]

	result[0] = rookCalc(whiteRookPosition, unionBlackMask, unionWhiteMaskWithoutRook)
	result[1] = bishopCalc(whiteBishopPosition, unionBlackMask, unionWhiteMaskWithoutBishop)
	result[2] = rookCalc(whiteQueenPosition, unionBlackMask, unionWhiteMaskWithoutQueen) |
		bishopCalc(whiteQueenPosition, unionBlackMask, unionWhiteMaskWithoutQueen)

	return result
}

func rookCalc(whiteRookPosition, unionBlackMask, unionWhiteMaskWithoutRook uint64) uint64 {
	startPosOfRook := -1
	for i := whiteRookPosition; i > 0; i = i >> 1 {
		startPosOfRook++
	}

	resRook := uint64(0)
	posLeft := startPosOfRook - 1
	posRight := startPosOfRook + 1
	posUp := startPosOfRook + 8
	posDown := startPosOfRook - 8

	for {
		if posLeft >= 0 && posLeft % 8 != 7 &&
			(uint64(1) << (posLeft + 1)) | unionBlackMask != unionBlackMask &&
			((uint64(1) << posLeft) | unionWhiteMaskWithoutRook != unionWhiteMaskWithoutRook) {
			resRook = resRook | (uint64(1) << posLeft)
			posLeft--
		} else {
			break
		}
	}

	for {
		if  posRight <= 63 && posRight % 8 != 0 &&
			(uint64(1) << (posRight - 1)) | unionBlackMask != unionBlackMask &&
			((uint64(1) << posRight) | unionWhiteMaskWithoutRook != unionWhiteMaskWithoutRook) {
			resRook = resRook | (uint64(1) << posRight)
			posRight++
		} else {
			break
		}
	}

	for {
		if posUp <= 63 &&
			(uint64(1) << (posUp - 8)) | unionBlackMask != unionBlackMask &&
			((uint64(1) << posUp) | unionWhiteMaskWithoutRook != unionWhiteMaskWithoutRook) {
			resRook = resRook | (uint64(1) << posUp)
			posUp += 8
		} else {
			break
		}
	}

	for {
		if posDown >= 0 &&
			(uint64(1) << (posDown + 8)) | unionBlackMask != unionBlackMask &&
			(uint64(1) << posDown)|unionWhiteMaskWithoutRook != unionWhiteMaskWithoutRook {
			resRook = resRook | (uint64(1) << posDown)
			posDown -= 8
		} else {
			break
		}
	}

	return resRook
}

func bishopCalc(whiteBishopPosition, unionBlackMask, unionWhiteMaskWithoutBishop uint64) uint64 {
	startPosOfBishop := -1
	for i := whiteBishopPosition; i > 0; i = i >> 1 {
		startPosOfBishop++
	}

	resBishop := uint64(0)
	posLU := startPosOfBishop + 7
	posRU := startPosOfBishop + 9
	posLD := startPosOfBishop - 9
	posRD := startPosOfBishop - 7

	for {
		if posLU % 8 != 7 && posLU <= 63 &&
			(uint64(1) << (posLU - 7)) | unionBlackMask != unionBlackMask &&
			((uint64(1) << posLU) | unionWhiteMaskWithoutBishop != unionWhiteMaskWithoutBishop) {
			resBishop = resBishop | (uint64(1) << posLU)
			posLU += 7
		} else {
			break
		}
	}

	for {
		if posRU % 8 != 0 && posRU >= 0 &&
			(uint64(1) << (posRU - 9)) | unionBlackMask != unionBlackMask &&
			((uint64(1) << posRU) | unionWhiteMaskWithoutBishop != unionWhiteMaskWithoutBishop) {
			resBishop = resBishop | (uint64(1) << posRU)
			posRU += 9
		} else {
			break
		}
	}

	for {
		if posLD >= 0 && posLD % 8 != 7 &&
			(uint64(1) << (posLD + 9)) | unionBlackMask != unionBlackMask &&
			((uint64(1) << posLD) | unionWhiteMaskWithoutBishop != unionWhiteMaskWithoutBishop) {
			resBishop = resBishop | (uint64(1) << posLD)
			posLD -= 9
		} else {
			break
		}
	}

	for {
		if posRD >= 0 && posRD % 8 != 0 &&
			(uint64(1) << (posRD + 7)) | unionBlackMask != unionBlackMask &&
			((uint64(1) << posRD) | unionWhiteMaskWithoutBishop != unionWhiteMaskWithoutBishop) {
			resBishop = resBishop | (uint64(1) << posRD)
			posRD -= 7
		} else {
			break
		}
	}

	return resBishop
}
