package handlers

import (
	"net/http"
)

func PrimHandler(r http.ResponseWriter, w *http.Request) {
	r.Write([]byte("Primary Handelr"))
}
