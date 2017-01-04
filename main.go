package main

import (
	"crypto/dsa"
	"math/big"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var params *dsa.Parameters

const (
	g = "66151740938618228538134463788752921881465865634116427855903259928217197023736720482246784966798915292296790442657509500343711040954802848949865022646557813731571483166610115253689915444465414329516618883440091175584031606804493399092777583782461287015697321400031375804268796659842610014495553794588435255771"
	p = "131222800294604502705481226543618822278818598931593836076045315117623598543398227051785405478070663258892570114762445323344032825659091236986600088795444247135566925897245478896409670574614564066632394221386933786203613371345104104774557513996484736737168115298197440091287272010224976808076237774275400965909"
	q = "812432499621524350136872323481113002945912958851"
)

func main() {
	log.SetLevel(log.DebugLevel)

	params = new(dsa.Parameters)
	params.G, params.P, params.Q = big.NewInt(0), big.NewInt(0), big.NewInt(0)
	params.G.SetString(g, 10)
	params.P.SetString(p, 10)
	params.Q.SetString(q, 10)

	router := mux.NewRouter()
	router.HandleFunc("/sign", sign)
	router.HandleFunc("/generate", generate)
	router.HandleFunc("/verify", verify)
	if err := http.ListenAndServeTLS(":443", "public.crt", "private.key", router); err != nil {
		log.Errorln(err)
	}
}
