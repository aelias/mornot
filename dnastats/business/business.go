package business

import (
	"log"
	"meli/dnastats/model"
	"meli/rabbit"
	"sync"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Database constants
const (
	DbName       = "dnastats"
	DbCollection = "dna"
)

// MongoDB constants
var saveMutex sync.Mutex
var mongoSession *mgo.Session
var dnaMongoCol *mgo.Collection

func init() {
	// Initialize mongodb
	log.Println("Initialize module business")
	saveMutex = sync.Mutex{}

	var err error
	mongoSession, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	// Get the proper collection
	dnaMongoCol = mongoSession.DB(DbName).C(DbCollection)
	// Ensure the unique index exists
	uniqueIDIndex := mgo.Index{
		Key:        []string{"unique_id"},
		Unique:     true,
		DropDups:   true,
		Background: false, // Need to have unique objects
		Sparse:     true,
	}
	dnaMongoCol.EnsureIndex(uniqueIDIndex)
}

// SaveDNAIfNotExists save the dna in the database checking if it exists
func SaveDNAIfNotExists(dnaMessage rabbit.DnaMessage) {
	saveMutex.Lock()
	// Prepare the record to be added
	var dnaModel model.DnaMatrix
	dnaModel.IsMutant = dnaMessage.IsMutant
	dnaModel.Dna = dnaMessage.Dna
	dnaModel.UniqueID = dnaModel.GetUniqueID()
	err := dnaMongoCol.Insert(dnaModel)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Saved DNA: %v\n", dnaModel)
	}
	// Free mutex
	saveMutex.Unlock()
}

// SaveDNAIfNotExists save the dna in the database checking if it exists
// func SaveDNAIfNotExists(isMutant bool, container container.DnaMatrix) {
// 	// Only one dna can be saved at the same time, for ensure uniqueness
// 	saveMutex.Lock()
// 	// Prepare the record to be added
// 	dnaModel := model.NewFromContainer(isMutant, container)
// 	// Save the record
// 	err := dnaMongoCol.Insert(dnaModel)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		log.Printf("Saved DNA: %v\n", dnaModel)
// 	}
// 	// Free mutex
// 	saveMutex.Unlock()
// }

// GetDNAStats returns the amount of mutant and human DNA in the collection
func GetDNAStats() (mutants int, humans int, ratio float64) {
	mutants, _ = dnaMongoCol.Find(bson.M{"is_mutant": true}).Count()
	humans, _ = dnaMongoCol.Find(bson.M{"is_mutant": false}).Count()
	ratio = 0
	// Be careful with cero division
	if humans != 0 {
		ratio = float64(mutants) / float64(humans)
	}
	return
}
