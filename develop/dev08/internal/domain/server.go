package domain

import (
	"errors"
	"io"
	"strings"
)

func (d *Domain) ServeCommand(cmd string, args []string) (io.Reader, error) {
	switch cmd {
	case "cd":
		if err := d.cd(args); err != nil {
			d.Log.Error(err.Error())
			return nil, err
		}
	case "pwd":
		pwd, err := d.pwd()
		if err != nil {
			d.Log.Error(err.Error())
			return nil, err
		}
		return strings.NewReader(pwd), nil
	case "echo":
		return strings.NewReader(d.echo(args)), nil
	case "kill":
		out, err := d.kill(args)
		if err != nil {
			d.Log.Error(err.Error())
			return nil, err
		}
		return out, nil
	case "ps":
		out, err := d.ps()
		if err != nil {
			d.Log.Error(err.Error())
			return nil, err
		}
		return out, nil
	case "exec":
		out, err := d.exec(args[0], args[1:])
		if err != nil {
			d.Log.Error(err.Error())
			return nil, err
		}
		return out, nil
	case "exit":
		d.exit()
	default:
		return nil, errors.ErrUnsupported
	}

	return nil, nil
}
