package tray

import (
	"fmt"
	config "github.com/dkunin/open-in-sublime/settings"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
)

func OnReady(settings *config.Settings) func() {
	return func() {
		systray.SetIcon(getIcon("assets/icon.ico"))

		mCurrentEditor := systray.AddMenuItem("Editor: " + settings.GetEditor(), "")
		mCurrentEditor.Disable()
		mCurrentPort := systray.AddMenuItem("PORT: " + settings.GetPort(), "")
		mCurrentPort.Disable()
		systray.AddSeparator()
		mSettings := systray.AddMenuItem("Settings", "")
		//mStop := systray.AddMenuItem("Stop Server", "")
		mQuit := systray.AddMenuItem("Quit", "")

		go func() {
			for {
				select {
				case <-mSettings.ClickedCh:
					fmt.Printf("%s", settings.GetEditor())
					open.Run("http://127.0.0.1:" + settings.GetPort() + "/settings")
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

func OnExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}