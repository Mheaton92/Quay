package ssh

import (
	"fmt"
	"syscall"
	"github.com/mheaton92/quay/internal/connection"
	"os/exec"
	"os"
	"golang.org/x/term"
)

func BuildCmd(conn connection.Connection) *exec.Cmd {
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

	return exec.Command(args[0], args[1:]...)
}

func Connect(conn connection.Connection) error {
	cmd := BuildCmd(conn)
	return cmd.Run()
}

func ExecSSH(conn connection.Connection) error {
	sshCmd := BuildCmd(conn)
	sshPath, err := exec.LookPath(sshCmd.Path)
	if err != nil {
		return err
	}
	args := append([]string{sshPath}, sshCmd.Args[1:]...)
	// Restore terminal before SSH takes over
	if oldState, err := term.GetState(int(os.Stdin.Fd())); err == nil {
		term.Restore(int(os.Stdin.Fd()), oldState)
	}
	return syscall.Exec(sshPath, args, os.Environ())
}
