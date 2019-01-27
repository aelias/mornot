package util

import (
	"testing"
)

func TestSearchCoincidences(t *testing.T) {
	if searchCoincidences("JJJJ") > 0 {
		t.Error("This shouldn't be a coincidence")
		t.Fail()
	}
}

func TestSearchCoincidencesGoodString(t *testing.T) {
	if searchCoincidences("CCCC") == 0 {
		t.Error("Should be a coincidence CCCC")
	}
}

func TestHumanDNA(t *testing.T) {
	// With
	dnaHumanSample := []string{"AAAAG", "CTCTT", "TGTGT", "ATATA", "GGTTT"}
	// Do
	if isMutant, _ := IsMutant(dnaHumanSample); isMutant {
		t.Error("DNA sample is not from a mutant")
	}
}

func TestMutantDNA(t *testing.T) {
	// With
	dnaMutantSample := []string{"AAAA", "CCCC", "TGTG", "TTTT"}
	// Do
	if isMutant, _ := IsMutant(dnaMutantSample); !isMutant {
		t.Error("DNA sample IS from a mutant")
	}

}

func TestNotValidMatrixNoCuadratic(t *testing.T) {
	invalidMatrix := []string{"AA", "BBB", "CC"}
	if isValidDNAMatrix(invalidMatrix) {
		t.Error("Matrix is not a valid one!")
	}

}

func TestNotValidMatrixNoDNA(t *testing.T) {
	invalidMatrix := []string{"AAA", "BBB", "CCC"}
	if isValidDNAMatrix(invalidMatrix) {
		t.Error("Matrix is not a valid one!")
	}

}

func TestValidMatrix(t *testing.T) {
	invalidMatrix := []string{"AAA", "TTT", "CCC"}
	if !isValidDNAMatrix(invalidMatrix) {
		t.Error("Matrix is not a valid one!")
	}

}

func TestInvalidSize(t *testing.T) {
	invalidMatrix := []string{"AAA", "TTT", "CCC"}
	if _, err := IsMutant(invalidMatrix); err == nil {
		t.Error("Matrix is not a valid one!")
	}

}
