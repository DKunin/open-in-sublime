package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type Settings struct {
	port string
	editor string
}

func (s *Settings) setEditor(editor string) {
	s.editor = editor
}

func (s *Settings) setPort(port string) {
	s.port = port
}

func (s *Settings) getSettings() string {
	return "{\"port\": "+ s.port + ", \"editor\": \"" + s.editor + "\"}"
}

func renderCommand(settings *Settings) string {
	// vim ~/Projects/release-notification-service/app/modules/scheduler.js \\"+call cursor(10, 10)\\"

	switch settings.editor {
	case "sublime":
		return "/Applications/Sublime Text.app/Contents/SharedSupport/bin/subl"
	case "goland":
		return "/Applications/GoLand.app/Contents/MacOS/goland"
	default:
		return ""
	}

}

func renderOptions(settings *Settings, filename string, row string) string {
	switch settings.editor {
	case "sublime":
		return  filename + ":" + row
	case "goland":
		return " --line "+ row + " " + filename
	default:
		return ""
	}
}


func openHandler(settings *Settings) func(w http.ResponseWriter, r *http.Request) {
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

func settingsJsonHandler(settings *Settings) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(settings.getSettings()))
	}
}

func settingsHandler(settings *Settings) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				w.WriteHeader(500)
				log.Printf("error with form %s", err)
				return
			}

			settings.setEditor(r.PostFormValue("editor"))
			settings.setPort(r.PostFormValue("port"))

			w.WriteHeader(200)
			w.Write([]byte(settings.getSettings()))
			return
		}

		b, err := ioutil.ReadFile("assets/settings.html")

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error reading file"))
			log.Fatal(err)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(b))
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	settings := Settings{"9898", "vim"}
	defer systray.Run(onReady(&settings), onExit)
	go func() {
		http.HandleFunc("/open", openHandler(&settings))
		http.HandleFunc("/settings", settingsHandler(&settings))
		http.HandleFunc("/settings-json", settingsJsonHandler(&settings))
		err := http.ListenAndServe("127.0.0.1:" + settings.port, nil)

		if err != nil {
			log.Fatal(err)
		}
	}()


}

func onReady(settings *Settings) func() {
	return func() {
		systray.SetIcon(getIcon("assets/icon.ico"))

		mCurrentEditor := systray.AddMenuItem("Editor: " + settings.editor, "")
		mCurrentEditor.Disable()
		mCurrentPort := systray.AddMenuItem("PORT: " + settings.port, "")
		mCurrentPort.Disable()
		systray.AddSeparator()
		mSettings := systray.AddMenuItem("Settings", "")
		//mStop := systray.AddMenuItem("Stop Server", "")
		mQuit := systray.AddMenuItem("Quit", "")

		go func() {
			for {
				select {
				case <-mSettings.ClickedCh:
					fmt.Printf("%s", settings.editor)
					open.Run("http://127.0.0.1:" + settings.port + "/settings")
				//case <-mStop.ClickedCh:
					//settings.setEditor("vim")

				case <-mQuit.ClickedCh:
					systray.Quit()
					return
				}
			}
		}()
	}

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