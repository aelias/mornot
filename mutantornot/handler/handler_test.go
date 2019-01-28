package handler

import (
	"io/ioutil"
	"meli/mutantornot/container"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

func TestHandlePOSTMutantHumanDNA(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)

	body := container.DnaMatrix{
		Dna: []string{"ATGC", "TAGC", "CCCC", "GGGG"},
	}

	req, _ := sling.New().Post("/mutant").BodyJSON(body).Request()
	con.Request = req
	HandlePOSTMutant(con)
	if w.Result().StatusCode != http.StatusForbidden {
		t.Error("DNA Matrix is not human. Result code should be 403")
	}
}

func TestHandlePOSTMutantMutantDNA(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)

	body := container.DnaMatrix{
		Dna: []string{"TTTT", "TAGC", "CCCC", "GGGG"},
	}

	req, _ := sling.New().Post("/mutant").BodyJSON(body).Request()
	con.Request = req
	HandlePOSTMutant(con)
	if w.Result().StatusCode != http.StatusOK {
		t.Error("DNA Matrix is MUTANT. Result code should be 200")
	}
}

func TestHandlePOSTMutantInvalidDNA(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)

	body := container.DnaMatrix{
		Dna: []string{},
	}

	req, _ := sling.New().Post("/mutant").BodyJSON(body).Request()
	con.Request = req
	HandlePOSTMutant(con)
	if w.Result().StatusCode != http.StatusNotFound {
		t.Error("DNA Matrix is INVALID. Response code sould be 404")
	}
	// Check message
	data, _ := ioutil.ReadAll(w.Result().Body)
	tempStr := string(data)
	if tempStr != "\"The DNA matrix is invalid\"" {
		t.Error("Message sould be \"The DNA matrix is invalid\"")
	}

}

func TestHandlePOSTMutantInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)

	type IncorrectType struct {
		NoDnaPresent string
	}

	body := IncorrectType{
		NoDnaPresent: "",
	}

	req, _ := sling.New().Post("/mutant").BodyJSON(body).Request()
	con.Request = req
	HandlePOSTMutant(con)
	if w.Result().StatusCode != http.StatusNotFound {
		t.Error("This is an Invalid request")
	}
	// Check message
	data, _ := ioutil.ReadAll(w.Result().Body)
	if string(data) != "\"Invalid request\"" {
		t.Error("Message sould be \"Invalid request\"")
	}

}
