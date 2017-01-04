package main

import (
	"crypto/dsa"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Verify(t *testing.T) {
	Convey("Verify Signature", t, func() {
		params = new(dsa.Parameters)
		params.G, params.P, params.Q = big.NewInt(0), big.NewInt(0), big.NewInt(0)
		params.G.SetString(g, 10)
		params.P.SetString(p, 10)
		params.Q.SetString(q, 10)

		req, err := http.NewRequest("POST", "/verify",
			strings.NewReader("Some dummy text to sign"))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("X-Public-Key", "95844969447922433977339473624550925036347612000276854027151969825972546689660825290553136028027736298093359581423836527778446402317015414531334103284121961520585680585161943766567116924878902577361757013847580356733745521762681741114500536993923329936051538811898183055511862474696825601113862077149209590782")
		req.Header.Add("X-Signature-S", "476268709127830776100463584598141902668545680448")
		req.Header.Add("X-Signature-R", "403810297134553044788220380459807600595922071859")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(verify)
		handler.ServeHTTP(rr, req)

		So(rr.Code, ShouldEqual, http.StatusOK)
		So(rr.Body.String(), ShouldEqual, "true\n")

		req, err = http.NewRequest("POST", "/verify",
			strings.NewReader("Some mangled dummy text to sign"))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("X-Public-Key", "95844969447922433977339473624550925036347612000276854027151969825972546689660825290553136028027736298093359581423836527778446402317015414531334103284121961520585680585161943766567116924878902577361757013847580356733745521762681741114500536993923329936051538811898183055511862474696825601113862077149209590782")
		req.Header.Add("X-Signature-S", "476268709127830776100463584598141902668545680448")
		req.Header.Add("X-Signature-R", "403810297134553044788220380459807600595922071859")
		rr = httptest.NewRecorder()
		handler = http.HandlerFunc(verify)
		handler.ServeHTTP(rr, req)
		So(rr.Code, ShouldEqual, http.StatusOK)
		So(rr.Body.String(), ShouldEqual, "false\n")
	})
}
