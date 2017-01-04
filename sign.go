package main

import (
	"crypto/dsa"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type sig struct {
	R string
	S string
}

func sign(w http.ResponseWriter, req *http.Request) {
	d, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	h512 := hashDoc(d)
	log.Debugln("req.Body hash: ", h512)

	privatekey := new(dsa.PrivateKey)
	privatekey.PublicKey.Parameters = *params
	privatekey.X = big.NewInt(0)
	privatekey.PublicKey.Y = big.NewInt(0)
	privatekey.X.SetString(req.Header.Get("X-Private-Key"), 10)
	privatekey.PublicKey.Y.SetString(req.Header.Get("X-Public-Key"), 10)

	r, s, err := dsa.Sign(rand.Reader, privatekey, h512)
	log.Debugln("r:", r)
	log.Debugln("s:", s)
	if err != nil {
		log.Errorln("Error signing document: ", err)
	}

	signature := sig{r.String(), s.String()}
	log.Debugln(signature)
	jData, err := json.Marshal(signature)
	if err != nil {
		log.Debugln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
