package model

import (
	"meli/dnastats/container"
	"testing"
)

func TestNewFormContainer(t *testing.T) {
	con := container.DnaMatrix{
		Dna: []string{"AAAA", "TTTT", "CCCC", "GGGG"},
	}
	model := NewFromContainer(true, con)

	if len(model.Dna) != len(con.Dna) {
		t.Error("Len fot both should be the same")
	}
	for i, str := range model.Dna {
		if str != con.Dna[i] {
			t.Fatal("Elements in the same index should be just equals")
		}
	}
}

func TestGetUniqueID(t *testing.T) {
	expectedUniqueID := "da99d6c998b54eeebeff835c4b8104a6"
	model := DnaMatrix{
		IsMutant: true,
		Dna:      []string{"AAAA", "TTTT", "CCCC", "GGGG"},
	}
	model.UniqueID = model.GetUniqueID()
	if model.UniqueID != expectedUniqueID {
		t.Fatal("It is not the expected uniqueid")
	}
}

func TestDisctintUniqueID(t *testing.T) {
	model1 := DnaMatrix{
		IsMutant: true,
		Dna:      []string{"AAAA", "TTTT", "CCCC", "GGGG"},
	}
	model2 := DnaMatrix{
		IsMutant: true,
		Dna:      []string{"AAAA", "ATTT", "CCCC", "GGGG"},
	}
	model1.UniqueID = model1.GetUniqueID()
	model2.UniqueID = model2.GetUniqueID()

	if model1.UniqueID == model2.UniqueID {
		t.Fatal("UniqueID shouldn't be the same for disctint dna matrix")
	}
}
