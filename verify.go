package main

import (
	"crypto/dsa"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func verify(w http.ResponseWriter, req *http.Request) {
	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Debugln(err)
	}
	defer req.Body.Close()
	h512 := hashDoc(d)
	log.Debugln("hash: ", h512)

	var pubkey dsa.PublicKey
	pubkey.Y = big.NewInt(0)
	pubkey.Parameters = *params
	pubkey.Y.SetString(req.Header.Get("X-Public-Key"), 10)

	r, s := big.NewInt(0), big.NewInt(0)
	log.Debugln("R header: ", req.Header.Get("X-Signature-R"))
	log.Debugln("S header: ", req.Header.Get("X-Signature-S"))

	r.SetString(req.Header.Get("X-Signature-R"), 10)
	s.SetString(req.Header.Get("X-Signature-S"), 10)

	verifystatus := dsa.Verify(&pubkey, h512, r, s)

	log.Debugln("verify result: ", verifystatus)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, strconv.FormatBool(verifystatus)+"\n")
}
