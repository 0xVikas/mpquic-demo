package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	quic "github.com/lucas-clemente/quic-go"

	"github.com/lucas-clemente/quic-go/h2quic"
//	"github.com/lucas-clemente/quic-go/internal/utils"
)

func init(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

//		utils.SetLogTimeFormat("")

		quicConfig := &quic.Config{
			CreatePaths: true,
		}

		hclient := &http.Client{
			Transport: &h2quic.RoundTripper{QuicConfig: quicConfig},
		}

		addr := "https://localhost:6121" + r.URL.Path

//		utils.Infof("GET %s", addr)

			rsp, err := hclient.Get(addr)
			if err != nil {
				panic(err)
			}
//			utils.Infof("Got response for %s: %#v", addr, rsp)

			body, err := ioutil.ReadAll(rsp.Body)

			if err != nil {
				fmt.Printf("error reading body while handling /echo: %s\n", err.Error())
			}
			w.Write(body)

	})
}

func main() {

	http.ListenAndServe("0.0.0.0:1337", nil)

}
