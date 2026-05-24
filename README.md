# ⚓ Quay

> A terminal-native SSH manager and homelab control center

![quay screenshot](docs/screenshot.png)

Quay is a keyboard-driven TUI for managing SSH connections, monitoring your homelab, and running network diagnostics — all from the terminal.

## Features

### SSH Management
- Add, edit, and delete SSH connections
- Connect with a single keypress
- SCP file transfer — upload and download
- SSH key management — generate, deploy, copy, delete
- Import connections from `~/.ssh/config`
- Tracks last connected time and connection count

### Network Monitor
- Live ping latency and packet loss for all connections
- Sparkline graphs updated in real time
- Pin connections to the bottom bar — session or persistent

### Networking Tools
- **Port Scanner** — check if services are running
- **Wake on LAN** — wake sleeping machines remotely
- **DNS Lookup** — resolve hostnames and IPs
- **Traceroute** — visualize network hops
- **SSL Checker** — monitor certificate expiry
- **Subnet Scanner** — discover devices on your network
- **Bandwidth Test** — measure throughput over SSH

## Installation

### From source
```bash
git clone https://github.com/mheaton92/quay.git
cd quay
go build -o quay ./cmd/quay
sudo mv quay /usr/local/bin/
```

### Termux (Android)
```bash
pkg install golang git openssh
git clone https://github.com/mheaton92/quay.git
cd quay
go build -o quay ./cmd/quay
mv quay ~/bin/
```

## Keybinds

| Key | Action |
|-----|--------|
| `enter` | Connect via SSH |
| `a` | Add connection |
| `e` | Edit connection |
| `d` | Delete connection |
| `s` | SCP file transfer |
| `K` | SSH key manager |
| `N` | Networking tools |
| `i` | Import from ~/.ssh/config |
| `p` | Pin to monitor bar (session) |
| `P` | Pin to monitor bar (persistent) |
| `?` | Toggle keybind help |
| `j/k` | Navigate |
| `q` | Quit |

## Configuration

Connections are stored in `~/.config/quay/connections.json`

Keybinds can be customized in `~/.config/quay/keybinds.toml` (coming soon)

## Credits

See [CREDITS.md](CREDITS.md)

## License

MIT
