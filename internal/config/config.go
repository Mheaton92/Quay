package config

import "os"

type Keybinds struct {
	Connect 	string `toml:"connect"`
	Add 		string `toml:"add"`
	Edit 		string `toml:"edit"`
	Delete 		string `toml:"delete"`
	Quit 		string `toml:"quit"`
	SCP 		string `toml:"scp"`
	Keys 		string `toml:"keys"`
	Help 		string `toml:"help"`
	Up 			string `toml:"up"`
	Down 		string `toml:"down"`
	NextTab 	string `toml:"next_tab"`
	PrevTab 	string `toml:"prev_tab"`
	Networking 	string `toml:"networking"`
	Save 		string `toml:"save"`
	ForceQuit 	string `toml:"force_quit"`
	NavLeft 	string `toml:"nav_left"`
	NavRight 	string `toml:"nav_right"`
	Test	string "toml:test"
}

func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/.config/quay", nil
}

func EnsureConfigDir() error {
	path, err := ConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(path, 0700)
}

func DefaultKeybinds() Keybinds {
	return Keybinds{
		Connect:	"enter",
		Add:		"a",
		Edit:		"e",
		Delete:		"d",
		Quit:		"q",
		SCP:		"s",
		Keys:		"K",
		Help:		"?",
		Up:			"k",
		Down:		"j",
		NextTab:	"tab",
		PrevTab:	"shift+tab",
		Networking:	"N",
		Save:		"ctrl+s",
		ForceQuit:	"ctrl+c",
		NavLeft:	"left",
		NavRight:	"right",
		Test:	"test",
	}
}