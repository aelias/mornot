package util

import (
	"errors"
	"strings"
)

func isValidDNAMatrix(matrix []string) (isValid bool) {
	matrixLen := len(matrix)
	isValid = true
	for _, str := range matrix {
		if matrixLen != len(str) {
			isValid = false
			break
		}
		// Only allows the desired characters
		for _, chr := range str {
			if !strings.Contains("ATCG", string(chr)) {
				isValid = false
				break
			}
		}
	}
	return
}

func searchCoincidences(str string) (counter int) {
	counter += strings.Count(str, "AAAA")
	counter += strings.Count(str, "TTTT")
	counter += strings.Count(str, "CCCC")
	counter += strings.Count(str, "GGGG")
	return
}

func searchHorizontal(dna []string) (counter int) {
	for _, str := range dna {
		counter += searchCoincidences(str)
	}
	return
}

func searchVertical(dna []string) (counter int) {

	verticalDNA := make([]strings.Builder, len(dna))
	for _, str := range dna {
		for i := 0; i < len(dna); i++ {
			verticalDNA[i].WriteByte(str[i])
		}
	}

	for i := 0; i < len(dna); i++ {
		counter += searchCoincidences(verticalDNA[i].String())
	}

	return
}

func searchDiagonal(dna []string) (counter int) {
	diagonalSize := len(dna)
	var sb strings.Builder
	// From left to right, bottom triangle
	fromY := 0
	for diagonalSize >= 4 {
		fromX := 0
		for y := fromY; y < len(dna); y++ {
			sb.WriteByte(dna[y][fromX])
			fromX++
		}
		diagonal := sb.String()
		counter += searchCoincidences(diagonal)
		sb.Reset()
		diagonalSize--
		fromY++
	}

	// From left to right, top triangle
	diagonalSize = len(dna) - 1
	fromX := 1
	for diagonalSize >= 4 {
		fromY := 0
		for x := fromX; x < len(dna); x++ {
			sb.WriteByte(dna[fromY][x])
			fromY++
		}
		diagonal := sb.String()
		counter += searchCoincidences(diagonal)
		sb.Reset()
		diagonalSize--
		fromX++
	}

	// From right to left, bottom triangle
	fromY = 0
	diagonalSize = len(dna)
	for diagonalSize >= 4 {
		fromX := len(dna) - 1
		for y := fromY; y < len(dna); y++ {
			sb.WriteByte(dna[y][fromX])
			fromX--
		}
		diagonal := sb.String()
		counter += searchCoincidences(diagonal)
		sb.Reset()
		diagonalSize--
		fromY++
	}

	// From right to left, top triangle
	diagonalSize = len(dna) - 1
	fromX = len(dna) - 2
	for diagonalSize >= 4 {
		fromY := 0
		for x := fromX; x >= 0; x-- {
			sb.WriteByte(dna[fromY][x])
			fromY++
		}
		diagonal := sb.String()
		counter += searchCoincidences(diagonal)
		sb.Reset()
		diagonalSize--
		fromX--
	}

	return
}

// IsMutant detemine if a DNA secuence belongs to a mutant being
func IsMutant(dna []string) (mutant bool, err error) {

	// Ensure valid matrix
	if len(dna) < 4 || !isValidDNAMatrix(dna) {

		return false, errors.New("DNA matrix is not valid")
	}

	counter := searchHorizontal(dna)
	counter += searchVertical(dna)
	counter += searchDiagonal(dna)

	mutant = false
	if counter >= 3 {
		mutant = true
	}
	return
}
