package server

import (
	"fmt"
	config "github.com/dkunin/open-in-sublime/settings"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func renderCommand(settings *config.Settings) string {
	switch settings.GetEditor() {
	case "sublime":
		return "/Applications/Sublime Text.app/Contents/SharedSupport/bin/subl"
	case "goland":
		return "/Applications/GoLand.app/Contents/MacOS/goland"
	default:
		return ""
	}

}

func renderOptions(settings *config.Settings, filename string, row string) string {
	switch settings.GetEditor() {
	case "sublime":
		return  filename + ":" + row
	case "goland":
		return " --line "+ row + " " + filename
	default:
		return ""
	}
}


func OpenHandler(settings *config.Settings) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filename, _ := r.URL.Query()["filename"]
		row, _ := r.URL.Query()["row"]
		log.Printf("render options: %s", renderOptions(settings, string(filename[0]), string(row[0])))
		if err := exec.Command(renderCommand(settings), renderOptions(settings, string(filename[0]), string(row[0]))).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func SettingsJsonHandler(settings *config.Settings) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(settings.GetSettings()))
	}
}

//func SettingsHandler(settings *config.Settings) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == "POST" {
//			err := r.ParseForm()
//			if err != nil {
//				w.WriteHeader(500)
//				log.Printf("error with form %s", err)
//				return
//			}
//
//			settings.SetEditor(r.PostFormValue("editor"))
//			//settings.SetPort(r.PostFormValue("port"))
//
//			w.WriteHeader(200)
//			w.Write([]byte(settings.GetSettings()))
//			return
//		}
//
//		b, err := ioutil.ReadFile("assets/settings.html")
//
//		if err != nil {
//			w.WriteHeader(500)
//			w.Write([]byte("Error reading file"))
//			log.Fatal(err)
//			return
//		}
//		w.WriteHeader(200)
//		w.Write([]byte(b))
//	}
//}