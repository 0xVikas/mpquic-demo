package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	_ "net/http/pprof"

	"github.com/lucas-clemente/quic-go/h2quic"
	"github.com/gorilla/mux"
)

type binds []string

func (b binds) String() string {
	return strings.Join(b, ",")
}

func (b *binds) Set(v string) error {
	*b = strings.Split(v, ",")
	return nil
}

func handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/media/{mId:[0-9]+}/stream/", streamHandler).Methods("GET")
	router.HandleFunc("/media/{mId:[0-9]+}/stream/{segName:index[0-9]+.ts}", streamHandler).Methods("GET")
	fs := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return router
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/index.html")
	w.Header().Set("Cache-Control", "public, max-age=0")
}

func streamHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	mId, err := strconv.Atoi(vars["mId"])
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	segName, ok := vars["segName"]
	if !ok {
		mediaBase := getMediaBase(mId)
		m3u8Name := "index.m3u8"
		serveHlsM3u8(response, request, mediaBase, m3u8Name)
	} else {
		mediaBase := getMediaBase(mId)
		serveHlsTs(response, request, mediaBase, segName)
	}
}

func getMediaBase(mId int) string {
	mediaRoot := "assets/media"
	return fmt.Sprintf("%s/%d", mediaRoot, mId)
}

func serveHlsM3u8(w http.ResponseWriter, r *http.Request, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/hls/%s", mediaBase, m3u8Name)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "application/x-mpegURL")
	w.Header().Set("Cache-Control", "public, max-age=10")
}

func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/hls/%s", mediaBase, segName)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "video/MP2T")
	w.Header().Set("Cache-Control", "public, max-age=10")
}

func main() {

	certDir :="./assets/certificates"
	certFile := certDir+"/fullchain.pem"
	keyFile := certDir+"/privkey.pem"

	http.Handle("/", handlers())

	bs := binds{"0.0.0.0:6121"}
	
	var wg sync.WaitGroup
	wg.Add(len(bs))
	for _, b := range bs {
		bCap := b
		go func() {
			var err error
			err = h2quic.ListenAndServeQUIC(bCap, certFile, keyFile, nil)
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
