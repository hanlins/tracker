package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// original functions
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/dev", DevHandler)
	r.HandleFunc("/session", SessionCreateHandler)
	r.HandleFunc("/listen", ListenHandler)
	r.HandleFunc("/test", TestHandler)

	// tracker handlers goes here
	r.HandleFunc("/a/{script:[a-z]+}.js", TrackerAHandler)
	r.HandleFunc("/listenera", ListenerAHandler)

	r.HandleFunc("/b/{file}", TrackerBHandler)

	r.HandleFunc("/c/setcookiec", TrackerCSetCookie)
	r.HandleFunc("/c/{file:[a-z]+}.html", TrackerCHtmlHandler)
	r.HandleFunc("/c/{script:[a-z]+}.js", TrackerCScriptHandler)

	r.HandleFunc("/listenerd", ListenerDHandler)
	r.HandleFunc("/d/{script:[a-z]+}.js", TrackerDScriptHandler)

	r.HandleFunc("/e/home", TrackerEHomeHandler)
	r.HandleFunc("/e/{file}", TrackerEHandler)

	r.HandleFunc("/f/{file}", TrackerFHandler)

	// handle result request pages for testbed
	r.HandleFunc("/result/{file}", ResultRequestHandler)

	// Start listening on the given IP address and port
	http.Handle("/", r)
	var httpListenAddr = fmt.Sprintf("%s:%d",
		"127.0.0.1",
		8002)
	if err := http.ListenAndServe(httpListenAddr, nil); err != nil {
		log.Fatalf("Could not start HTTP server listening: %s\n", err)
	}
}

// original functions
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./index.html")
	ServeFileWrapper(w, r, "./index.html")
}

func DevHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./dev.html")
	ServeFileWrapper(w, r, "./dev.html")
}

func SessionCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	r.ParseForm()
	fmt.Println("username:", r.Form["uname"])
	fmt.Println("password:", r.Form["passwd"])
}

func ListenHandler(w http.ResponseWriter, r *http.Request) {
	// http://mylab.com/listen?login=param_value
	//	httpHeader := http.Header(r.Header)
	//	log.Println(httpHeader)
	param1 := r.URL.Query().Get("login")
	fmt.Println(param1)
	// http.ServeFile(w, r, "./static/1.png")
	ServeFileWrapper(w, r, "./static/1.png")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./test.html")
	ServeFileWrapper(w, r, "./test.html")
}

func ResultRequestHandler(w http.ResponseWriter, r *http.Request) {
	file := mux.Vars(r)["file"]
	// http.ServeFile(w, r, "../result/"+file)
	ServeFileWrapper(w, r, "../result/"+file)
}
