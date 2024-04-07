package Handlers

import (
	"fmt"
	"net/http"
)

func ProfileHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Profile_page lol")
}
