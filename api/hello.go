package api

import (
	"fmt"
	"net/http"

	"github.com/aronhoyer/vercel-go-repro/util"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]

	if !ok || len(name) < 1 {
		name = append(name, "World")
	}

	fmt.Fprint(w, util.Greet(name[0]))
}
