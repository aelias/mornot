package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)
	HandleGETStats(con)
	fmt.Printf("%v", con)
	fmt.Println(con.Keys)
	fmt.Println(w.Result().StatusCode)
	if w.Result().StatusCode != http.StatusOK {
		t.Error("The response sould be 200 - OK")
	}
}
