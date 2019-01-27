package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"meli/dnastats/container"
	"net/http"
	"testing"

	"github.com/dghubble/sling"
)

func TestHandlerGETStats(t *testing.T) {

	req, err := sling.New().Get("http://localhost:8082/stats").Request()
	if err != nil {
		log.Printf("%v", err)
	}

	var client = http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("Cannot GET stats")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	// message := string(body)
	stats := &container.Stat{}
	err = json.Unmarshal(body, &stats)
	if err != nil {
		t.Error("Response not formated properly")
	}

}
