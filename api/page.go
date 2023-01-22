package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MrHuxu/blogo/posts"
)

// Page ...
func Page(w http.ResponseWriter, r *http.Request) {
	names, err := posts.GetNames()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, strings.Join(names, ", "))
}
