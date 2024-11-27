package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

type oneFileServer struct {
	name    string
	ct      string
	content []byte
}

func (fs *oneFileServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("cache-control", "no-store")
	if req.URL.Path == "/favicon.ico" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if base := path.Base(fs.name); req.URL.Path != "/"+base {
		http.Redirect(w, req, "/"+base, http.StatusSeeOther)
		return
	}
	w.Header().Add("content-type", fs.ct)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(fs.content); err != nil {
		log.Printf("error writing content: %v", err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: serveme <file or directory>")
		os.Exit(3)
	}
	server := &http.Server{Addr: "0.0.0.0:1234"}
	file := os.Args[1]
	if st, err := os.Stat(file); err != nil {
		panic(err)
	} else if st.IsDir() {
		server.Handler = http.FileServer(http.Dir(file))
	} else {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read: %v", err)
			os.Exit(1)
		}
		kind := http.DetectContentType(content)
		fmt.Fprintf(os.Stderr, "detected content type: %s\n", kind)
		server.Handler = &oneFileServer{name: path.Base(file), ct: kind, content: content}
	}
	if err := server.ListenAndServe(); err != nil {
		log.Printf("http server exited: %v", err)
	}
}
