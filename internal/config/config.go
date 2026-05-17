package config

import "os"

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
