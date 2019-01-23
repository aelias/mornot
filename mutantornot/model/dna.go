package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"meli/mutantornot/container"
)

// DnaMatrix is a model for a DNA in mongodb
type DnaMatrix struct {
	IsMutant bool     `bson:"is_mutant"`
	UniqueID string   `bson:"unique_id"`
	Dna      []string `bson:"dna"`
}

// LoadFromContainer loads the model from a given container
func (dnaMatrix *DnaMatrix) LoadFromContainer(container container.DnaMatrix) {
	dnaMatrix.Dna = container.Dna
	dnaMatrix.UniqueID = dnaMatrix.GetUniqueID()
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

// Serialize return the struct converted to a string for being saved in the database
func (dnaMatrix *DnaMatrix) Serialize() (serialized string) {
	hash, _ := json.Marshal(dnaMatrix)
	serialized = string(hash)
	return
}

// Deserialize creates a DnaMatrix object from a byte array
func (dnaMatrix *DnaMatrix) Deserialize(serialized string) (deserialized DnaMatrix) {
	err := json.Unmarshal([]byte(serialized), &deserialized)
	if err != nil {
		log.Println("Error deserializing: ", err)
	}
	return
}
