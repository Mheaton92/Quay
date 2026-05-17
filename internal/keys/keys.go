package keys

import (
	"os"
	"path/filepath"
	"strings"
	"os/exec"
	"fmt"
)

type Key struct {
	Name string
	Type string
	Fingerprint string
	PublicKey string
	Path string
}


func List() ([]Key, error) {
    homeDir := os.Getenv("HOME")
    sshDir := filepath.Join(homeDir, ".ssh")
    
    entries, err := os.ReadDir(sshDir)
    if err != nil {
        return nil, err
    }
    
    var keys []Key
    for _, entry := range entries {
        if !strings.HasSuffix(entry.Name(), ".pub") {
            continue
        }
        pubPath := filepath.Join(sshDir, entry.Name())
        content, err := os.ReadFile(pubPath)
        if err != nil {
            continue
        }
        parts := strings.Fields(string(content))
        if len(parts) < 2 {
            continue
        }
		k := Key{
		Name:      strings.TrimSuffix(entry.Name(), ".pub"),
		Type:      parts[0],
		PublicKey: strings.TrimSpace(string(content)),
		Path:      strings.TrimSuffix(pubPath, ".pub"),
	}

	// Get fingerprint
	if fp, err := GetFingerprint(k); err == nil {
		k.Fingerprint = fp
	}

	keys = append(keys, k)
    }
    return keys, nil
}

func Generate(name string, keyType string, comment string) error {
    homeDir := os.Getenv("HOME")
    keyPath := filepath.Join(homeDir, ".ssh", name)
    
    args := []string{
        "-t", keyType,
        "-f", keyPath,
        "-C", comment,
        "-N", "",  // empty passphrase for now
    }
    
    cmd := exec.Command("ssh-keygen", args...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func Deploy(key Key, host string, user string, port int) error {
    args := []string{
        "-i", key.Path + ".pub",
        "-p", fmt.Sprintf("%d", port),
        user + "@" + host,
    }
    
    cmd := exec.Command("ssh-copy-id", args...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func CopyToClipboard(key Key) error {
    commands := [][]string{
        {"xclip", "-selection", "clipboard"},
        {"xsel", "--clipboard", "--input"},
        {"wl-copy"},
    }
    
    for _, args := range commands {
        cmd := exec.Command(args[0], args[1:]...)
        cmd.Stdin = strings.NewReader(key.PublicKey)
        if err := cmd.Run(); err == nil {
            return nil
        }
    }
    return fmt.Errorf("no clipboard tool found (install xclip, xsel, or wl-copy)")
}

func GetFingerprint(key Key) (string, error) {
    cmd := exec.Command("ssh-keygen", "-l", "-f", key.Path+".pub")
    out, err := cmd.Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(out)), nil
}
func Delete(key Key) error {
    // Remove private key
    if err := os.Remove(key.Path); err != nil {
        return err
    }
    // Remove public key
    return os.Remove(key.Path + ".pub")
}