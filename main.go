package main

import (
	"github.com/dkunin/open-in-sublime/server"
	config "github.com/dkunin/open-in-sublime/settings"
	"github.com/dkunin/open-in-sublime/tray"
	"github.com/getlantern/systray"
	"log"
	"net/http"
	"runtime"
)


func main() {
	runtime.GOMAXPROCS(2)

	settings := config.Settings{}
	settings.SetPort("9898")
	settings.SetEditor("sublime")

	defer systray.Run(tray.OnReady(&settings), tray.OnExit)
	go func() {
		http.HandleFunc("/open", server.OpenHandler(&settings))
		http.HandleFunc("/settings", server.SettingsHandler(&settings))
		http.HandleFunc("/settings-json", server.SettingsJsonHandler(&settings))
		err := http.ListenAndServe("127.0.0.1:" + settings.GetPort(), nil)

		if err != nil {
			log.Fatal(err)
		}
	}()
}
