package main

import (
	"crypto/dsa"
	"crypto/rand"
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

//GenResponse to return the json response with key pair
type GenResponse struct {
	PublicKey  string
	PrivateKey string
}

func generate(w http.ResponseWriter, req *http.Request) {
	privatekey := new(dsa.PrivateKey)
	privatekey.PublicKey.Parameters = *params
	dsa.GenerateKey(privatekey, rand.Reader)
	log.Debugln("Priv Key: ", privatekey.X)
	log.Debugln("Pub Key: ", privatekey.PublicKey.Y)

	resp := GenResponse{privatekey.PublicKey.Y.String(), privatekey.X.String()}
	log.Debugln("Response: ", resp)

	jData, err := json.Marshal(resp)
	if err != nil {
		log.Debugln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
