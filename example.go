package main

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"
	"math/big"

	log "github.com/Sirupsen/logrus"
)

const (
	privkeyConst = "336471247502148589687947524635172263994733687355"
	pubkeyConst  = "95844969447922433977339473624550925036347612000276854027151969825972546689660825290553136028027736298093359581423836527778446402317015414531334103284121961520585680585161943766567116924878902577361757013847580356733745521762681741114500536993923329936051538811898183055511862474696825601113862077149209590782"
)

func mainEx() {
	log.SetLevel(log.DebugLevel)
	params = new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	//params = new(dsa.Parameters)
	// if err := dsa.GenerateParameters(params, rand.Reader, dsa.L2048N256); err != nil {
	// 	log.Errorln(err)
	// }
	params.G, params.P, params.Q = new(big.Int), new(big.Int), new(big.Int)
	params.G.SetString(g, 10)
	params.P.SetString(p, 10)
	params.Q.SetString(q, 10)

	privatekey := new(dsa.PrivateKey)
	privatekey.X = new(big.Int)
	privatekey.X.SetString(privkeyConst, 10)
	privatekey.PublicKey.Parameters = *params
	privatekey.PublicKey.Y = new(big.Int)
	privatekey.PublicKey.Y.SetString(pubkeyConst, 10)
	pubkey := privatekey.PublicKey
	// pubkey.G.SetString(g, 10)
	// pubkey.P.SetString(p, 10)
	// pubkey.Q.SetString(q, 10)

	// privatekey2 := new(dsa.PrivateKey)
	// privatekey2.PublicKey.Parameters = *params
	//dsa.GenerateKey(privatekey2, rand.Reader) // this generates a public & private key pair

	// log.Debugln("My key: ", privatekey)
	// log.Debugln("Other key: ", privatekey2)
	// privatekey.Parameters = *params
	// privatekey.G, privatekey.P, privatekey.Q = new(big.Int), new(big.Int), new(big.Int)
	//
	// privatekey.G.SetString(g, 10)
	// privatekey.P.SetString(p, 10)
	// privatekey.Q.SetString(q, 10)
	// privatekey.Y.SetString(pubkeyConst, 10)
	//
	//var pubkey dsa.PublicKey
	// log.Debugln("Priv Key struct: ", privatekey)
	// log.Debugln("Priv Key struct: ", privatekey.Parameters)
	// log.Debugln("Priv Key struct: ", privatekey.PublicKey)
	// log.Debugln("Priv Key struct: ", privatekey.X)

	// fmt.Println("Private Key :")
	// fmt.Printf("%x \n", privatekey)
	//
	// fmt.Println("Public Key :")
	// fmt.Printf("%x \n", pubkey)

	// Sign
	// var h hash.Hash
	// h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	h := sha512.New()
	io.WriteString(h, "This idjfldjvldnvs the message to be signed and verified!")
	log.Debugln("hash size", h.Size())
	signhash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("Signature : %x\n", signature)

	r2, s2, err := dsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
	}

	verifystatus := dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be true

	privatekey.PublicKey.Parameters = *params

	signature2 := r2.Bytes()
	signature2 = append(signature2, s2.Bytes()...)

	fmt.Printf("Signature : %x\n", signature2)

	// Verify

	verifystatus2 := dsa.Verify(&pubkey, signhash, r2, s2)
	fmt.Println(verifystatus2) // should be true

	// we add additional data to change the signhash
	io.WriteString(h, "This message is NOT to be signed and verified!")
	signhash = h.Sum(nil)

	verifystatus = dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be false
}
