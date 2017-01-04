package main

import (
	"crypto/dsa"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Generate(t *testing.T) {
	Convey("Generate pub/private certs", t, func() {
		params = new(dsa.Parameters)
		params.G, params.P, params.Q = big.NewInt(0), big.NewInt(0), big.NewInt(0)
		params.G.SetString(g, 10)
		params.P.SetString(p, 10)
		params.Q.SetString(q, 10)

		req, err := http.NewRequest("GET", "/generate", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(generate)
		handler.ServeHTTP(rr, req)

		So(rr.Code, ShouldEqual, http.StatusOK)
		So(rr.Body.String(), ShouldContainSubstring, "PublicKey")
		So(rr.Body.String(), ShouldContainSubstring, "PrivateKey")
	})
}
