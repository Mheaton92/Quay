package scp

import (
	"github.com/mheaton92/quay/internal/connection"
	"os/exec"
	"fmt"
	"os"
)

func Upload(conn connection.Connection, localPath string, remotePath string) error {
	args := []string{"scp"}

	if conn.Port != 22 {
		args = append(args, "-P", fmt.Sprintf("%d", conn.Port))
	}
	if conn.IdentityFile != "" {
		args = append(args, "-i", conn.IdentityFile)
	}
	
	args = append(args, localPath, conn.User+"@"+conn.Host+":"+remotePath)

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Download(conn connection.Connection, remotePath string, localPath string) error {
	args := []string{"scp"}

	if conn.Port != 22 {
		args = append(args, "-P", fmt.Sprintf("%d", conn.Port))
	}
	if conn.IdentityFile != "" {
		args = append(args, "-i", conn.IdentityFile)
	}
	
	args = append(args, conn.User+"@"+conn.Host+":"+remotePath, localPath)

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}