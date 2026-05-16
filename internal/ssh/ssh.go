package ssh

import (
	"fmt"
	"os"
	"github.com/mheaton92/quay/internal/connection"
	"os/exec"
)

func Connect(conn connection.Connection) error {
	args := []string{"ssh"}

	if conn.Port != 22 {
		args = append(args, "-p", fmt.Sprintf("%d", conn.Port))
	}
	if conn.IdentityFile != "" {
		args = append(args, "-i", conn.IdentityFile)
	}
	if conn.ProxyJump != "" {
		args = append(args, "-J", conn.ProxyJump)
	}
	if conn.ConnectTimeout != "" {
		args = append(args, "-o", "ConnectTimeout="+conn.ConnectTimeout)
	}
	if conn.ForwardAgent != "" {
		args = append(args, "-o", "ForwardAgent="+conn.ForwardAgent)
	}
	if conn.ServerAliveInterval != "" {
		args = append(args, "-o", "ServerAliveInterval="+conn.ServerAliveInterval)
	}
	if conn.ServerAliveCountMax != "" {
		args = append(args, "-o", "ServerAliveCountMax="+conn.ServerAliveCountMax)
	}
	if conn.TCPKeepAlive != "" {
		args = append(args, "-o", "TCPKeepAlive="+conn.TCPKeepAlive)
	}
	if conn.LocalForward != "" {
		args = append(args, "-L", conn.LocalForward)
	}
	if conn.RemoteForward != "" {
		args = append(args, "-R", conn.RemoteForward)
	}
	if conn.DynamicForward != "" {
		args = append(args, "-D", conn.DynamicForward)
	}
	if conn.Args != "" {
		args = append(args, conn.Args)
	}


	args = append(args, conn.User+"@"+conn.Host)
	cmd := exec.Command("ssh", args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}