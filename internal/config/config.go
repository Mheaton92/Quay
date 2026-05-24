package config

import (
	"encoding/json"
	"os"
)

type Keybinds struct {
	Connect       string `toml:"connect"`
	Add           string `toml:"add"`
	Edit          string `toml:"edit"`
	Delete        string `toml:"delete"`
	Quit          string `toml:"quit"`
	SCP           string `toml:"scp"`
	Keys          string `toml:"keys"`
	Help          string `toml:"help"`
	Up            string `toml:"up"`
	Down          string `toml:"down"`
	NextTab       string `toml:"next_tab"`
	PrevTab       string `toml:"prev_tab"`
	Networking    string `toml:"networking"`
	Save          string `toml:"save"`
	ForceQuit     string `toml:"force_quit"`
	NavLeft       string `toml:"nav_left"`
	NavRight      string `toml:"nav_right"`
	PinSession    string `tom1:"pin_session"`
	PinPersistent string `tom1:"pin_persistent"`
	Import				string `toml:"import"`
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
		Connect:       "enter",
		Add:           "a",
		Edit:          "e",
		Delete:        "d",
		Quit:          "q",
		SCP:           "s",
		Keys:          "K",
		Help:          "?",
		Up:            "k",
		Down:          "j",
		NextTab:       "tab",
		PrevTab:       "shift+tab",
		Networking:    "N",
		Save:          "ctrl+s",
		ForceQuit:     "ctrl+c",
		NavLeft:       "left",
		NavRight:      "right",
		PinSession:    "p",
		PinPersistent: "P",
		Import:				 "i",
	}
}

func PinsFile() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return dir + "/pins.json", nil
}

func LoadPins() ([]string, error) {
	path, err := PinsFile()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{}, nil // no pins file yet is fine
	}
	var pins []string
	if err := json.Unmarshal(data, &pins); err != nil {
		return nil, err
	}
	return pins, nil
}

func SavePins(pins []string) error {
	path, err := PinsFile()
	if err != nil {
		return err
	}
	data, err := json.Marshal(pins)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
