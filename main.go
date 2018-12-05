package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	//"github.com/getlantern/systray"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", "127.0.0.1:9898", "addr to bind to")
)

func openHandler(w http.ResponseWriter, r *http.Request) {
	filename, _ := r.URL.Query()["filename"]
	row, _ := r.URL.Query()["row"]

	if err := exec.Command("/Applications/Sublime Text.app/Contents/SharedSupport/bin/subl", string(filename[0]) + ":" + string(row[0])).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	io.WriteString(w, "ok")
}

func main() {
	//systray.Run(onReady, onExit)
	http.HandleFunc("/open", openHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))

}

func onReady() {
	//systray.SetIcon(getIcon("assets/code.ico"))
	//systray.SetTitle("I'm alive!")
	//systray.SetTooltip("Look at me, I'm a tooltip!")
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}