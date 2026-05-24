package ssh

import (
    "bufio"
    "os"
    "strconv"
    "strings"
    "github.com/mheaton92/quay/internal/connection"
)

func ImportSSHConfig() ([]connection.Connection, error) {
    path := os.Getenv("HOME") + "/.ssh/config"
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var connections []connection.Connection
    var current *connection.Connection

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        parts := strings.Fields(line)
        if len(parts) < 2 {
            continue
        }

        key := strings.ToLower(parts[0])
        value := strings.Join(parts[1:], " ")

        switch key {
        case "host":
            if current != nil && current.Name != "*" {
                connections = append(connections, *current)
            }
            current = &connection.Connection{
                Name: value,
                Host: value,
                Port: 22,
            }
        case "hostname":
            if current != nil {
                current.Host = value
            }
        case "user":
            if current != nil {
                current.User = value
            }
        case "port":
            if current != nil {
                if port, err := strconv.Atoi(value); err == nil {
                    current.Port = port
                }
            }
        case "identityfile":
            if current != nil {
                current.IdentityFile = value
            }
        case "proxyjump":
            if current != nil {
                current.ProxyJump = value
            }
        case "forwardagent":
            if current != nil {
                current.ForwardAgent = value
            }
        }
    }

    if current != nil && current.Name != "*" {
        connections = append(connections, *current)
    }

    return connections, nil
}
