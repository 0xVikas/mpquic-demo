package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	quic "github.com/lucas-clemente/quic-go"

	"github.com/lucas-clemente/quic-go/h2quic"
)

// Function for printing the HTTP Request
func formatRequest(r *http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
	  name = strings.ToLower(name)
	  for _, h := range headers {
		request = append(request, fmt.Sprintf("%v: %v", name, h))
	  }
	}

	if r.Method == "POST" {
	   r.ParseForm()
	   request = append(request, "\n")
	   request = append(request, r.Form.Encode())
	} 
	 return strings.Join(request, "\n")
}

func init(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(formatRequest(r))

		quicConfig := &quic.Config{
			CreatePaths: true,
		}

		hclient := &http.Client{
			Transport: &h2quic.RoundTripper{QuicConfig: quicConfig},
		}

		addr := "https://localhost:6121" + r.URL.Path

		rsp, err := hclient.Get(addr)
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(rsp.Body)

		if err != nil {
			fmt.Printf("error reading body while handling /echo: %s\n", err.Error())
		}

		w.Write(body)

	})
}

func main() {

	fmt.Println("Started the client on : localhost:1337")
	http.ListenAndServe("0.0.0.0:1337", nil)

}
