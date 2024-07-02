package domain

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func (d *Domain) cd(args []string) error {
	if len(args) == 0 {
		d.Log.Error("no such file or directory")
		return errors.New("cd: no such file or directory")
	}

	if err := os.Chdir(args[0]); err != nil {
		d.Log.Error(err.Error())
		return err
	}

	return nil
}

func (d *Domain) pwd() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		d.Log.Error(err.Error())
		return "", errors.New("pwd: " + err.Error())
	}

	return pwd, nil
}

func (d *Domain) echo(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return args[0]
}

func (d *Domain) kill(args []string) (io.Reader, error) {
	if len(args) == 0 {
		d.Log.Error("no process")
		return nil, errors.New("kill: no process")
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		d.Log.Error(err.Error())
		return nil, errors.New("kill: " + err.Error())
	}

	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		d.Log.Error(err.Error())
		return nil, errors.New("kill: " + err.Error())
	}
	return strings.NewReader(fmt.Sprintf("kill: %d", pid)), nil
}

func (d *Domain) ps() (io.Reader, error) {
	cmd := exec.Command("ps")

	out, err := cmd.Output()
	if err != nil {
		d.Log.Error(err.Error())
		return nil, errors.New("ps: " + err.Error())
	}

	return bytes.NewReader(out), nil
}

func (d *Domain) exec(command string, args []string) (io.Reader, error) {
	cmd := exec.Command(command, args...)

	out, err := cmd.Output()
	if err != nil {
		d.Log.Error(err.Error())
		return nil, errors.New(command + ": " + err.Error())
	}

	return bytes.NewReader(out), nil
}

func (d *Domain) exit() {
	/*
		Custom logic might be here
	*/
	os.Exit(0)
}
