package settings

import (
	"fmt"
	"testing"
)

func TestSettings_GetEditor(t *testing.T) {
	settings := Settings{}
	settings.SetEditor("sublime")

	if settings.GetEditor() != "sublime" {
		t.Fatal("wrong editor")
	}
}

func TestSettings_GetSettings(t *testing.T) {
	settings := Settings{}
	settings.SetEditor("sublime")
	settings.SetPort("9898")
	fmt.Println(settings.GetSettings())
	if settings.GetSettings() != "{\"port\": 9898, \"editor\": \"sublime\"}" {
		t.Fatal("wrong json")
	}
}

func TestSettings_GetPort(t *testing.T) {
	settings := Settings{}
	settings.SetPort("9898")

	if settings.GetPort() != "9898" {
		t.Fatal("wrong port")
	}
}