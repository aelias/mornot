package handler

import (
	"log"
	"meli/mutantornot/container"
	"net/http"
	"testing"

	"github.com/dghubble/sling"
)

func TestMutantHandlerNotMutantDNA(t *testing.T) {
	body := container.DnaMatrix{
		Dna: []string{"ATGC", "TAGC", "CCCC", "GGGG"},
	}
	req, err := sling.New().Post("http://localhost/mutant").BodyJSON(body).Request()
	if err != nil {
		log.Printf("%v", err)
	}
	var client = http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusForbidden {
		t.Error("Incorrect message")
	}

}

func TestMutantHandlerMutantDNA(t *testing.T) {
	body := container.DnaMatrix{
		Dna: []string{"ATGC", "TTTT", "CCCC", "GGGG"},
	}
	req, err := sling.New().Post("http://localhost/mutant").BodyJSON(body).Request()
	if err != nil {
		log.Printf("%v", err)
	}
	var client = http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("DNA is mutant. Response status should be 200 - OK")
	}

}

func TestMutantHandlerInvalidMatrix(t *testing.T) {
	body := container.DnaMatrix{
		Dna: []string{"ATG", "TTT", "CCC", "GGG"},
	}
	req, err := sling.New().Post("http://localhost/mutant").BodyJSON(body).Request()
	if err != nil {
		log.Printf("%v", err)
	}
	var client = http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Error("DNA is mutant. Response status should be 200 - OK")
	}

}

func TestMutantHandlerInvalidRequest(t *testing.T) {
	// Data type for emulate an invalid POST request
	type NotValid struct {
		NotValid string
	}

	body := NotValid{
		NotValid: "",
	}

	req, err := sling.New().Post("http://localhost/mutant").BodyJSON(body).Request()
	if err != nil {
		log.Printf("%v", err)
	}
	var client = http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Error("Not valid request accepted: ", resp.StatusCode)
	}

}
