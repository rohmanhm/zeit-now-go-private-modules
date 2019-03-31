package handler

import (
	"fmt"
	"net/http"

	msg "github.com/rohmanhm/go-private/message"
)

// Handler func
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, msg.SayHi("rohman"))
}
