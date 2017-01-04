package main

import (
	"crypto/sha512"
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func httpResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")

	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(body)
	}
}

func hashDoc(doc []byte) []byte {
	h512 := sha512.New()
	h512.Write(doc)
	log.Debugln(h512.Size())
	return h512.Sum(nil)
}
