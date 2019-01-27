package model

import (
	"crypto/md5"
	"encoding/hex"
	"meli/dnastats/container"
)

// DnaMatrix is a model for a DNA in mongodb
type DnaMatrix struct {
	IsMutant bool     `bson:"is_mutant"`
	UniqueID string   `bson:"unique_id"`
	Dna      []string `bson:"dna"`
}

// NewFromContainer creates a new DnaMatrix model from a given container
func NewFromContainer(isMutant bool, container container.DnaMatrix) DnaMatrix {
	var dnaModel DnaMatrix
	dnaModel.IsMutant = isMutant
	dnaModel.Dna = container.Dna
	dnaModel.UniqueID = dnaModel.GetUniqueID()
	return dnaModel
}

// GetUniqueID calculate a uniqueID for a DNA matrix
func (dnaMatrix *DnaMatrix) GetUniqueID() (uniqueID string) {
	hash := md5.New()
	for _, str := range dnaMatrix.Dna {
		hash.Write([]byte(str))
	}
	uniqueID = hex.EncodeToString(hash.Sum(nil))
	dnaMatrix.UniqueID = uniqueID
	return
}
